package table

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/coschain/contentos-go/common/encoding/kope"
	"github.com/coschain/contentos-go/iservices"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	ExtTrxTrxIdTable          uint32 = 1916120438
	ExtTrxBlockHeightTable    uint32 = 3799341326
	ExtTrxBlockTimeTable      uint32 = 1025113122
	ExtTrxTrxCreateOrderTable uint32 = 1760958085
	ExtTrxTrxIdUniTable       uint32 = 334659987

	ExtTrxTrxIdRow uint32 = 2158991352
)

////////////// SECTION Wrap Define ///////////////
type SoExtTrxWrap struct {
	dba         iservices.IDatabaseRW
	mainKey     *prototype.Sha256
	watcherFlag *ExtTrxWatcherFlag
	mKeyFlag    int    //the flag of the main key exist state in db, -1:has not judged; 0:not exist; 1:already exist
	mKeyBuf     []byte //the buffer after the main key is encoded with prefix
	mBuf        []byte //the value after the main key is encoded
	mdFuncMap   map[string]interface{}
}

func NewSoExtTrxWrap(dba iservices.IDatabaseRW, key *prototype.Sha256) *SoExtTrxWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoExtTrxWrap{dba, key, nil, -1, nil, nil, nil}
	result.initWatcherFlag()
	return result
}

func (s *SoExtTrxWrap) CheckExist() bool {
	if s.dba == nil {
		return false
	}
	if s.mKeyFlag != -1 {
		//if you have already obtained the existence status of the primary key, use it directly
		if s.mKeyFlag == 0 {
			return false
		}
		return true
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}
	if res == false {
		s.mKeyFlag = 0
	} else {
		s.mKeyFlag = 1
	}
	return res
}

func (s *SoExtTrxWrap) MustExist(errMsgs ...interface{}) *SoExtTrxWrap {
	if !s.CheckExist() {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.MustExist: %v not found", s.mainKey), errMsgs...))
	}
	return s
}

func (s *SoExtTrxWrap) MustNotExist(errMsgs ...interface{}) *SoExtTrxWrap {
	if s.CheckExist() {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.MustNotExist: %v already exists", s.mainKey), errMsgs...))
	}
	return s
}

func (s *SoExtTrxWrap) initWatcherFlag() {
	if s.watcherFlag == nil {
		s.watcherFlag = new(ExtTrxWatcherFlag)
		*(s.watcherFlag) = ExtTrxWatcherFlagOfDb(s.dba.ServiceId())
	}
}

func (s *SoExtTrxWrap) create(f func(tInfo *SoExtTrx)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoExtTrx{}
	f(val)
	if val.TrxId == nil {
		val.TrxId = s.mainKey
	}
	if s.CheckExist() {
		return errors.New("the main key is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}

	buf, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	err = s.dba.Put(keyBuf, buf)
	if err != nil {
		return err
	}

	// update srt list keys
	if err = s.insertAllSortKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.dba.Delete(keyBuf)
		return err
	}

	//update unique list
	if sucNames, err := s.insertAllUniKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.delUniKeysWithNames(sucNames, val)
		s.dba.Delete(keyBuf)
		return err
	}

	s.mKeyFlag = 1

	// call watchers
	s.initWatcherFlag()
	if s.watcherFlag.AnyWatcher {
		ReportTableRecordInsert(s.dba.ServiceId(), s.dba.BranchId(), s.mainKey, val)
	}

	return nil
}

func (s *SoExtTrxWrap) Create(f func(tInfo *SoExtTrx), errArgs ...interface{}) *SoExtTrxWrap {
	err := s.create(f)
	if err != nil {
		panic(bindErrorInfo(fmt.Errorf("SoExtTrxWrap.Create failed: %s", err.Error()), errArgs...))
	}
	return s
}

func (s *SoExtTrxWrap) getMainKeyBuf() ([]byte, error) {
	if s.mainKey == nil {
		return nil, errors.New("the main key is nil")
	}
	if s.mBuf == nil {
		var err error = nil
		s.mBuf, err = kope.Encode(s.mainKey)
		if err != nil {
			return nil, err
		}
	}
	return s.mBuf, nil
}

func (s *SoExtTrxWrap) modify(f func(tInfo *SoExtTrx)) error {
	if !s.CheckExist() {
		return errors.New("the SoExtTrx table does not exist. Please create a table first")
	}
	oriTable := s.getExtTrx()
	if oriTable == nil {
		return errors.New("fail to get origin table SoExtTrx")
	}

	curTable := s.getExtTrx()
	if curTable == nil {
		return errors.New("fail to create current table SoExtTrx")
	}
	f(curTable)

	//the main key is not support modify
	if !reflect.DeepEqual(curTable.TrxId, oriTable.TrxId) {
		return errors.New("primary key does not support modification")
	}

	s.initWatcherFlag()
	modifiedFields, hasWatcher, err := s.getModifiedFields(oriTable, curTable)
	if err != nil {
		return err
	}

	if modifiedFields == nil || len(modifiedFields) < 1 {
		return nil
	}

	//check whether modify sort and unique field to nil
	err = s.checkSortAndUniFieldValidity(curTable, modifiedFields)
	if err != nil {
		return err
	}

	//check unique
	err = s.handleFieldMd(FieldMdHandleTypeCheck, curTable, modifiedFields)
	if err != nil {
		return err
	}

	//delete sort and unique key
	err = s.handleFieldMd(FieldMdHandleTypeDel, oriTable, modifiedFields)
	if err != nil {
		return err
	}

	//update table
	err = s.updateExtTrx(curTable)
	if err != nil {
		return err
	}

	//insert sort and unique key
	err = s.handleFieldMd(FieldMdHandleTypeInsert, curTable, modifiedFields)
	if err != nil {
		return err
	}

	// call watchers
	if hasWatcher {
		ReportTableRecordUpdate(s.dba.ServiceId(), s.dba.BranchId(), s.mainKey, oriTable, curTable, modifiedFields)
	}

	return nil

}

func (s *SoExtTrxWrap) Modify(f func(tInfo *SoExtTrx), errArgs ...interface{}) *SoExtTrxWrap {
	err := s.modify(f)
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.Modify failed: %s", err.Error()), errArgs...))
	}
	return s
}

func (s *SoExtTrxWrap) SetBlockHeight(p uint64, errArgs ...interface{}) *SoExtTrxWrap {
	err := s.modify(func(r *SoExtTrx) {
		r.BlockHeight = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.SetBlockHeight( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoExtTrxWrap) SetBlockId(p *prototype.Sha256, errArgs ...interface{}) *SoExtTrxWrap {
	err := s.modify(func(r *SoExtTrx) {
		r.BlockId = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.SetBlockId( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoExtTrxWrap) SetBlockTime(p *prototype.TimePointSec, errArgs ...interface{}) *SoExtTrxWrap {
	err := s.modify(func(r *SoExtTrx) {
		r.BlockTime = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.SetBlockTime( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoExtTrxWrap) SetTrxCreateOrder(p *prototype.UserTrxCreateOrder, errArgs ...interface{}) *SoExtTrxWrap {
	err := s.modify(func(r *SoExtTrx) {
		r.TrxCreateOrder = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.SetTrxCreateOrder( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoExtTrxWrap) SetTrxWrap(p *prototype.TransactionWrapper, errArgs ...interface{}) *SoExtTrxWrap {
	err := s.modify(func(r *SoExtTrx) {
		r.TrxWrap = p
	})
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.SetTrxWrap( %v ) failed: %s", p, err.Error()), errArgs...))
	}
	return s
}

func (s *SoExtTrxWrap) checkSortAndUniFieldValidity(curTable *SoExtTrx, fields map[string]bool) error {
	if curTable != nil && fields != nil && len(fields) > 0 {

		if fields["BlockTime"] && curTable.BlockTime == nil {
			return errors.New("sort field BlockTime can't be modified to nil")
		}

		if fields["TrxCreateOrder"] && curTable.TrxCreateOrder == nil {
			return errors.New("sort field TrxCreateOrder can't be modified to nil")
		}

	}
	return nil
}

//Get all the modified fields in the table
func (s *SoExtTrxWrap) getModifiedFields(oriTable *SoExtTrx, curTable *SoExtTrx) (map[string]bool, bool, error) {
	if oriTable == nil {
		return nil, false, errors.New("table info is nil, can't get modified fields")
	}
	hasWatcher := false
	fields := make(map[string]bool)

	if !reflect.DeepEqual(oriTable.BlockHeight, curTable.BlockHeight) {
		fields["BlockHeight"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasBlockHeightWatcher
	}

	if !reflect.DeepEqual(oriTable.BlockId, curTable.BlockId) {
		fields["BlockId"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasBlockIdWatcher
	}

	if !reflect.DeepEqual(oriTable.BlockTime, curTable.BlockTime) {
		fields["BlockTime"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasBlockTimeWatcher
	}

	if !reflect.DeepEqual(oriTable.TrxCreateOrder, curTable.TrxCreateOrder) {
		fields["TrxCreateOrder"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasTrxCreateOrderWatcher
	}

	if !reflect.DeepEqual(oriTable.TrxWrap, curTable.TrxWrap) {
		fields["TrxWrap"] = true
		hasWatcher = hasWatcher || s.watcherFlag.HasTrxWrapWatcher
	}

	hasWatcher = hasWatcher || s.watcherFlag.WholeWatcher
	return fields, hasWatcher, nil
}

func (s *SoExtTrxWrap) handleFieldMd(t FieldMdHandleType, so *SoExtTrx, fields map[string]bool) error {
	if so == nil {
		return errors.New("fail to modify empty table")
	}

	//there is no field need to modify
	if fields == nil || len(fields) < 1 {
		return nil
	}

	errStr := ""

	if fields["BlockHeight"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldBlockHeight(so.BlockHeight, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "BlockHeight")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldBlockHeight(so.BlockHeight, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "BlockHeight")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldBlockHeight(so.BlockHeight, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "BlockHeight")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	if fields["BlockId"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldBlockId(so.BlockId, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "BlockId")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldBlockId(so.BlockId, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "BlockId")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldBlockId(so.BlockId, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "BlockId")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	if fields["BlockTime"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldBlockTime(so.BlockTime, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "BlockTime")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldBlockTime(so.BlockTime, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "BlockTime")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldBlockTime(so.BlockTime, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "BlockTime")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	if fields["TrxCreateOrder"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldTrxCreateOrder(so.TrxCreateOrder, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "TrxCreateOrder")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldTrxCreateOrder(so.TrxCreateOrder, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "TrxCreateOrder")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldTrxCreateOrder(so.TrxCreateOrder, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "TrxCreateOrder")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	if fields["TrxWrap"] {
		res := true
		if t == FieldMdHandleTypeCheck {
			res = s.mdFieldTrxWrap(so.TrxWrap, true, false, false, so)
			errStr = fmt.Sprintf("fail to modify exist value of %v", "TrxWrap")
		} else if t == FieldMdHandleTypeDel {
			res = s.mdFieldTrxWrap(so.TrxWrap, false, true, false, so)
			errStr = fmt.Sprintf("fail to delete  sort or unique field  %v", "TrxWrap")
		} else if t == FieldMdHandleTypeInsert {
			res = s.mdFieldTrxWrap(so.TrxWrap, false, false, true, so)
			errStr = fmt.Sprintf("fail to insert  sort or unique field  %v", "TrxWrap")
		}
		if !res {
			return errors.New(errStr)
		}
	}

	return nil
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoExtTrxWrap) delSortKeyTrxId(sa *SoExtTrx) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListExtTrxByTrxId{}
	if sa == nil {
		val.TrxId = s.GetTrxId()
	} else {
		val.TrxId = sa.TrxId
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) insertSortKeyTrxId(sa *SoExtTrx) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListExtTrxByTrxId{}
	val.TrxId = sa.TrxId
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) delSortKeyBlockHeight(sa *SoExtTrx) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListExtTrxByBlockHeight{}
	if sa == nil {
		val.BlockHeight = s.GetBlockHeight()
		val.TrxId = s.mainKey

	} else {
		val.BlockHeight = sa.BlockHeight
		val.TrxId = sa.TrxId
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) insertSortKeyBlockHeight(sa *SoExtTrx) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListExtTrxByBlockHeight{}
	val.TrxId = sa.TrxId
	val.BlockHeight = sa.BlockHeight
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) delSortKeyBlockTime(sa *SoExtTrx) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListExtTrxByBlockTime{}
	if sa == nil {
		val.BlockTime = s.GetBlockTime()
		val.TrxId = s.mainKey

	} else {
		val.BlockTime = sa.BlockTime
		val.TrxId = sa.TrxId
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) insertSortKeyBlockTime(sa *SoExtTrx) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListExtTrxByBlockTime{}
	val.TrxId = sa.TrxId
	val.BlockTime = sa.BlockTime
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) delSortKeyTrxCreateOrder(sa *SoExtTrx) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListExtTrxByTrxCreateOrder{}
	if sa == nil {
		val.TrxCreateOrder = s.GetTrxCreateOrder()
		val.TrxId = s.mainKey

	} else {
		val.TrxCreateOrder = sa.TrxCreateOrder
		val.TrxId = sa.TrxId
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) insertSortKeyTrxCreateOrder(sa *SoExtTrx) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListExtTrxByTrxCreateOrder{}
	val.TrxId = sa.TrxId
	val.TrxCreateOrder = sa.TrxCreateOrder
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoExtTrxWrap) delAllSortKeys(br bool, val *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delSortKeyTrxId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyBlockHeight(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyBlockTime(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyTrxCreateOrder(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoExtTrxWrap) insertAllSortKeys(val *SoExtTrx) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoExtTrx fail ")
	}
	if !s.insertSortKeyTrxId(val) {
		return errors.New("insert sort Field TrxId fail while insert table ")
	}
	if !s.insertSortKeyBlockHeight(val) {
		return errors.New("insert sort Field BlockHeight fail while insert table ")
	}
	if !s.insertSortKeyBlockTime(val) {
		return errors.New("insert sort Field BlockTime fail while insert table ")
	}
	if !s.insertSortKeyTrxCreateOrder(val) {
		return errors.New("insert sort Field TrxCreateOrder fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoExtTrxWrap) removeExtTrx() error {
	if s.dba == nil {
		return errors.New("database is nil")
	}

	s.initWatcherFlag()

	var oldVal *SoExtTrx
	if s.watcherFlag.AnyWatcher {
		oldVal = s.getExtTrx()
	}

	//delete sort list key
	if res := s.delAllSortKeys(true, nil); !res {
		return errors.New("delAllSortKeys failed")
	}

	//delete unique list
	if res := s.delAllUniKeys(true, nil); !res {
		return errors.New("delAllUniKeys failed")
	}

	//delete table
	key, err := s.encodeMainKey()
	if err != nil {
		return fmt.Errorf("encodeMainKey failed: %s", err.Error())
	}
	err = s.dba.Delete(key)
	if err == nil {
		s.mKeyBuf = nil
		s.mKeyFlag = -1

		// call watchers
		if s.watcherFlag.AnyWatcher && oldVal != nil {
			ReportTableRecordDelete(s.dba.ServiceId(), s.dba.BranchId(), s.mainKey, oldVal)
		}
		return nil
	} else {
		return fmt.Errorf("database.Delete failed: %s", err.Error())
	}
}

func (s *SoExtTrxWrap) RemoveExtTrx(errMsgs ...interface{}) *SoExtTrxWrap {
	err := s.removeExtTrx()
	if err != nil {
		panic(bindErrorInfo(fmt.Sprintf("SoExtTrxWrap.RemoveExtTrx failed: %s", err.Error()), errMsgs...))
	}
	return s
}

////////////// SECTION Members Get/Modify ///////////////

func (s *SoExtTrxWrap) GetBlockHeight() uint64 {
	res := true
	msg := &SoExtTrx{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.BlockHeight
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.BlockHeight
}

func (s *SoExtTrxWrap) mdFieldBlockHeight(p uint64, isCheck bool, isDel bool, isInsert bool,
	so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkBlockHeightIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldBlockHeight(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldBlockHeight(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoExtTrxWrap) delFieldBlockHeight(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if !s.delSortKeyBlockHeight(so) {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) insertFieldBlockHeight(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if !s.insertSortKeyBlockHeight(so) {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) checkBlockHeightIsMetMdCondition(p uint64) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) GetBlockId() *prototype.Sha256 {
	res := true
	msg := &SoExtTrx{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.BlockId
			}
		}
	}
	if !res {
		return nil

	}
	return msg.BlockId
}

func (s *SoExtTrxWrap) mdFieldBlockId(p *prototype.Sha256, isCheck bool, isDel bool, isInsert bool,
	so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkBlockIdIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldBlockId(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldBlockId(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoExtTrxWrap) delFieldBlockId(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) insertFieldBlockId(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) checkBlockIdIsMetMdCondition(p *prototype.Sha256) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) GetBlockTime() *prototype.TimePointSec {
	res := true
	msg := &SoExtTrx{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.BlockTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.BlockTime
}

func (s *SoExtTrxWrap) mdFieldBlockTime(p *prototype.TimePointSec, isCheck bool, isDel bool, isInsert bool,
	so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkBlockTimeIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldBlockTime(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldBlockTime(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoExtTrxWrap) delFieldBlockTime(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if !s.delSortKeyBlockTime(so) {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) insertFieldBlockTime(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if !s.insertSortKeyBlockTime(so) {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) checkBlockTimeIsMetMdCondition(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) GetTrxCreateOrder() *prototype.UserTrxCreateOrder {
	res := true
	msg := &SoExtTrx{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.TrxCreateOrder
			}
		}
	}
	if !res {
		return nil

	}
	return msg.TrxCreateOrder
}

func (s *SoExtTrxWrap) mdFieldTrxCreateOrder(p *prototype.UserTrxCreateOrder, isCheck bool, isDel bool, isInsert bool,
	so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkTrxCreateOrderIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldTrxCreateOrder(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldTrxCreateOrder(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoExtTrxWrap) delFieldTrxCreateOrder(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if !s.delSortKeyTrxCreateOrder(so) {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) insertFieldTrxCreateOrder(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if !s.insertSortKeyTrxCreateOrder(so) {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) checkTrxCreateOrderIsMetMdCondition(p *prototype.UserTrxCreateOrder) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) GetTrxId() *prototype.Sha256 {
	res := true
	msg := &SoExtTrx{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.TrxId
			}
		}
	}
	if !res {
		return nil

	}
	return msg.TrxId
}

func (s *SoExtTrxWrap) GetTrxWrap() *prototype.TransactionWrapper {
	res := true
	msg := &SoExtTrx{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMainKey()
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.TrxWrap
			}
		}
	}
	if !res {
		return nil

	}
	return msg.TrxWrap
}

func (s *SoExtTrxWrap) mdFieldTrxWrap(p *prototype.TransactionWrapper, isCheck bool, isDel bool, isInsert bool,
	so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	if isCheck {
		res := s.checkTrxWrapIsMetMdCondition(p)
		if !res {
			return false
		}
	}

	if isDel {
		res := s.delFieldTrxWrap(so)
		if !res {
			return false
		}
	}

	if isInsert {
		res := s.insertFieldTrxWrap(so)
		if !res {
			return false
		}
	}
	return true
}

func (s *SoExtTrxWrap) delFieldTrxWrap(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) insertFieldTrxWrap(so *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}

	return true
}

func (s *SoExtTrxWrap) checkTrxWrapIsMetMdCondition(p *prototype.TransactionWrapper) bool {
	if s.dba == nil {
		return false
	}

	return true
}

////////////// SECTION List Keys ///////////////
type SExtTrxTrxIdWrap struct {
	Dba iservices.IDatabaseRW
}

func NewExtTrxTrxIdWrap(db iservices.IDatabaseRW) *SExtTrxTrxIdWrap {
	if db == nil {
		return nil
	}
	wrap := SExtTrxTrxIdWrap{Dba: db}
	return &wrap
}

func (s *SExtTrxTrxIdWrap) GetMainVal(val []byte) *prototype.Sha256 {
	res := &SoListExtTrxByTrxId{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.TrxId

}

func (s *SExtTrxTrxIdWrap) GetSubVal(val []byte) *prototype.Sha256 {
	res := &SoListExtTrxByTrxId{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.TrxId

}

func (m *SoListExtTrxByTrxId) OpeEncode() ([]byte, error) {
	pre := ExtTrxTrxIdTable
	sub := m.TrxId
	if sub == nil {
		return nil, errors.New("the pro TrxId is nil")
	}
	sub1 := m.TrxId
	if sub1 == nil {
		return nil, errors.New("the mainkey TrxId is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtTrxTrxIdWrap) ForEachByOrder(start *prototype.Sha256, end *prototype.Sha256, lastMainKey *prototype.Sha256,
	lastSubVal *prototype.Sha256, f func(mVal *prototype.Sha256, sVal *prototype.Sha256, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtTrxTrxIdTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SExtTrxBlockHeightWrap struct {
	Dba iservices.IDatabaseRW
}

func NewExtTrxBlockHeightWrap(db iservices.IDatabaseRW) *SExtTrxBlockHeightWrap {
	if db == nil {
		return nil
	}
	wrap := SExtTrxBlockHeightWrap{Dba: db}
	return &wrap
}

func (s *SExtTrxBlockHeightWrap) GetMainVal(val []byte) *prototype.Sha256 {
	res := &SoListExtTrxByBlockHeight{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.TrxId

}

func (s *SExtTrxBlockHeightWrap) GetSubVal(val []byte) *uint64 {
	res := &SoListExtTrxByBlockHeight{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return &res.BlockHeight

}

func (m *SoListExtTrxByBlockHeight) OpeEncode() ([]byte, error) {
	pre := ExtTrxBlockHeightTable
	sub := m.BlockHeight

	sub1 := m.TrxId
	if sub1 == nil {
		return nil, errors.New("the mainkey TrxId is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtTrxBlockHeightWrap) ForEachByOrder(start *uint64, end *uint64, lastMainKey *prototype.Sha256,
	lastSubVal *uint64, f func(mVal *prototype.Sha256, sVal *uint64, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtTrxBlockHeightTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SExtTrxBlockTimeWrap struct {
	Dba iservices.IDatabaseRW
}

func NewExtTrxBlockTimeWrap(db iservices.IDatabaseRW) *SExtTrxBlockTimeWrap {
	if db == nil {
		return nil
	}
	wrap := SExtTrxBlockTimeWrap{Dba: db}
	return &wrap
}

func (s *SExtTrxBlockTimeWrap) GetMainVal(val []byte) *prototype.Sha256 {
	res := &SoListExtTrxByBlockTime{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.TrxId

}

func (s *SExtTrxBlockTimeWrap) GetSubVal(val []byte) *prototype.TimePointSec {
	res := &SoListExtTrxByBlockTime{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.BlockTime

}

func (m *SoListExtTrxByBlockTime) OpeEncode() ([]byte, error) {
	pre := ExtTrxBlockTimeTable
	sub := m.BlockTime
	if sub == nil {
		return nil, errors.New("the pro BlockTime is nil")
	}
	sub1 := m.TrxId
	if sub1 == nil {
		return nil, errors.New("the mainkey TrxId is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtTrxBlockTimeWrap) ForEachByOrder(start *prototype.TimePointSec, end *prototype.TimePointSec, lastMainKey *prototype.Sha256,
	lastSubVal *prototype.TimePointSec, f func(mVal *prototype.Sha256, sVal *prototype.TimePointSec, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtTrxBlockTimeTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

//Query srt by reverse order
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtTrxBlockTimeWrap) ForEachByRevOrder(start *prototype.TimePointSec, end *prototype.TimePointSec, lastMainKey *prototype.Sha256,
	lastSubVal *prototype.TimePointSec, f func(mVal *prototype.Sha256, sVal *prototype.TimePointSec, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtTrxBlockTimeTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey)
		}
		skeyList = append(skeyList, kope.MaximumKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(eBuf, sBuf, true, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SExtTrxTrxCreateOrderWrap struct {
	Dba iservices.IDatabaseRW
}

func NewExtTrxTrxCreateOrderWrap(db iservices.IDatabaseRW) *SExtTrxTrxCreateOrderWrap {
	if db == nil {
		return nil
	}
	wrap := SExtTrxTrxCreateOrderWrap{Dba: db}
	return &wrap
}

func (s *SExtTrxTrxCreateOrderWrap) GetMainVal(val []byte) *prototype.Sha256 {
	res := &SoListExtTrxByTrxCreateOrder{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.TrxId

}

func (s *SExtTrxTrxCreateOrderWrap) GetSubVal(val []byte) *prototype.UserTrxCreateOrder {
	res := &SoListExtTrxByTrxCreateOrder{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.TrxCreateOrder

}

func (m *SoListExtTrxByTrxCreateOrder) OpeEncode() ([]byte, error) {
	pre := ExtTrxTrxCreateOrderTable
	sub := m.TrxCreateOrder
	if sub == nil {
		return nil, errors.New("the pro TrxCreateOrder is nil")
	}
	sub1 := m.TrxId
	if sub1 == nil {
		return nil, errors.New("the mainkey TrxId is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtTrxTrxCreateOrderWrap) ForEachByOrder(start *prototype.UserTrxCreateOrder, end *prototype.UserTrxCreateOrder, lastMainKey *prototype.Sha256,
	lastSubVal *prototype.UserTrxCreateOrder, f func(mVal *prototype.Sha256, sVal *prototype.UserTrxCreateOrder, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtTrxTrxCreateOrderTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

//Query srt by reverse order
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SExtTrxTrxCreateOrderWrap) ForEachByRevOrder(start *prototype.UserTrxCreateOrder, end *prototype.UserTrxCreateOrder, lastMainKey *prototype.Sha256,
	lastSubVal *prototype.UserTrxCreateOrder, f func(mVal *prototype.Sha256, sVal *prototype.UserTrxCreateOrder, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := ExtTrxTrxCreateOrderTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey)
		}
		skeyList = append(skeyList, kope.MaximumKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(eBuf, sBuf, true, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

/////////////// SECTION Private function ////////////////

func (s *SoExtTrxWrap) update(sa *SoExtTrx) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	buf, err := proto.Marshal(sa)
	if err != nil {
		return false
	}

	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	return s.dba.Put(keyBuf, buf) == nil
}

func (s *SoExtTrxWrap) getExtTrx() *SoExtTrx {
	if s.dba == nil {
		return nil
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return nil
	}
	resBuf, err := s.dba.Get(keyBuf)

	if err != nil {
		return nil
	}

	res := &SoExtTrx{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoExtTrxWrap) updateExtTrx(so *SoExtTrx) error {
	if s.dba == nil {
		return errors.New("update fail:the db is nil")
	}

	if so == nil {
		return errors.New("update fail: the SoExtTrx is nil")
	}

	key, err := s.encodeMainKey()
	if err != nil {
		return nil
	}

	buf, err := proto.Marshal(so)
	if err != nil {
		return err
	}

	err = s.dba.Put(key, buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *SoExtTrxWrap) encodeMainKey() ([]byte, error) {
	if s.mKeyBuf != nil {
		return s.mKeyBuf, nil
	}
	pre := ExtTrxTrxIdRow
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	preBuf, err := kope.Encode(pre)
	if err != nil {
		return nil, err
	}
	mBuf, err := s.getMainKeyBuf()
	if err != nil {
		return nil, err
	}
	list := make([][]byte, 2)
	list[0] = preBuf
	list[1] = mBuf
	s.mKeyBuf = kope.PackList(list)
	return s.mKeyBuf, nil
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoExtTrxWrap) delAllUniKeys(br bool, val *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyTrxId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoExtTrxWrap) delUniKeysWithNames(names map[string]string, val *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if len(names["TrxId"]) > 0 {
		if !s.delUniKeyTrxId(val) {
			res = false
		}
	}

	return res
}

func (s *SoExtTrxWrap) insertAllUniKeys(val *SoExtTrx) (map[string]string, error) {
	if s.dba == nil {
		return nil, errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return nil, errors.New("insert uniuqe Field fail,get the SoExtTrx fail ")
	}
	sucFields := map[string]string{}
	if !s.insertUniKeyTrxId(val) {
		return sucFields, errors.New("insert unique Field TrxId fail while insert table ")
	}
	sucFields["TrxId"] = "TrxId"

	return sucFields, nil
}

func (s *SoExtTrxWrap) delUniKeyTrxId(sa *SoExtTrx) bool {
	if s.dba == nil {
		return false
	}
	pre := ExtTrxTrxIdUniTable
	kList := []interface{}{pre}
	if sa != nil {
		if sa.TrxId == nil {
			return false
		}

		sub := sa.TrxId
		kList = append(kList, sub)
	} else {
		sub := s.GetTrxId()
		if sub == nil {
			return true
		}

		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoExtTrxWrap) insertUniKeyTrxId(sa *SoExtTrx) bool {
	if s.dba == nil || sa == nil {
		return false
	}

	pre := ExtTrxTrxIdUniTable
	sub := sa.TrxId
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	res, err := s.dba.Has(kBuf)
	if err == nil && res == true {
		//the unique key is already exist
		return false
	}
	val := SoUniqueExtTrxByTrxId{}
	val.TrxId = sa.TrxId

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniExtTrxTrxIdWrap struct {
	Dba iservices.IDatabaseRW
}

func NewUniExtTrxTrxIdWrap(db iservices.IDatabaseRW) *UniExtTrxTrxIdWrap {
	if db == nil {
		return nil
	}
	wrap := UniExtTrxTrxIdWrap{Dba: db}
	return &wrap
}

func (s *UniExtTrxTrxIdWrap) UniQueryTrxId(start *prototype.Sha256) *SoExtTrxWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := ExtTrxTrxIdUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueExtTrxByTrxId{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoExtTrxWrap(s.Dba, res.TrxId)

			return wrap
		}
	}
	return nil
}

////////////// SECTION Watchers ///////////////

type ExtTrxWatcherFlag struct {
	HasBlockHeightWatcher bool

	HasBlockIdWatcher bool

	HasBlockTimeWatcher bool

	HasTrxCreateOrderWatcher bool

	HasTrxWrapWatcher bool

	WholeWatcher bool
	AnyWatcher   bool
}

var (
	ExtTrxTable = &TableInfo{
		Name:    "ExtTrx",
		Primary: "TrxId",
		Record:  reflect.TypeOf((*SoExtTrx)(nil)).Elem(),
	}
	ExtTrxWatcherFlags     = make(map[uint32]ExtTrxWatcherFlag)
	ExtTrxWatcherFlagsLock sync.RWMutex
)

func ExtTrxWatcherFlagOfDb(dbSvcId uint32) ExtTrxWatcherFlag {
	ExtTrxWatcherFlagsLock.RLock()
	defer ExtTrxWatcherFlagsLock.RUnlock()
	return ExtTrxWatcherFlags[dbSvcId]
}

func ExtTrxRecordWatcherChanged(dbSvcId uint32) {
	var flag ExtTrxWatcherFlag
	flag.WholeWatcher = HasTableRecordWatcher(dbSvcId, ExtTrxTable.Record, "")
	flag.AnyWatcher = flag.WholeWatcher

	flag.HasBlockHeightWatcher = HasTableRecordWatcher(dbSvcId, ExtTrxTable.Record, "BlockHeight")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasBlockHeightWatcher

	flag.HasBlockIdWatcher = HasTableRecordWatcher(dbSvcId, ExtTrxTable.Record, "BlockId")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasBlockIdWatcher

	flag.HasBlockTimeWatcher = HasTableRecordWatcher(dbSvcId, ExtTrxTable.Record, "BlockTime")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasBlockTimeWatcher

	flag.HasTrxCreateOrderWatcher = HasTableRecordWatcher(dbSvcId, ExtTrxTable.Record, "TrxCreateOrder")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasTrxCreateOrderWatcher

	flag.HasTrxWrapWatcher = HasTableRecordWatcher(dbSvcId, ExtTrxTable.Record, "TrxWrap")
	flag.AnyWatcher = flag.AnyWatcher || flag.HasTrxWrapWatcher

	ExtTrxWatcherFlagsLock.Lock()
	ExtTrxWatcherFlags[dbSvcId] = flag
	ExtTrxWatcherFlagsLock.Unlock()
}

////////////// SECTION Json query ///////////////

func ExtTrxQuery(db iservices.IDatabaseRW, keyJson string) (valueJson string, err error) {
	k := new(prototype.Sha256)
	d := json.NewDecoder(bytes.NewReader([]byte(keyJson)))
	d.UseNumber()
	if err = d.Decode(k); err != nil {
		return
	}
	if v := NewSoExtTrxWrap(db, k).getExtTrx(); v == nil {
		err = errors.New("not found")
	} else {
		var jbytes []byte
		if jbytes, err = json.Marshal(v); err == nil {
			valueJson = string(jbytes)
		}
	}
	return
}

func init() {
	RegisterTableWatcherChangedCallback(ExtTrxTable.Record, ExtTrxRecordWatcherChanged)
	RegisterTableJsonQuery("ExtTrx", ExtTrxQuery)
}

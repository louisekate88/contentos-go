package plugins

import (
	"database/sql"
	"fmt"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/iservices/itype"
	"github.com/coschain/contentos-go/iservices/service-configs"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/prototype"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

//type Row map[string]int

const INTERVAL = 60 * 60
// for test easily
//const INTERVAL = 60 * 5

type DailyStatisticService struct {
	node.Service
	config *service_configs.DatabaseConfig
	consensus iservices.IConsensus
	outDb *sql.DB
	log *logrus.Logger
	ctx *node.ServiceContext
	ticker *time.Ticker
	quit chan bool
}

func NewDailyStatisticService(ctx *node.ServiceContext, config *service_configs.DatabaseConfig, log *logrus.Logger) (*DailyStatisticService, error) {
	return &DailyStatisticService{ctx: ctx, log: log, config: config}, nil
}

func (s *DailyStatisticService) Start(node *node.Node) error {
	s.quit = make(chan bool)
	// dns: data source name
	consensus, err := s.ctx.Service(iservices.ConsensusServerName)
	if err != nil {
		return err
	}
	s.consensus = consensus.(iservices.IConsensus)
	dsn := fmt.Sprintf("%s:%s@/%s", s.config.User, s.config.Password, s.config.Db)
	outDb, err := sql.Open(s.config.Driver, dsn)

	if err != nil {
		return err
	}
	s.outDb = outDb

	s.ticker = time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <- s.ticker.C:
				if err := s.pollLIB(); err != nil {
					s.log.Error(err)
				}
			case <- s.quit:
				s.stop()
				return
			}
		}
	}()
	return nil
}

// trigger by ticker cannot statistic past dau
// trigger by lib is ok
func (s *DailyStatisticService) pollLIB() error  {
	lib := s.consensus.GetLIB().BlockNum()
	s.log.Debugf("[trx db] sync lib: %d \n", lib)
	var lastLib uint64 = 0
	var lastDate string
	_ = s.outDb.QueryRow("SELECT lib, date from dailystatinfo limit 1").Scan(&lastLib, &lastDate)
	//updateStmt, _ := s.outDb.Prepare("UPDATE dxulibinfo SET lib=?, last_check_time=?")
	//defer updateStmt.Close()
	var waitingSyncLib []uint64
	var count = 0
	for lastLib < lib {
		if count > 100 {
			break
		}
		waitingSyncLib = append(waitingSyncLib, lastLib)
		lastLib ++
		count ++
	}
	for _, block := range waitingSyncLib {
		blks , err := s.consensus.FetchBlocks(block, block)
		if err != nil {
			s.log.Error(err)
		}
		if len(blks) == 0 {
			s.log.Errorf("cannot fetch block %d in consensus", block)
		}
		blk := blks[0].(*prototype.SignedBlock)
		blockTime := blk.Timestamp()
		datetime := time.Unix(int64(blockTime), 0)
		date := fmt.Sprintf("%d-%02d-%02d", datetime.Year(), datetime.Month(), datetime.Day())
		// trigger
		if lastDate != date {
			s.handleDAUStatistic(blk)
			s.handleDNUStatistic(blk)
		}
		utcTimestamp := time.Now().UTC().Unix()
		_, _ = s.outDb.Exec("UPDATE dailystatinfo SET lib = ?, date = ?, last_check_time = ?", block, date, utcTimestamp)
	}
	return nil
}

func (s *DailyStatisticService) handleDAUStatistic(block *prototype.SignedBlock) {
	blockTime := block.Timestamp()
	datetime := time.Unix(int64(blockTime), 0)
	date := fmt.Sprintf("%d-%02d-%02d", datetime.Year(), datetime.Month(), datetime.Day())

	end := blockTime
	start := end - 86400
	statRows, _ := s.outDb.Query("SELECT creator, count(distinct creator) FROM trxinfo WHERE block_time >= ? and block_time < ? group by creator", start, end)
	dapps := make(map[string]string)
	dappRows, _ := s.outDb.Query("select dapp, prefix from dailystatdapp where status=1")
	for dappRows.Next() {
		var dapp, prefix string
		_ = dappRows.Scan(&dapp, &prefix)
		dapps[prefix] = dapp
	}
	for statRows.Next() {
		var (
			creator string
			counter uint32
		)
		if err := statRows.Scan(&creator, &counter); err != nil {
			s.log.Error(err)
			continue
		}
		for prefix, dapp := range dapps {
			if strings.HasPrefix(creator, prefix) {
				_, _ = s.outDb.Exec("insert ignore into daustat (date, dapp, count) values (?, ?, ?)", date, dapp, counter)
			}
		}
	}
}

func (s *DailyStatisticService) handleDNUStatistic(block *prototype.SignedBlock) {
	blockTime := block.Timestamp()
	datetime := time.Unix(int64(blockTime), 0)
	date := fmt.Sprintf("%d-%02d-%02d", datetime.Year(), datetime.Month(), datetime.Day())

	end := blockTime
	start := end - 86400
	statRows, _ := s.outDb.Query("SELECT creator, count(distinct creator) FROM createaccountinfo WHERE create_time >= ? and create_time < ? group by creator", start, end)
	dapps := make(map[string]string)
	dappRows, _ := s.outDb.Query("select dapp, prefix from dailystatdapp where status=1")
	for dappRows.Next() {
		var dapp, prefix string
		_ = dappRows.Scan(&dapp, &prefix)
		dapps[prefix] = dapp
	}
	for statRows.Next() {
		var (
			creator string
			counter uint32
		)
		if err := statRows.Scan(&creator, &counter); err != nil {
			s.log.Error(err)
			continue
		}
		for prefix, dapp := range dapps {
			if strings.HasPrefix(creator, prefix) {
				_, _ = s.outDb.Exec("insert ignore into dnustat (date, dapp, count) values (?, ?, ?)", date, dapp, counter)
			}
		}
	}
}

func (s *DailyStatisticService) DAUStatsOn(date string, dapp string) *itype.Row {
	var count uint32
	_ = s.outDb.QueryRow("select count from daustat where date=? and dapp=?", date, dapp).Scan(&count)
	return &itype.Row{Date: date, Dapp: dapp, Count: count}
}

func (s *DailyStatisticService) DAUStatsSince(days int, dapp string) []*itype.Row {
	now := time.Now().UTC()
	d, _ := time.ParseDuration("-1d")
	then := now.Add(d * time.Duration(days))
	date := fmt.Sprintf("%d-%02d-%02d", then.Year(), then.Month(), then.Day())
	var dauRows []*itype.Row
	rows, err := s.outDb.Query("select count from daustat where date >= ? and dapp = ?  order by date", date, dapp)
	if err != nil {
		return dauRows
	}
	for rows.Next() {
		var count uint32
		_ = rows.Scan(&count)
		r := &itype.Row{Date: date, Dapp: dapp, Count:count}
		dauRows = append(dauRows, r)
	}
	return dauRows
}

func (s *DailyStatisticService) DNUStatsOn(date string, dapp string) *itype.Row {
	var count uint32
	_ = s.outDb.QueryRow("select count from dnustat where date=? and dapp=?", date, dapp).Scan(&count)
	return &itype.Row{Date: date, Dapp: dapp, Count: count}
}

func (s *DailyStatisticService) DNUStatsSince(days int, dapp string) []*itype.Row {
	now := time.Now().UTC()
	d, _ := time.ParseDuration("-1d")
	then := now.Add(d * time.Duration(days))
	date := fmt.Sprintf("%d-%02d-%02d", then.Year(), then.Month(), then.Day())
	var dauRows []*itype.Row
	rows, err := s.outDb.Query("select count from dnustat where date >= ? and dapp = ?  order by date", date, dapp)
	if err != nil {
		return dauRows
	}
	for rows.Next() {
		var count uint32
		_ = rows.Scan(&count)
		r := &itype.Row{Date: date, Dapp: dapp, Count:count}
		dauRows = append(dauRows, r)
	}
	return dauRows
}

func (s *DailyStatisticService) stop() {
	_ = s.outDb.Close()
	s.ticker.Stop()
}


func (t *DailyStatisticService) Stop() error {
	t.quit <- true
	close(t.quit)
	return nil
}
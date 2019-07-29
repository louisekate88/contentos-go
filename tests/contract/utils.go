package contracts

import (
	"errors"
	"fmt"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/dandelion"
	"github.com/coschain/contentos-go/prototype"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func Deploy(d *dandelion.Dandelion, owner, contract string) (err error) {
	var (
		code, abi []byte
	)
	if code, err = ioutil.ReadFile("./testdata/" + contract + ".wasm"); err != nil {
		return err
	}
	if abi, err = ioutil.ReadFile("./testdata/" + contract + ".abi"); err != nil {
		return err
	}
	r := d.Account(owner).TrxReceipt(dandelion.ContractDeployUncompressed(owner, contract, abi, code, true, "", ""))
	if r == nil || r.Status != prototype.StatusSuccess {
		err = fmt.Errorf("contract deployment failed, receipt = %v", r)
	}
	return
}

func Apply(d *dandelion.Dandelion, caller, owner, contract, method, jsonParams string, amount uint64) *prototype.TransactionReceiptWithInfo {
	return d.Account(caller).TrxReceipt(dandelion.ContractApply(caller, owner, contract, method, jsonParams, amount))
}

func applyAndCheck(t *testing.T, checker func(*prototype.TransactionReceiptWithInfo) error, d *dandelion.Dandelion, caller, owner, contract, method, jsonParams string, amount uint64) {
	if err := checker(Apply(d, caller, owner, contract, method, jsonParams, amount)); err != nil {
		t.Fatal(err)
	}
}

func extractCall(call string) (caller, owner, contract, method, params string, amount uint64, err error) {
	re, _ := regexp.Compile("\\s*(\\w+):\\s*(\\d*)\\s+(\\w+)\\.(\\w+).(\\w+)\\s*")
	if matches := re.FindStringSubmatch(call); len(matches) >= 6 {
		caller, owner, contract, method = matches[1], matches[3], matches[4], matches[5]
		params = fmt.Sprintf("[%s]", strings.Trim(call[len(matches[0]):], " \t\r\n"))
		if len(matches[2]) > 0 {
			if coins, err := strconv.ParseUint(matches[2], 10, 64); err == nil {
				amount = coins
			}
		}
	} else {
		err = errors.New("invalid call-string")
	}
	return
}

func ApplyNoError(t *testing.T, d *dandelion.Dandelion, call string) {
	caller, owner, contract, method, params, amount, err := extractCall(call)
	if err != nil {
		t.Fatal(err)
	}
	applyAndCheck(t, func(r *prototype.TransactionReceiptWithInfo) error {
		if r == nil {
			return errors.New("apply failed, receipt is nil")
		}
		if r.Status != prototype.StatusSuccess {
			return fmt.Errorf("apply failed, receipt.status=%d, err=%s", r.Status, r.ErrorInfo)
		}
		return nil
	}, d, caller, owner, contract, method, params, amount)
}

func ApplyError(t *testing.T, d *dandelion.Dandelion, call string) {
	caller, owner, contract, method, params, amount, err := extractCall(call)
	if err != nil {
		t.Fatal(err)
	}
	applyAndCheck(t, func(r *prototype.TransactionReceiptWithInfo) error {
		if r == nil {
			return errors.New("apply failed, receipt is nil")
		}
		if r.Status == prototype.StatusSuccess {
			return errors.New("apply succeeded, but a failure was expected")
		}
		return nil
	}, d, caller, owner, contract, method, params, amount)
}

func NewDandelionContractTest(f func(*testing.T, *dandelion.Dandelion), actors int, contracts...string) func(*testing.T) {
	return dandelion.NewDandelionTest(func(t *testing.T, d *dandelion.Dandelion) {
		const stakeAmount = 10000 * constants.COSTokenDecimals
		var err error
		err = d.Account(constants.COSInitMiner).SendTrx(dandelion.Stake(constants.COSInitMiner, constants.COSInitMiner, stakeAmount))
		if err != nil {
			t.Fatalf("staking %d (%s->%s) error: %s", stakeAmount, constants.COSInitMiner, constants.COSInitMiner, err.Error())
		}
		for i := 0; i < actors; i++ {
			name := fmt.Sprintf("actor%d", i)
			err = d.Account(constants.COSInitMiner).SendTrx(dandelion.Stake(constants.COSInitMiner, name, stakeAmount))
			if err != nil {
				t.Fatalf("staking %d (%s->%s) error: %s", stakeAmount, constants.COSInitMiner, name, err.Error())
			}
		}
		err = d.ProduceBlocks(1)
		if err != nil {
			t.Fatal(err)
		}
		for _, contract := range contracts {
			parts := strings.Split(contract, ".")
			if len(parts) == 2 {
				err = Deploy(d, parts[0], parts[1])
				if err != nil {
					t.Fatal(err.Error())
				}
			}
		}
		f(t, d)
	}, actors)
}

func BytesToJson(data []byte) string {
	var s []string
	for _, b := range data {
		s = append(s, strconv.FormatUint(uint64(b), 10))
	}
	return fmt.Sprintf("[%s]", strings.Join(s, ","))
}

package op

import (
	"github.com/coschain/contentos-go/dandelion"
	"testing"
)

func TestOperations(t *testing.T) {
	t.Run("transfer", dandelion.NewDandelionTest(new(TransferTester).Test, 3))
	t.Run("bp", dandelion.NewDandelionTest(new(BpTest).Test, 3))
	t.Run("vote", dandelion.NewDandelionTest(new(VoteTester).TestNormal, 3))
	t.Run("vote", dandelion.NewDandelionTest(new(VoteTester).TestRevote, 3))
	t.Run("vote", dandelion.NewDandelionTest(new(VoteTester).TestFullPower, 3))
	t.Run("follow", dandelion.NewDandelionTest(new(FollowTester).Test, 3))
	t.Run("transfer to vesting", dandelion.NewDandelionTest(new(TransferToVestingTester).Test, 3))
	t.Run("contract_deploy", dandelion.NewDandelionTest(new(ContractDeployTester).Test, 3))
	t.Run("contract_lib", NewDandelionContractTest(new(ContractTester).Test, 2, "actor0.native_tester", "actor1.native_tester"))
}

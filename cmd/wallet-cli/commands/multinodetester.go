package commands

import (
	"context"
	"fmt"
	"github.com/coschain/cobra"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/constants"
	//"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"strconv"
	//"time"
)

var MultinodetesterCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "multinodetester",
		Short:   "multinodetester count",
		Example: "multinodetester count",
		Args:    cobra.ExactArgs(1),
		Run:     multinodetester,
	}
	return cmd
}

func transferCosToNewAccountTrx(cmd *cobra.Command, client grpcpb.ApiServiceClient, count int64) (*prototype.SignedTransaction, error) {
	resp, _ := client.GetChainState( context.Background(), &grpcpb.NonParamsRequest{} )
	refBlockPrefix := common.TaposRefBlockPrefix(resp.State.Dgpo.HeadBlockId.Hash)
	refBlockNum := common.TaposRefBlockNum(resp.State.Dgpo.HeadBlockNumber)
	tx := &prototype.Transaction{RefBlockNum: refBlockNum, RefBlockPrefix: refBlockPrefix, Expiration: &prototype.TimePointSec{UtcSeconds: resp.State.Dgpo.Time.UtcSeconds + 30}}
	trx := &prototype.SignedTransaction{Trx: tx}

	fromAccount := prototype.NewAccountName(constants.COSInitMiner)
	fromAccountPriKey, err := prototype.PrivateKeyFromWIF(constants.InitminerPrivKey)
	if err != nil {
		return nil, err
	}

	for index := int64(1); index < count; index++ {
		toAccountName := fmt.Sprintf("%s%d", constants.COSInitMiner, index)

		op_transfer := &prototype.TransferOperation{
			From:   fromAccount,
			To:     &prototype.AccountName{Value: toAccountName},
			Amount: prototype.NewCoin(constants.MinBpRegisterVest),
			Memo:   "",
		}

		trx.Trx.AddOperation(op_transfer)
	}

	res := trx.Sign(fromAccountPriKey, cmd.Context["chain_id"].(prototype.ChainId))
	trx.Signature = &prototype.SignatureType{Sig: res}

	if err := trx.Validate(); err != nil {
		return nil, err
	}
	return trx, nil
}

func transferCosToVestTrx(cmd *cobra.Command, client grpcpb.ApiServiceClient, count int64) (*prototype.SignedTransaction, error) {
	resp, _ := client.GetChainState( context.Background(), &grpcpb.NonParamsRequest{} )
	refBlockPrefix := common.TaposRefBlockPrefix(resp.State.Dgpo.HeadBlockId.Hash)
	refBlockNum := common.TaposRefBlockNum(resp.State.Dgpo.HeadBlockNumber)
	tx := &prototype.Transaction{RefBlockNum: refBlockNum, RefBlockPrefix: refBlockPrefix, Expiration: &prototype.TimePointSec{UtcSeconds: resp.State.Dgpo.Time.UtcSeconds + 30}}
	trx := &prototype.SignedTransaction{Trx: tx}

	accountName := fmt.Sprintf("%s%d", constants.COSInitMiner, count)
	keys, err := prototype.FixBytesToPrivateKey([]byte(accountName))
	if err != nil {
		return nil, err
	}

	transferv_op := &prototype.TransferToVestOperation{
		From:   &prototype.AccountName{Value: accountName},
		To:     &prototype.AccountName{Value: accountName},
		Amount: prototype.NewCoin(constants.MinBpRegisterVest),
	}

	trx.Trx.AddOperation(transferv_op)

	res := trx.Sign(keys, cmd.Context["chain_id"].(prototype.ChainId))
	trx.Signature = &prototype.SignatureType{Sig: res}

	if err := trx.Validate(); err != nil {
		return nil, err
	}
	return trx, nil
}


func makeBpRegVoteTrx(cmd *cobra.Command, client grpcpb.ApiServiceClient, count int64) (*prototype.SignedTransaction, error) {

	resp, _ := client.GetChainState( context.Background(), &grpcpb.NonParamsRequest{} )
	refBlockPrefix := common.TaposRefBlockPrefix(resp.State.Dgpo.HeadBlockId.Hash)
	refBlockNum := common.TaposRefBlockNum(resp.State.Dgpo.HeadBlockNumber)
	tx := &prototype.Transaction{RefBlockNum: refBlockNum, RefBlockPrefix: refBlockPrefix, Expiration: &prototype.TimePointSec{UtcSeconds: resp.State.Dgpo.Time.UtcSeconds + 30}}
	trx := &prototype.SignedTransaction{Trx: tx}

	bpName := fmt.Sprintf("%s%d", constants.COSInitMiner, count)
	keys, err := prototype.FixBytesToPrivateKey([]byte(bpName))
	if err != nil {
		return nil, err
	}

	pubKey, err := keys.PubKey()
	if err != nil {
		return nil, err
	}

	opBpReg := &prototype.BpRegisterOperation{
		Owner:           &prototype.AccountName{Value: bpName},
		Url:             bpName,
		Desc:            bpName,
		BlockSigningKey: pubKey,
		Props: &prototype.ChainProperties{
			AccountCreationFee: prototype.NewCoin(constants.DefaultAccountCreateFee),
			StaminaFree:        constants.DefaultStaminaFree,
			TpsExpected:        constants.DefaultTPSExpected,
			TopNAcquireFreeToken: constants.InitTopN,
			EpochDuration: constants.InitEpochDuration,
			PerTicketPrice: prototype.NewCoin(constants.PerTicketPrice * constants.COSTokenDecimals),
			PerTicketWeight: constants.PerTicketWeight,
		},
	}

	opBpVote := &prototype.BpVoteOperation{Voter: prototype.NewAccountName(bpName), BlockProducer: prototype.NewAccountName(bpName), Cancel: false}

	trx.Trx.AddOperation(opBpReg)
	trx.Trx.AddOperation(opBpVote)


	res := trx.Sign(keys, cmd.Context["chain_id"].(prototype.ChainId))
	trx.Signature = &prototype.SignatureType{Sig: res}

	if err := trx.Validate(); err != nil {
		return nil, err
	}
	return trx, nil
}


func createMNTAccountTrx(cmd *cobra.Command, client grpcpb.ApiServiceClient, count int64) (*prototype.SignedTransaction, error) {

	resp, _ := client.GetChainState( context.Background(), &grpcpb.NonParamsRequest{} )
	refBlockPrefix := common.TaposRefBlockPrefix(resp.State.Dgpo.HeadBlockId.Hash)
	refBlockNum := common.TaposRefBlockNum(resp.State.Dgpo.HeadBlockNumber)
	tx := &prototype.Transaction{RefBlockNum: refBlockNum, RefBlockPrefix: refBlockPrefix, Expiration: &prototype.TimePointSec{UtcSeconds: resp.State.Dgpo.Time.UtcSeconds + 30}}
	trx := &prototype.SignedTransaction{Trx: tx}

	creator := prototype.NewAccountName(constants.COSInitMiner)

	creatorPriKey, err := prototype.PrivateKeyFromWIF(constants.InitminerPrivKey)
	if err != nil {
		return nil, err
	}

	opCreatorBpVote := &prototype.BpVoteOperation{Voter: creator, BlockProducer: creator, Cancel: false}

	trx.Trx.AddOperation(opCreatorBpVote)

	for index := int64(1); index < count; index++ {
		bpName := fmt.Sprintf("%s%d", constants.COSInitMiner, index)
		keys, err := prototype.FixBytesToPrivateKey([]byte(bpName))
		if err != nil {
			return nil, err
		}

		pubKey, err := keys.PubKey()
		if err != nil {
			return nil, err
		}

		opCreate := &prototype.AccountCreateOperation{
			Fee:            prototype.NewCoin(constants.DefaultAccountCreateFee),
			Creator:        creator,
			NewAccountName: &prototype.AccountName{Value: bpName},
			PubKey:          pubKey,
		}

		trx.Trx.AddOperation(opCreate)
	}

	res := trx.Sign(creatorPriKey, cmd.Context["chain_id"].(prototype.ChainId))
	trx.Signature = &prototype.SignatureType{Sig: res}

	if err := trx.Validate(); err != nil {
		return nil, err
	}
	return trx, nil
}

/*
func makeMultiNodeTeseterTrx(count int64, onlyCreate bool) (*prototype.SignedTransaction, error) {

	var priKey *prototype.PrivateKeyType = nil

	tx := &prototype.Transaction{RefBlockNum: 0, RefBlockPrefix: 0, Expiration: &prototype.TimePointSec{UtcSeconds: uint32(time.Now().Unix()) + 30}}
	trx := &prototype.SignedTransaction{Trx: tx}

	creator := prototype.NewAccountName(constants.COSInitMiner)

	creatorPriKey, err := prototype.PrivateKeyFromWIF(constants.InitminerPrivKey)
	if err != nil {
		return nil, err
	}

	opCreatorBpVote := &prototype.BpVoteOperation{Voter: creator, BlockProducer: creator, Cancel: false}

	if !onlyCreate {
		trx.Trx.AddOperation(opCreatorBpVote)
	}

	for index := int64(1); index < count; index++ {
		bpName := fmt.Sprintf("%s%d", constants.COSInitMiner, index)
		keys, err := prototype.FixBytesToPrivateKey([]byte(bpName))
		if err != nil {
			return nil, err
		}

		pubKey, err := keys.PubKey()
		if err != nil {
			return nil, err
		}

		opCreate := &prototype.AccountCreateOperation{
			Fee:            prototype.NewCoin(1),
			Creator:        creator,
			NewAccountName: &prototype.AccountName{Value: bpName},
			Owner:          prototype.NewAuthorityFromPubKey(pubKey),
		}

		opBpReg := &prototype.BpRegisterOperation{
			Owner:           &prototype.AccountName{Value: bpName},
			Url:             bpName,
			Desc:            bpName,
			BlockSigningKey: pubKey,
			Props: &prototype.ChainProperties{
				AccountCreationFee: prototype.NewCoin(1),
				MaximumBlockSize:   10 * 1024 * 1024,
			},
		}

		opBpVote := &prototype.BpVoteOperation{Voter: prototype.NewAccountName(bpName), BlockProducer: prototype.NewAccountName(bpName), Cancel: false}

		if !onlyCreate {
			trx.Trx.AddOperation(opBpReg)
			trx.Trx.AddOperation(opBpVote)
			priKeys = append(priKeys, keys)
		} else {
			trx.Trx.AddOperation(opCreate)
		}

	}

	priKeys = append(priKeys, creatorPriKey)

	for _, k := range priKeys {
		res := trx.Sign(k, prototype.ChainId{Value: 0})
		trx.Signatures = append(trx.Signatures, &prototype.SignatureType{Sig: res})
	}

	if err := trx.Validate(); err != nil {
		return nil, err
	}
	return trx, nil
}*/

func multinodetester(cmd *cobra.Command, args []string) {
	c := cmd.Context["rpcclient"]
	client := c.(grpcpb.ApiServiceClient)
	idx, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create new account
	{
		signTx, err := createMNTAccountTrx(cmd, client, idx)
		if err != nil {
			fmt.Println(err)
			return
		}
		req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
		resp, err := client.BroadcastTrx(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("create Result: %v", resp))
		}
	}

	// transfer cos to new account
	{
		signTx, err := transferCosToNewAccountTrx(cmd, client, idx)
		if err != nil {
			fmt.Println(err)
			return
		}
		req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
		resp, err := client.BroadcastTrx(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("transfer Result: %v", resp))
		}
	}

	var i int64
	// new account transfer cos to vest
	for i = 1; i < idx; i++ {
		signTx, err := transferCosToVestTrx(cmd, client, i)
		if err != nil {
			fmt.Println(err)
			return
		}
		req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
		resp, err := client.BroadcastTrx(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("transfer to VEST Result: %v", resp))
		}
	}

	// new account register and vote bp
	for i = 1; i < idx; i++ {
		signTx, err := makeBpRegVoteTrx(cmd, client, i)
		if err != nil {
			fmt.Println(err)
			return
		}
		req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
		//req.OnlyDeliver = true
		resp, err := client.BroadcastTrx(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("bpvote Result: %v", resp))
		}
	}
}

package txsFee

import (
	"math/big"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	wasmConfig "github.com/TerraDharitri/drt-go-chain-vm/config"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm/txsFee/utils"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/stretchr/testify/require"
)

func TestMultiDCDTTransferShouldWork(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	rcvAddr := []byte("12345678901234567890123456789022")

	rewaBalance := big.NewInt(100000000)
	dcdtBalance := big.NewInt(100000000)
	token := []byte("miiutoken")
	utils.CreateAccountWithDCDTBalance(t, testContext.Accounts, sndAddr, rewaBalance, token, 0, dcdtBalance, uint32(core.Fungible))
	secondToken := []byte("second")
	utils.CreateAccountWithDCDTBalance(t, testContext.Accounts, sndAddr, big.NewInt(0), secondToken, 0, dcdtBalance, uint32(core.Fungible))

	gasLimit := uint64(4000)
	tx := utils.CreateMultiTransferTX(0, sndAddr, rcvAddr, gasPrice, gasLimit, &utils.TransferDCDTData{
		Token: token,
		Value: big.NewInt(100),
	}, &utils.TransferDCDTData{
		Token: secondToken,
		Value: big.NewInt(200),
	})
	retCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Equal(t, vmcommon.Ok, retCode)
	require.Nil(t, err)
	require.Nil(t, testContext.GetCompositeTestError())

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	expectedBalanceSnd := big.NewInt(99999900)
	utils.CheckDCDTBalance(t, testContext, sndAddr, token, expectedBalanceSnd)

	expectedReceiverBalance := big.NewInt(100)
	utils.CheckDCDTBalance(t, testContext, rcvAddr, token, expectedReceiverBalance)

	expectedBalanceSndSecondToken := big.NewInt(99999800)
	utils.CheckDCDTBalance(t, testContext, sndAddr, secondToken, expectedBalanceSndSecondToken)

	expectedReceiverBalanceSecondToken := big.NewInt(200)
	utils.CheckDCDTBalance(t, testContext, rcvAddr, secondToken, expectedReceiverBalanceSecondToken)

	expectedREWABalance := big.NewInt(99960000)
	utils.TestAccount(t, testContext.Accounts, sndAddr, 1, expectedREWABalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(40000), accumulatedFees)

	allLogs := testContext.TxsLogsProcessor.GetAllCurrentLogs()
	require.NotNil(t, allLogs)
}

func TestMultiDCDTTransferFailsBecauseOfMaxLimit(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMsAndCustomGasSchedule(config.EnableEpochs{},
		func(gasMap wasmConfig.GasScheduleMap) {
			gasMap[common.MaxPerTransaction]["MaxNumberOfTransfersPerTx"] = 1
		}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	rcvAddr := []byte("12345678901234567890123456789022")

	rewaBalance := big.NewInt(100000000)
	dcdtBalance := big.NewInt(100000000)
	token := []byte("miiutoken")
	utils.CreateAccountWithDCDTBalance(t, testContext.Accounts, sndAddr, rewaBalance, token, 0, dcdtBalance, uint32(core.Fungible))
	secondToken := []byte("second")
	utils.CreateAccountWithDCDTBalance(t, testContext.Accounts, sndAddr, big.NewInt(0), secondToken, 0, dcdtBalance, uint32(core.Fungible))

	gasLimit := uint64(4000)
	tx := utils.CreateMultiTransferTX(0, sndAddr, rcvAddr, gasPrice, gasLimit, &utils.TransferDCDTData{
		Token: token,
		Value: big.NewInt(100),
	}, &utils.TransferDCDTData{
		Token: secondToken,
		Value: big.NewInt(200),
	})
	retCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.NotNil(t, err)
	require.Equal(t, vmcommon.UserError, retCode)
	require.Contains(t, testContext.GetCompositeTestError().Error(), process.ErrMaxCallsReached.Error())
}

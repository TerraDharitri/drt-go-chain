package txsFee

import (
	"encoding/hex"
	"math/big"
	"testing"

	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/integrationTests"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm/txsFee/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRelayedAsyncCallShouldWork(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	senderAddr := []byte("12345678901234567890123456789011")

	t.Run("nonce fix is disabled, should increase the sender's nonce", func(t *testing.T) {
		testContext := testRelayedAsyncCallShouldWork(t, config.EnableEpochs{
			RelayedNonceFixEnableEpoch: 100000,
		}, senderAddr)
		defer testContext.Close()

		senderAccount := getAccount(t, testContext, senderAddr)
		assert.Equal(t, uint64(1), senderAccount.GetNonce())
	})
	t.Run("nonce fix is enabled, should still increase the sender's nonce", func(t *testing.T) {
		testContext := testRelayedAsyncCallShouldWork(t, config.EnableEpochs{
			RelayedNonceFixEnableEpoch: 0,
		}, senderAddr)
		defer testContext.Close()

		senderAccount := getAccount(t, testContext, senderAddr)
		assert.Equal(t, uint64(1), senderAccount.GetNonce())
	})
}

func testRelayedAsyncCallShouldWork(t *testing.T, enableEpochs config.EnableEpochs, senderAddr []byte) *vm.VMTestContext {
	testContext, err := vm.CreatePreparedTxProcessorWithVMs(enableEpochs, 1)
	require.Nil(t, err)

	localRewaBalance := big.NewInt(100000000)
	ownerAddr := []byte("12345678901234567890123456789010")
	relayerAddr := []byte("12345678901234567890123456789017")

	_, _ = vm.CreateAccount(testContext.Accounts, ownerAddr, 0, localRewaBalance)
	_, _ = vm.CreateAccount(testContext.Accounts, relayerAddr, 0, localRewaBalance)

	ownerAccount, _ := testContext.Accounts.LoadAccount(ownerAddr)
	deployGasLimit := uint64(50000)

	pathToContract := "testdata/first/output/first.wasm"
	firstScAddress := utils.DoDeploySecond(t, testContext, pathToContract, ownerAccount, gasPrice, deployGasLimit, nil, big.NewInt(50))

	gasLimit := uint64(5000000)
	args := [][]byte{[]byte(hex.EncodeToString(firstScAddress))}
	ownerAccount, _ = testContext.Accounts.LoadAccount(ownerAddr)
	pathToContract = "testdata/second/output/async.wasm"
	secondSCAddress := utils.DoDeploySecond(t, testContext, pathToContract, ownerAccount, gasPrice, deployGasLimit, args, big.NewInt(50))

	utils.CleanAccumulatedIntermediateTransactions(t, testContext)
	testContext.TxFeeHandler.CreateBlockStarted(getZeroGasAndFees())

	innerTx := vm.CreateTransaction(0, big.NewInt(0), senderAddr, secondSCAddress, gasPrice, gasLimit, []byte("doSomething"))

	rtxData := integrationTests.PrepareRelayedTxDataV1(innerTx)
	rTxGasLimit := 1 + gasLimit + uint64(len(rtxData))
	rtx := vm.CreateTransaction(0, innerTx.Value, relayerAddr, senderAddr, gasPrice, rTxGasLimit, rtxData)

	retCode, err := testContext.TxProcessor.ProcessTransaction(rtx)
	require.Equal(t, vmcommon.Ok, retCode)
	require.Nil(t, err)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	utils.TestAccount(t, testContext.Accounts, relayerAddr, 1, big.NewInt(49998050))

	intermediateTxs := testContext.GetIntermediateTransactions(t)
	require.NotNil(t, intermediateTxs)

	res := vm.GetIntValueFromSC(nil, testContext.Accounts, firstScAddress, "numCalled")
	require.Equal(t, big.NewInt(1), res)

	require.Equal(t, big.NewInt(50001950), testContext.TxFeeHandler.GetAccumulatedFees())
	require.Equal(t, big.NewInt(4999988), testContext.TxFeeHandler.GetDeveloperFees())

	return testContext
}

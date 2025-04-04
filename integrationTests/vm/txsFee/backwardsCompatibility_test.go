package txsFee

import (
	"errors"
	"math/big"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/data/receipt"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/integrationTests"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// minGasPrice = 1, gasPerDataByte = 1, minGasLimit = 1
func TestMoveBalanceSelfShouldWorkAndConsumeTxFeeWhenAllFlagsAreDisabled(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{
		PenalizedTooMuchGasEnableEpoch: 100,
		SCDeployEnableEpoch:            100,
		MetaProtectionEnableEpoch:      100,
		RelayedTransactionsEnableEpoch: 100,
	}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(10000)
	gasLimit := uint64(100)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(senderNonce, big.NewInt(100), sndAddr, sndAddr, gasPrice, gasLimit, []byte("aaaa"))

	returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Nil(t, err)
	require.Equal(t, vmcommon.Ok, returnCode)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// 10_000 - gasPrice(10)*gasLimit(1) + gasPerDataByte(1)*gasPrice(10)*4 = 10000 - 50 = 9950
	expectedBalance := big.NewInt(9950)
	vm.TestAccount(t, testContext.Accounts, sndAddr, senderNonce+1, expectedBalance)

	// check receipts
	require.Equal(t, 1, len(testContext.GetIntermediateTransactions(t)))
	rcpt := testContext.GetIntermediateTransactions(t)[0].(*receipt.Receipt)
	assert.Equal(t, "950", rcpt.Value.String())

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(50), accumulatedFees)
}

// minGasPrice = 1, gasPerDataByte = 1, minGasLimit = 1
func TestMoveBalanceAllFlagsDisabledLessBalanceThanGasLimitMulGasPrice(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{
		PenalizedTooMuchGasEnableEpoch: integrationTests.UnreachableEpoch,
		BuiltInFunctionsEnableEpoch:    integrationTests.UnreachableEpoch,
		SCDeployEnableEpoch:            integrationTests.UnreachableEpoch,
		MetaProtectionEnableEpoch:      integrationTests.UnreachableEpoch,
		RelayedTransactionsEnableEpoch: integrationTests.UnreachableEpoch,
	}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(10000)
	gasLimit := uint64(10000)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(senderNonce, big.NewInt(100), sndAddr, sndAddr, gasPrice, gasLimit, []byte("aaaa"))

	_, err = testContext.TxProcessor.ProcessTransaction(tx)
	require.True(t, errors.Is(err, process.ErrInsufficientFee))
}

func TestMoveBalanceSelfShouldWorkAndConsumeTxFeeWhenSomeFlagsAreDisabled(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(
		config.EnableEpochs{
			PenalizedTooMuchGasEnableEpoch:   0,
			BuiltInFunctionsEnableEpoch:      100,
			SCDeployEnableEpoch:              100,
			MetaProtectionEnableEpoch:        100,
			RelayedTransactionsV2EnableEpoch: 100,
		}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(10000)
	gasLimit := uint64(100)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(senderNonce, big.NewInt(100), sndAddr, sndAddr, gasPrice, gasLimit, []byte("aaaa"))

	returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Nil(t, err)
	require.Equal(t, vmcommon.Ok, returnCode)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// 10_000 - gasPrice(10)*gasLimit(1) + gasPerDataByte(1)*gasPrice(10)*4 = 10000 - 50 = 9950
	expectedBalance := big.NewInt(9950)
	vm.TestAccount(t, testContext.Accounts, sndAddr, senderNonce+1, expectedBalance)

	// check receipts
	require.Equal(t, 1, len(testContext.GetIntermediateTransactions(t)))
	rcpt := testContext.GetIntermediateTransactions(t)[0].(*receipt.Receipt)
	assert.Equal(t, "950", rcpt.Value.String())

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(50), accumulatedFees)
}

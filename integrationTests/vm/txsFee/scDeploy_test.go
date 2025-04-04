package txsFee

import (
	"math/big"
	"testing"

	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm/wasm"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/stretchr/testify/require"
)

func TestScDeployShouldWork(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(100000)
	gasLimit := uint64(1962)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)

	scCode := wasm.GetSCCode("../wasm/testdata/misc/fib_wasm/output/fib_wasm.wasm")
	tx := vm.CreateTransaction(senderNonce, big.NewInt(0), sndAddr, vm.CreateEmptyAddress(), gasPrice, gasLimit, []byte(wasm.CreateDeployTxData(scCode)))

	_, err = testContext.TxProcessor.ProcessTransaction(tx)
	require.Nil(t, err)
	require.Nil(t, testContext.GetCompositeTestError())

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// 8490 gas units the sc deploy consumed
	expectedBalance := big.NewInt(80400)
	vm.TestAccount(t, testContext.Accounts, sndAddr, senderNonce+1, expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(19600), accumulatedFees)
}

func TestScDeployInvalidContractCodeShouldConsumeGas(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(100000)
	gasLimit := uint64(1960)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)

	scCode := wasm.GetSCCode("../wasm/testdata/misc/fib_wasm/output/fib_wasm.wasm")
	scCodeBytes := []byte(scCode)
	scCodeBytes = append(scCodeBytes, []byte("aaa")...)
	txDeployData := []byte(wasm.CreateDeployTxData(string(scCodeBytes)))
	tx := vm.CreateTransaction(senderNonce, big.NewInt(0), sndAddr, vm.CreateEmptyAddress(), gasPrice, gasLimit, txDeployData)

	_, err = testContext.TxProcessor.ProcessTransaction(tx)
	require.Nil(t, err)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	expectedBalance := big.NewInt(80400)
	vm.TestAccount(t, testContext.Accounts, sndAddr, senderNonce+1, expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(19600), accumulatedFees)
}

func TestScDeployInsufficientGasLimitShouldNotConsumeGas(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(100000)
	gasLimit := uint64(568)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)

	scCode := wasm.GetSCCode("../wasm/testdata/misc/fib_wasm/output/fib_wasm.wasm")
	txDeployData := []byte(wasm.CreateDeployTxData(scCode))
	tx := vm.CreateTransaction(senderNonce, big.NewInt(0), sndAddr, vm.CreateEmptyAddress(), gasPrice, gasLimit, txDeployData)

	_, err = testContext.TxProcessor.ProcessTransaction(tx)
	require.Equal(t, process.ErrInsufficientGasLimitInTx, err)
	require.Nil(t, testContext.GetCompositeTestError())

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	expectedBalance := big.NewInt(0).Set(senderBalance)
	vm.TestAccount(t, testContext.Accounts, sndAddr, senderNonce, expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(0), accumulatedFees)
}

func TestScDeployOutOfGasShouldConsumeGas(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{}, 1)
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(13100)
	gasLimit := uint64(1310)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)

	scCode := wasm.GetSCCode("../wasm/testdata/misc/fib_wasm/output/fib_wasm.wasm")
	txDeployData := []byte(wasm.CreateDeployTxData(scCode))
	tx := vm.CreateTransaction(senderNonce, big.NewInt(0), sndAddr, vm.CreateEmptyAddress(), gasPrice, gasLimit, txDeployData)

	returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Nil(t, err)
	require.Equal(t, returnCode, vmcommon.UserError)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	expectedBalance := big.NewInt(0)
	vm.TestAccount(t, testContext.Accounts, sndAddr, 1, expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(13100), accumulatedFees)
}

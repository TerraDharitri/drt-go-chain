package scenariosConverter

import (
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-chain-scenario/scenario/exporter"
	scenmodel "github.com/TerraDharitri/drt-go-chain-scenario/scenario/model"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/stretchr/testify/require"
)

var log = logger.GetOrCreate("scenariosConverter")

// CheckAccounts will verify if scenariosAccounts correspond to AccountsAdapter accounts
func CheckAccounts(t *testing.T, accAdapter state.AccountsAdapter, scenariosAccounts []*exporter.TestAccount) {
	for _, scenariosAcc := range scenariosAccounts {
		accHandler, err := accAdapter.LoadAccount(scenariosAcc.GetAddress())
		require.Nil(t, err)
		account := accHandler.(state.UserAccountHandler)

		require.Equal(t, scenariosAcc.GetBalance(), account.GetBalance())
		require.Equal(t, scenariosAcc.GetNonce(), account.GetNonce())

		scOwnerAddress := scenariosAcc.GetOwner()
		if len(scOwnerAddress) == 0 {
			require.Nil(t, account.GetOwnerAddress())
		} else {
			require.Equal(t, scenariosAcc.GetOwner(), account.GetOwnerAddress())
		}

		codeHash := account.GetCodeHash()
		code := accAdapter.GetCode(codeHash)
		require.Equal(t, len(scenariosAcc.GetCode()), len(code))

		scenariosAccStorage := scenariosAcc.GetStorage()
		CheckStorage(t, account, scenariosAccStorage)
	}
}

// CheckStorage checks if the dataTrie of an account equals with the storage of the corresponding scenariosAccount
func CheckStorage(t *testing.T, dataTrie state.UserAccountHandler, scenariosAccStorage map[string][]byte) {
	for key := range scenariosAccStorage {
		dataTrieValue, _, err := dataTrie.RetrieveValue([]byte(key))
		require.Nil(t, err)
		if len(scenariosAccStorage[key]) == 0 {
			require.Nil(t, dataTrieValue)
		} else {
			require.Equal(t, scenariosAccStorage[key], dataTrieValue)
		}
	}
}

// CheckTransactions checks if the transactions correspond with the scenariosTransactions
func CheckTransactions(t *testing.T, transactions []*transaction.Transaction, scenariosTransactions []*exporter.Transaction) {
	expectedLength := len(scenariosTransactions)
	require.Equal(t, expectedLength, len(transactions))
	for i := 0; i < expectedLength; i++ {
		expectedSender := scenariosTransactions[i].GetSenderAddress()
		expectedReceiver := scenariosTransactions[i].GetReceiverAddress()
		expectedCallValue := scenariosTransactions[i].GetCallValue()
		expectedCallFunction := scenariosTransactions[i].GetCallFunction()
		expectedCallArguments := scenariosTransactions[i].GetCallArguments()
		expectedDcdtTransfers := scenariosTransactions[i].GetDCDTTransfers()
		expectedGasLimit, expectedGasPrice := scenariosTransactions[i].GetGasLimitAndPrice()
		expectedNonce := scenariosTransactions[i].GetNonce()

		require.Equal(t, expectedSender, transactions[i].GetSndAddr())
		require.Equal(t, expectedCallValue, transactions[i].GetValue())
		require.Equal(t, expectedGasLimit, transactions[i].GetGasLimit())
		require.Equal(t, expectedGasPrice, transactions[i].GetGasPrice())
		require.Equal(t, expectedNonce, transactions[i].GetNonce())

		var expectedData []byte
		if len(expectedDcdtTransfers) != 0 {
			expectedData = scenmodel.CreateMultiTransferData(expectedReceiver, expectedDcdtTransfers, expectedCallFunction, expectedCallArguments)
			require.Equal(t, expectedSender, transactions[i].GetRcvAddr())
		} else {
			require.Equal(t, expectedReceiver, transactions[i].GetRcvAddr())
			expectedData = createData(expectedCallFunction, expectedCallArguments)
		}

		actualData := transactions[i].GetData()
		require.Equal(t, expectedData, actualData)
	}
}

// BenchmarkScenariosSpecificTx -
func BenchmarkScenariosSpecificTx(b *testing.B, scenariosTestPath string) {
	testContext, transactions, benchmarkTxPos, err := SetStateFromScenariosTest(scenariosTestPath)
	if err != nil {
		log.Trace("Setting state went wrong:", "error", err)
		return
	}
	defer testContext.Close()
	if benchmarkTxPos == exporter.InvalidBenchmarkTxPos {
		log.Trace("no transactions marked for benchmarking")
	}
	if len(transactions) > 1 {
		err = ProcessAllTransactions(testContext, transactions[:benchmarkTxPos])
		if err != nil {
			log.Trace("Processing transactions went wrong:", "error", err)
		}
	}

	err = RunSingleTransactionBenchmark(b, testContext, transactions[benchmarkTxPos])
	if err != nil {
		log.Trace("Processing benchmark transaction went wrong:", "error", err)
	}
}

// SetStateFromScenariosTest recieves path to scenariosTest, returns a VMTestContext with the specified accounts, an array with the specified transactions and an error
func SetStateFromScenariosTest(scenariosTestPath string) (testContext *vm.VMTestContext, transactions []*transaction.Transaction, bechmarkTxPos int, err error) {
	stateAndBenchmarkInfo, err := exporter.GetAccountsAndTransactionsFromScenarios(scenariosTestPath)
	if err != nil {
		return nil, nil, exporter.InvalidBenchmarkTxPos, err
	}
	testContext, err = vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{}, 1)
	if err != nil {
		return nil, nil, exporter.InvalidBenchmarkTxPos, err
	}
	err = CreateAccountsFromScenariosAccs(testContext, stateAndBenchmarkInfo.Accs)
	if err != nil {
		return nil, nil, exporter.InvalidBenchmarkTxPos, err
	}
	newAddresses, err := DeploySCsFromScenariosDeployTxs(testContext, stateAndBenchmarkInfo.DeployTxs)
	if err != nil {
		return nil, nil, exporter.InvalidBenchmarkTxPos, err
	}
	ReplaceScenariosScAddressesWithNewScAddresses(stateAndBenchmarkInfo.DeployedAccs, newAddresses, stateAndBenchmarkInfo.Txs)
	transactions = CreateTransactionsFromScenariosTxs(stateAndBenchmarkInfo.Txs)
	return testContext, transactions, stateAndBenchmarkInfo.BenchmarkTxPos, nil
}

// CheckConverter -
func CheckConverter(t *testing.T, scenariosTestPath string) {
	stateAndBenchmarkInfo, err := exporter.GetAccountsAndTransactionsFromScenarios(scenariosTestPath)
	require.Nil(t, err)
	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{}, 1)
	require.Nil(t, err)
	err = CreateAccountsFromScenariosAccs(testContext, stateAndBenchmarkInfo.Accs)
	require.Nil(t, err)
	CheckAccounts(t, testContext.Accounts, stateAndBenchmarkInfo.Accs)
	transactions := CreateTransactionsFromScenariosTxs(stateAndBenchmarkInfo.Txs)
	CheckTransactions(t, transactions, stateAndBenchmarkInfo.Txs)
}

// ProcessAllTransactions -
func ProcessAllTransactions(testContext *vm.VMTestContext, transactions []*transaction.Transaction) error {
	for _, tx := range transactions {
		sndrAccHandler, err := testContext.Accounts.LoadAccount(tx.SndAddr)
		if err != nil {
			return err
		}
		sndrAcc := sndrAccHandler.(state.UserAccountHandler)
		tx.Nonce = sndrAcc.GetNonce()
		returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
		if err != nil {
			return err
		} else if returnCode != vmcommon.Ok {
			return errReturnCodeNotOk
		}
	}
	return nil
}

// RunSingleTransactionBenchmark receives the VMTestContext (which can be created with SetStateFromScenariosTest), a tx and performs a benchmark on that specific tx. If processing transaction fails, it will return error, else will return nil
func RunSingleTransactionBenchmark(b *testing.B, testContext *vm.VMTestContext, tx *transaction.Transaction) error {
	var returnCode vmcommon.ReturnCode
	var err error
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		returnCode, err = testContext.TxProcessor.ProcessTransaction(tx)
		tx.Nonce++
	}
	b.StopTimer()
	if err != nil {
		return err
	}
	if returnCode != vmcommon.Ok {
		return errReturnCodeNotOk
	}
	return nil
}

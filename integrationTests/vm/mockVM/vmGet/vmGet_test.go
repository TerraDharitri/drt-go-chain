package vmGet

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"sync"
	"testing"

	vmData "github.com/TerraDharitri/drt-go-chain-core/data/vm"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/mock"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/process/factory"
	"github.com/TerraDharitri/drt-go-chain/process/smartContract"
	"github.com/TerraDharitri/drt-go-chain/process/sync/disabled"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/dblookupext"
	"github.com/TerraDharitri/drt-go-chain/testscommon/economicsmocks"
	"github.com/TerraDharitri/drt-go-chain/testscommon/marshallerMock"
	storageStubs "github.com/TerraDharitri/drt-go-chain/testscommon/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVmGetShouldReturnValue(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	accnts, destinationAddressBytes, expectedValueForVar := deploySmartContract(t)

	mockVM := vm.CreateOneSCExecutorMockVM(accnts)
	vmContainer := &mock.VMContainerMock{
		GetCalled: func(key []byte) (handler vmcommon.VMExecutionHandler, e error) {
			return mockVM, nil
		}}
	argsNewSCQueryService := smartContract.ArgsNewSCQueryService{
		VmContainer: vmContainer,
		EconomicsFee: &economicsmocks.EconomicsHandlerMock{
			MaxGasLimitPerBlockCalled: func(_ uint32) uint64 {
				return uint64(math.MaxUint64)
			},
		},
		BlockChainHook:           &testscommon.BlockChainHookStub{},
		MainBlockChain:           &testscommon.ChainHandlerStub{},
		APIBlockChain:            &testscommon.ChainHandlerStub{},
		WasmVMChangeLocker:       &sync.RWMutex{},
		Bootstrapper:             disabled.NewDisabledBootstrapper(),
		AllowExternalQueriesChan: common.GetClosedUnbufferedChannel(),
		HistoryRepository:        &dblookupext.HistoryRepositoryStub{},
		ShardCoordinator:         testscommon.NewMultiShardsCoordinatorMock(1),
		StorageService:           &storageStubs.ChainStorerStub{},
		Marshaller:               &marshallerMock.MarshalizerStub{},
		Hasher:                   &testscommon.HasherStub{},
		Uint64ByteSliceConverter: &mock.Uint64ByteSliceConverterMock{},
	}
	service, _ := smartContract.NewSCQueryService(argsNewSCQueryService)

	functionName := "Get"
	query := process.SCQuery{
		ScAddress: destinationAddressBytes,
		FuncName:  functionName,
		Arguments: [][]byte{},
	}

	vmOutput, _, err := service.ExecuteQuery(&query)
	assert.Nil(t, err)

	returnData, _ := vmOutput.GetFirstReturnData(vmData.AsBigInt)
	assert.Equal(t, expectedValueForVar, returnData)
}

func deploySmartContract(t *testing.T) (state.AccountsAdapter, []byte, *big.Int) {
	vmOpGas := uint64(1)
	senderAddressBytes := []byte("12345678901234567890123456789012")
	senderNonce := uint64(11)
	senderBalance := big.NewInt(100000000)
	gasPrice := uint64(1)
	gasLimit := vmOpGas + 1000
	transferOnCalls := big.NewInt(0)

	initialValueForInternalVariable := uint64(45)
	scCode := fmt.Sprintf("aaaa@%s@0000@%X", hex.EncodeToString(factory.InternalTestingVM), initialValueForInternalVariable)

	tx := vm.CreateTx(
		senderAddressBytes,
		vm.CreateEmptyAddress(),
		senderNonce,
		transferOnCalls,
		gasPrice,
		gasLimit,
		scCode,
	)

	txProc, accnts, err := vm.CreatePreparedTxProcessorAndAccountsWithMockedVM(
		vmOpGas,
		senderNonce,
		senderAddressBytes,
		senderBalance,
		config.EnableEpochs{},
		&sync.RWMutex{},
	)
	require.Nil(t, err)

	_, err = txProc.ProcessTransaction(tx)
	assert.Nil(t, err)

	_, err = accnts.Commit()
	assert.Nil(t, err)

	destinationAddressBytes, _ := hex.DecodeString("0000000000000000ffff1a2983b179a480a60c4308da48f13b4480dbb4d33132")
	return accnts, destinationAddressBytes, big.NewInt(0).SetUint64(initialValueForInternalVariable)
}

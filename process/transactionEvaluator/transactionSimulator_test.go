package transactionEvaluator

import (
	"encoding/hex"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/TerraDharitri/drt-go-chain-core/data/receipt"
	"github.com/TerraDharitri/drt-go-chain-core/data/smartContractResult"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	datafield "github.com/TerraDharitri/drt-go-chain-vm-common/parsers/dataField"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/process/mock"
	"github.com/TerraDharitri/drt-go-chain/storage/storageunit"
	"github.com/TerraDharitri/drt-go-chain/storage/txcache"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/hashingMocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTransactionSimulator(t *testing.T) {
	tests := []struct {
		name     string
		argsFunc func() ArgsTxSimulator
		exError  error
	}{
		{
			name: "NilShardCoordinator",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.ShardCoordinator = nil
				return args
			},
			exError: ErrNilShardCoordinator,
		},
		{
			name: "NilTransactionProcessor",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.TransactionProcessor = nil
				return args
			},
			exError: ErrNilTxSimulatorProcessor,
		},
		{
			name: "NilIntermProcessorContainer",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.IntermediateProcContainer = nil
				return args
			},
			exError: ErrNilIntermediateProcessorContainer,
		},
		{
			name: "NilPubkeyConverter",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.AddressPubKeyConverter = nil
				return args
			},
			exError: ErrNilPubkeyConverter,
		},
		{
			name: "NilHasher",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.Hasher = nil
				return args
			},
			exError: ErrNilHasher,
		},
		{
			name: "NilBlockChainHook",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.BlockChainHook = nil
				return args
			},
			exError: process.ErrNilBlockChainHook,
		},
		{
			name: "NilMarshalizer",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.Marshalizer = nil
				return args
			},
			exError: ErrNilMarshalizer,
		},
		{
			name: "NilVMOutputCacher",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				args.VMOutputCacher = nil
				return args
			},
			exError: ErrNilCacher,
		},
		{
			name: "Ok",
			argsFunc: func() ArgsTxSimulator {
				args := getTxSimulatorArgs()
				return args
			},
			exError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTransactionSimulator(tt.argsFunc())
			require.Equal(t, err, tt.exError)
		})
	}
}

func TestTransactionSimulator_ProcessTxProcessingErrShouldSignal(t *testing.T) {
	t.Parallel()

	expRetCode := vmcommon.Ok
	expErr := errors.New("transaction failed")
	args := getTxSimulatorArgs()
	args.TransactionProcessor = &testscommon.TxProcessorStub{
		ProcessTransactionCalled: func(transaction *transaction.Transaction) (vmcommon.ReturnCode, error) {
			return expRetCode, expErr
		},
	}
	ts, _ := NewTransactionSimulator(args)

	results, err := ts.ProcessTx(&transaction.Transaction{Nonce: 37}, &block.Header{})
	require.NoError(t, err)
	require.Equal(t, expErr.Error(), results.FailReason)
}

func TestTransactionSimulator_getVMOutputComputeHashFails(t *testing.T) {
	t.Parallel()

	args := getTxSimulatorArgs()
	args.Marshalizer = &mock.MarshalizerMock{
		Fail: true,
	}
	ts, _ := NewTransactionSimulator(args)
	require.False(t, ts.IsInterfaceNil())

	_, ok := ts.getVMOutputOfTx(nil)
	require.False(t, ok)
}

func TestTransactionSimulator_getVMOutput(t *testing.T) {
	t.Parallel()

	args := getTxSimulatorArgs()
	args.VMOutputCacher, _ = storageunit.NewCache(storageunit.CacheConfig{
		Type:     storageunit.LRUCache,
		Capacity: 100,
	})

	ts, _ := NewTransactionSimulator(args)
	require.False(t, ts.IsInterfaceNil())

	// cannot find vm output in cacher
	_, ok := ts.getVMOutputOfTx(&transaction.Transaction{})
	require.False(t, ok)

	// wrong data in cacher
	tx := &transaction.Transaction{}
	txHash, _ := core.CalculateHash(args.Marshalizer, args.Hasher, tx)
	args.VMOutputCacher.Put(txHash, 1, 0)
	_, ok = ts.getVMOutputOfTx(tx)
	require.False(t, ok)
}

func TestTransactionSimulator_ProcessTxShouldIncludeScrsAndReceipts(t *testing.T) {
	t.Parallel()

	expectedSCr := map[string]data.TransactionHandler{
		"keySCr": &smartContractResult.SmartContractResult{
			RcvAddr:        []byte("rcvr"),
			RelayerAddr:    []byte("relayer"),
			OriginalSender: []byte("original"),
		},
	}
	expectedReceipts := map[string]data.TransactionHandler{
		"keyReceipt": &receipt.Receipt{SndAddr: []byte("sndr")},
	}

	args := getTxSimulatorArgs()
	args.VMOutputCacher, _ = storageunit.NewCache(storageunit.CacheConfig{
		Type:     storageunit.LRUCache,
		Capacity: 100,
	})

	args.IntermediateProcContainer = &mock.IntermProcessorContainerStub{
		GetCalled: func(key block.Type) (process.IntermediateTransactionHandler, error) {
			return &mock.IntermediateTransactionHandlerStub{
				GetAllCurrentFinishedTxsCalled: func() map[string]data.TransactionHandler {
					if key == block.SmartContractResultBlock {
						return expectedSCr
					}
					return expectedReceipts
				},
			}, nil
		},
		KeysCalled: nil,
	}
	ts, _ := NewTransactionSimulator(args)

	tx := &transaction.Transaction{Nonce: 37}
	txHash, _ := core.CalculateHash(args.Marshalizer, args.Hasher, tx)
	args.VMOutputCacher.Put(txHash, &vmcommon.VMOutput{}, 0)

	results, err := ts.ProcessTx(tx, &block.Header{})
	require.NoError(t, err)
	require.Equal(
		t,
		hex.EncodeToString(expectedSCr["keySCr"].GetRcvAddr()),
		results.ScResults[hex.EncodeToString([]byte("keySCr"))].RcvAddr,
	)
	require.Equal(
		t,
		hex.EncodeToString(expectedReceipts["keyReceipt"].GetSndAddr()),
		results.Receipts[hex.EncodeToString([]byte("keyReceipt"))].SndAddr,
	)
}

func getTxSimulatorArgs() ArgsTxSimulator {
	pubKeyConverter := testscommon.NewPubkeyConverterMock(32)
	dataFieldParser, _ := datafield.NewOperationDataFieldParser(&datafield.ArgsOperationDataFieldParser{
		AddressLength: pubKeyConverter.Len(),
		Marshalizer:   &mock.MarshalizerMock{},
	})
	return ArgsTxSimulator{
		TransactionProcessor:      &testscommon.TxProcessorStub{},
		IntermediateProcContainer: &mock.IntermProcessorContainerStub{},
		AddressPubKeyConverter:    pubKeyConverter,
		ShardCoordinator:          mock.NewMultiShardsCoordinatorMock(2),
		VMOutputCacher:            txcache.NewDisabledCache(),
		Marshalizer:               &mock.MarshalizerMock{},
		Hasher:                    &hashingMocks.HasherMock{},
		DataFieldParser:           dataFieldParser,
		BlockChainHook:            &testscommon.BlockChainHookStub{},
	}
}

func TestTransactionSimulator_ProcessTxConcurrentCalls(t *testing.T) {
	t.Parallel()

	numTransactionProcessorCalls := 0
	args := getTxSimulatorArgs()
	args.TransactionProcessor = &testscommon.TxProcessorStub{
		ProcessTransactionCalled: func(transaction *transaction.Transaction) (vmcommon.ReturnCode, error) {
			// deliberately not used a mutex here as to catch race conditions
			numTransactionProcessorCalls++

			return vmcommon.Ok, nil
		},
	}
	txSimulator, _ := NewTransactionSimulator(args)
	tx := &transaction.Transaction{Nonce: 37}

	numCalls := 100
	wg := sync.WaitGroup{}
	wg.Add(numCalls)
	for i := 0; i < numCalls; i++ {
		go func(idx int) {
			time.Sleep(time.Millisecond * 10)
			_, _ = txSimulator.ProcessTx(tx, &block.Header{})
			wg.Done()
		}(i)
	}

	wg.Wait()
	assert.Equal(t, numCalls, numTransactionProcessorCalls)
}

package shard

import (
	"runtime"
	"sync"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	vmcommonBuiltInFunctions "github.com/TerraDharitri/drt-go-chain-vm-common/builtInFunctions"
	"github.com/TerraDharitri/drt-go-chain-vm-common/parsers"
	ipcNodePart1p2 "github.com/TerraDharitri/drt-go-chain-vm-v1/ipc/nodepart"
	wasmConfig "github.com/TerraDharitri/drt-go-chain-vm/config"
	"github.com/TerraDharitri/drt-go-chain/common/forking"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/process/factory"
	"github.com/TerraDharitri/drt-go-chain/process/mock"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/enableEpochsHandlerMock"
	"github.com/TerraDharitri/drt-go-chain/testscommon/epochNotifier"
	"github.com/TerraDharitri/drt-go-chain/testscommon/hashingMocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeVMConfig() config.VirtualMachineConfig {
	return config.VirtualMachineConfig{
		WasmVMVersions: []config.WasmVMVersionByEpoch{
			{StartEpoch: 0, Version: "v1.2"},
			{StartEpoch: 10, Version: "v1.2"},
			{StartEpoch: 12, Version: "v1.3"},
			{StartEpoch: 14, Version: "v1.4"},
		},
		TransferAndExecuteByUserAddresses: []string{"drt1qqqqqqqqqqqqqpgqr46jrxr6r2unaqh75ugd308dwx5vgnhwh47qkswz60"},
	}
}

func createMockVMAccountsArguments() ArgVMContainerFactory {
	dcdtTransferParser, _ := parsers.NewDCDTTransferParser(&mock.MarshalizerMock{})
	return ArgVMContainerFactory{
		Config:              makeVMConfig(),
		BlockGasLimit:       10000,
		GasSchedule:         testscommon.NewGasScheduleNotifierMock(wasmConfig.MakeGasMapForTests()),
		EpochNotifier:       &epochNotifier.EpochNotifierStub{},
		EnableEpochsHandler: &enableEpochsHandlerMock.EnableEpochsHandlerStub{},
		WasmVMChangeLocker:  &sync.RWMutex{},
		DCDTTransferParser:  dcdtTransferParser,
		BuiltInFunctions:    vmcommonBuiltInFunctions.NewBuiltInFunctionContainer(),
		BlockChainHook:      &testscommon.BlockChainHookStub{},
		Hasher:              &hashingMocks.HasherMock{},
		PubKeyConverter:     testscommon.RealWorldBech32PubkeyConverter,
	}
}

func TestNewVMContainerFactory_NilGasScheduleShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockVMAccountsArguments()
	args.GasSchedule = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilGasSchedule, err)
}

func TestNewVMContainerFactory_NilDCDTTransferParserShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockVMAccountsArguments()
	args.DCDTTransferParser = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilDCDTTransferParser, err)
}

func TestNewVMContainerFactory_NilLockerShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockVMAccountsArguments()
	args.WasmVMChangeLocker = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilLocker, err)
}

func TestNewVMContainerFactory_NilEpochNotifierShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockVMAccountsArguments()
	args.EpochNotifier = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilEpochNotifier, err)
}

func TestNewVMContainerFactory_NilEnableEpochsHandlerShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockVMAccountsArguments()
	args.EnableEpochsHandler = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilEnableEpochsHandler, err)
}

func TestNewVMContainerFactory_NilBuiltinFunctionsShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockVMAccountsArguments()
	args.BuiltInFunctions = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilBuiltInFunction, err)
}

func TestNewVMContainerFactory_NilBlockChainHookShouldErr(t *testing.T) {
	t.Parallel()

	args := createMockVMAccountsArguments()
	args.BlockChainHook = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilBlockChainHook, err)
}

func TestNewVMContainerFactory_NilHasherShouldErr(t *testing.T) {
	args := createMockVMAccountsArguments()
	args.Hasher = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilHasher, err)
}

func TestNewVMContainerFactory_NilPubKeyConverterShouldErr(t *testing.T) {
	args := createMockVMAccountsArguments()
	args.PubKeyConverter = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrNilPubkeyConverter, err)
}

func TestNewVMContainerFactory_EmptyOpcodeAddressListErr(t *testing.T) {
	args := createMockVMAccountsArguments()
	args.Config.TransferAndExecuteByUserAddresses = nil
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, process.ErrTransferAndExecuteByUserAddressesAreNil, err)
}

func TestNewVMContainerFactory_WrongAddressErr(t *testing.T) {
	args := createMockVMAccountsArguments()
	args.Config.TransferAndExecuteByUserAddresses = []string{"just"}
	vmf, err := NewVMContainerFactory(args)

	assert.Nil(t, vmf)
	assert.Equal(t, err.Error(), "invalid bech32 string length 4")
}

func TestNewVMContainerFactory_OkValues(t *testing.T) {
	if runtime.GOARCH == "arm64" {
		t.Skip("skipping test on arm64")
	}

	args := createMockVMAccountsArguments()
	vmf, err := NewVMContainerFactory(args)

	assert.NotNil(t, vmf)
	assert.Nil(t, err)
	assert.False(t, vmf.IsInterfaceNil())
}

func TestVmContainerFactory_Create(t *testing.T) {
	if runtime.GOARCH == "arm64" {
		t.Skip("skipping test on arm64")
	}

	args := createMockVMAccountsArguments()
	vmf, _ := NewVMContainerFactory(args)
	require.NotNil(t, vmf)

	container, err := vmf.Create()
	require.Nil(t, err)
	require.NotNil(t, container)
	defer func() {
		_ = container.Close()
	}()

	assert.Nil(t, err)
	assert.NotNil(t, container)

	vm, err := container.Get(factory.WasmVirtualMachine)
	assert.Nil(t, err)
	assert.NotNil(t, vm)

	acc := vmf.BlockChainHookImpl()
	assert.NotNil(t, acc)

	assert.Equal(t, len(vmf.mapOpcodeAddressIsAllowed), 1)
	assert.Equal(t, len(vmf.mapOpcodeAddressIsAllowed[managedMultiTransferDCDTNFTExecuteByUser]), 1)
}

func TestVmContainerFactory_ResolveWasmVMVersion(t *testing.T) {
	if runtime.GOARCH == "arm64" {
		t.Skip("skipping test on arm64")
	}

	epochNotifierInstance := forking.NewGenericEpochNotifier()

	numCalled := 0
	gasScheduleNotifier := testscommon.NewGasScheduleNotifierMock(wasmConfig.MakeGasMapForTests())
	gasScheduleNotifier.RegisterNotifyHandlerCalled = func(handler core.GasScheduleSubscribeHandler) {
		numCalled++
		handler.GasScheduleChange(gasScheduleNotifier.GasSchedule)
	}
	args := createMockVMAccountsArguments()
	args.GasSchedule = gasScheduleNotifier
	args.EpochNotifier = epochNotifierInstance
	vmf, _ := NewVMContainerFactory(args)
	require.NotNil(t, vmf)

	container, err := vmf.Create()
	require.Nil(t, err)
	require.NotNil(t, container)
	defer func() {
		_ = container.Close()
	}()
	require.Equal(t, "v1.2", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	epochNotifierInstance.CheckEpoch(makeHeaderHandlerStub(1))
	require.Equal(t, "v1.2", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	epochNotifierInstance.CheckEpoch(makeHeaderHandlerStub(6))
	require.Equal(t, "v1.2", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	epochNotifierInstance.CheckEpoch(makeHeaderHandlerStub(10))
	require.Equal(t, "v1.2", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	epochNotifierInstance.CheckEpoch(makeHeaderHandlerStub(11))
	require.Equal(t, "v1.2", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	epochNotifierInstance.CheckEpoch(makeHeaderHandlerStub(12))
	require.Equal(t, "v1.3", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	epochNotifierInstance.CheckEpoch(makeHeaderHandlerStub(13))
	require.Equal(t, "v1.3", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	epochNotifierInstance.CheckEpoch(makeHeaderHandlerStub(20))
	require.Equal(t, "v1.4", getWasmVMVersion(t, container))
	require.False(t, isOutOfProcess(t, container))

	require.Equal(t, numCalled, 1)
}

func makeHeaderHandlerStub(epoch uint32) data.HeaderHandler {
	return &testscommon.HeaderHandlerStub{
		EpochField: epoch,
	}
}

func isOutOfProcess(t testing.TB, container process.VirtualMachinesContainer) bool {
	vm, err := container.Get(factory.WasmVirtualMachine)
	require.Nil(t, err)
	require.NotNil(t, vm)

	_, ok := vm.(*ipcNodePart1p2.VMDriver)
	return ok
}

func getWasmVMVersion(t testing.TB, container process.VirtualMachinesContainer) string {
	vm, err := container.Get(factory.WasmVirtualMachine)
	require.Nil(t, err)
	require.NotNil(t, vm)

	return vm.GetVersion()
}

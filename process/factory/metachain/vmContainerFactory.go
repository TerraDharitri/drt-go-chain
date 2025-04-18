package metachain

import (
	"fmt"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/hashing"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain-vm-common/parsers"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/process/factory"
	"github.com/TerraDharitri/drt-go-chain/process/factory/containers"
	"github.com/TerraDharitri/drt-go-chain/process/smartContract/hooks"
	"github.com/TerraDharitri/drt-go-chain/sharding"
	"github.com/TerraDharitri/drt-go-chain/sharding/nodesCoordinator"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain/vm"
	systemVMFactory "github.com/TerraDharitri/drt-go-chain/vm/factory"
	systemVMProcess "github.com/TerraDharitri/drt-go-chain/vm/process"
	"github.com/TerraDharitri/drt-go-chain/vm/systemSmartContracts"
)

var _ process.VirtualMachinesContainerFactory = (*vmContainerFactory)(nil)

type vmContainerFactory struct {
	chanceComputer         nodesCoordinator.ChanceComputer
	validatorAccountsDB    state.AccountsAdapter
	userAccountsDB         state.AccountsAdapter
	blockChainHook         process.BlockChainHookWithAccountsAdapter
	cryptoHook             vmcommon.CryptoHook
	systemContracts        vm.SystemSCContainer
	economics              process.EconomicsDataHandler
	messageSigVerifier     vm.MessageSignVerifier
	nodesConfigProvider    vm.NodesConfigProvider
	gasSchedule            core.GasScheduleNotifier
	hasher                 hashing.Hasher
	marshalizer            marshal.Marshalizer
	systemSCConfig         *config.SystemSmartContractsConfig
	addressPubKeyConverter core.PubkeyConverter
	scFactory              vm.SystemSCContainerFactory
	shardCoordinator       sharding.Coordinator
	enableEpochsHandler    common.EnableEpochsHandler
	nodesCoordinator       vm.NodesCoordinator
}

// ArgsNewVMContainerFactory defines the arguments needed to create a new VM container factory
type ArgsNewVMContainerFactory struct {
	ArgBlockChainHook   hooks.ArgBlockChainHook
	Economics           process.EconomicsDataHandler
	MessageSignVerifier vm.MessageSignVerifier
	GasSchedule         core.GasScheduleNotifier
	NodesConfigProvider vm.NodesConfigProvider
	Hasher              hashing.Hasher
	Marshalizer         marshal.Marshalizer
	SystemSCConfig      *config.SystemSmartContractsConfig
	ValidatorAccountsDB state.AccountsAdapter
	UserAccountsDB      state.AccountsAdapter
	ChanceComputer      nodesCoordinator.ChanceComputer
	ShardCoordinator    sharding.Coordinator
	PubkeyConv          core.PubkeyConverter
	BlockChainHook      process.BlockChainHookWithAccountsAdapter
	EnableEpochsHandler common.EnableEpochsHandler
	NodesCoordinator    vm.NodesCoordinator
}

// NewVMContainerFactory is responsible for creating a new virtual machine factory object
func NewVMContainerFactory(args ArgsNewVMContainerFactory) (*vmContainerFactory, error) {
	if check.IfNil(args.Economics) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", process.ErrNilEconomicsData)
	}
	if check.IfNil(args.MessageSignVerifier) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", vm.ErrNilMessageSignVerifier)
	}
	if check.IfNil(args.NodesConfigProvider) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", process.ErrNilNodesConfigProvider)
	}
	if check.IfNil(args.Hasher) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", process.ErrNilHasher)
	}
	if check.IfNil(args.Marshalizer) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", process.ErrNilMarshalizer)
	}
	if args.SystemSCConfig == nil {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", process.ErrNilSystemSCConfig)
	}
	if check.IfNil(args.ValidatorAccountsDB) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", vm.ErrNilValidatorAccountsDB)
	}
	if check.IfNil(args.UserAccountsDB) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", vm.ErrNilUserAccountsDB)
	}
	if check.IfNil(args.ChanceComputer) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", vm.ErrNilChanceComputer)
	}
	if check.IfNil(args.GasSchedule) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", vm.ErrNilGasSchedule)
	}
	if check.IfNil(args.PubkeyConv) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", vm.ErrNilAddressPubKeyConverter)
	}
	if check.IfNil(args.ShardCoordinator) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", vm.ErrNilShardCoordinator)
	}
	if check.IfNil(args.BlockChainHook) {
		return nil, process.ErrNilBlockChainHook
	}
	if check.IfNil(args.EnableEpochsHandler) {
		return nil, vm.ErrNilEnableEpochsHandler
	}
	if check.IfNil(args.NodesCoordinator) {
		return nil, fmt.Errorf("%w in NewVMContainerFactory", process.ErrNilNodesCoordinator)
	}

	cryptoHook := hooks.NewVMCryptoHook()

	return &vmContainerFactory{
		blockChainHook:         args.BlockChainHook,
		cryptoHook:             cryptoHook,
		economics:              args.Economics,
		messageSigVerifier:     args.MessageSignVerifier,
		gasSchedule:            args.GasSchedule,
		nodesConfigProvider:    args.NodesConfigProvider,
		hasher:                 args.Hasher,
		marshalizer:            args.Marshalizer,
		systemSCConfig:         args.SystemSCConfig,
		validatorAccountsDB:    args.ValidatorAccountsDB,
		userAccountsDB:         args.UserAccountsDB,
		chanceComputer:         args.ChanceComputer,
		addressPubKeyConverter: args.PubkeyConv,
		shardCoordinator:       args.ShardCoordinator,
		enableEpochsHandler:    args.EnableEpochsHandler,
		nodesCoordinator:       args.NodesCoordinator,
	}, nil
}

// Create sets up all the needed virtual machines returning a container of all the VMs
func (vmf *vmContainerFactory) Create() (process.VirtualMachinesContainer, error) {
	container := containers.NewVirtualMachinesContainer()

	currVm, err := vmf.createSystemVM()
	if err != nil {
		return nil, err
	}

	err = container.Add(factory.SystemVirtualMachine, currVm)
	if err != nil {
		return nil, err
	}

	return container, nil
}

// Close closes the vm container factory
func (vmf *vmContainerFactory) Close() error {
	return vmf.blockChainHook.Close()
}

// CreateForGenesis sets up all the needed virtual machines returning a container of all the VMs to be used in the genesis process
// The system VM will have to contain the following and only following system smartcontracts:
// staking SC, validator SC, DCDT SC and governance SC. Including more system smartcontracts (or less) will trigger root hash mismatch
// errors when trying to sync the first metablock after the genesis event.
func (vmf *vmContainerFactory) CreateForGenesis() (process.VirtualMachinesContainer, error) {
	container := containers.NewVirtualMachinesContainer()

	currVm, err := vmf.createSystemVMForGenesis()
	if err != nil {
		return nil, err
	}

	err = container.Add(factory.SystemVirtualMachine, currVm)
	if err != nil {
		return nil, err
	}

	return container, nil
}

func (vmf *vmContainerFactory) createSystemVMFactoryAndEEI() (vm.SystemSCContainerFactory, vm.ContextHandler, error) {
	atArgumentParser := parsers.NewCallArgsParser()
	vmContextArgs := systemSmartContracts.VMContextArgs{
		BlockChainHook:      vmf.blockChainHook,
		CryptoHook:          vmf.cryptoHook,
		InputParser:         atArgumentParser,
		ValidatorAccountsDB: vmf.validatorAccountsDB,
		UserAccountsDB:      vmf.userAccountsDB,
		ChanceComputer:      vmf.chanceComputer,
		EnableEpochsHandler: vmf.enableEpochsHandler,
	}
	systemEI, err := systemSmartContracts.NewVMContext(vmContextArgs)
	if err != nil {
		return nil, nil, err
	}

	argsNewSystemScFactory := systemVMFactory.ArgsNewSystemSCFactory{
		SystemEI:               systemEI,
		SigVerifier:            vmf.messageSigVerifier,
		GasSchedule:            vmf.gasSchedule,
		NodesConfigProvider:    vmf.nodesConfigProvider,
		Hasher:                 vmf.hasher,
		Marshalizer:            vmf.marshalizer,
		SystemSCConfig:         vmf.systemSCConfig,
		Economics:              vmf.economics,
		AddressPubKeyConverter: vmf.addressPubKeyConverter,
		ShardCoordinator:       vmf.shardCoordinator,
		EnableEpochsHandler:    vmf.enableEpochsHandler,
		NodesCoordinator:       vmf.nodesCoordinator,
	}
	scFactory, err := systemVMFactory.NewSystemSCFactory(argsNewSystemScFactory)
	if err != nil {
		return nil, nil, err
	}

	return scFactory, systemEI, nil
}

func (vmf *vmContainerFactory) finalizeSystemVMCreation(systemEI vm.ContextHandler) (vmcommon.VMExecutionHandler, error) {
	err := systemEI.SetSystemSCContainer(vmf.systemContracts)
	if err != nil {
		return nil, err
	}

	argsNewSystemVM := systemVMProcess.ArgsNewSystemVM{
		SystemEI:        systemEI,
		SystemContracts: vmf.systemContracts,
		VmType:          factory.SystemVirtualMachine,
		GasSchedule:     vmf.gasSchedule,
	}
	systemVM, err := systemVMProcess.NewSystemVM(argsNewSystemVM)
	if err != nil {
		return nil, err
	}

	vmf.gasSchedule.RegisterNotifyHandler(systemVM)

	return systemVM, nil
}

func (vmf *vmContainerFactory) createSystemVM() (vmcommon.VMExecutionHandler, error) {
	scFactory, systemEI, err := vmf.createSystemVMFactoryAndEEI()
	if err != nil {
		return nil, err
	}

	vmf.systemContracts, err = scFactory.Create()
	if err != nil {
		return nil, err
	}

	vmf.scFactory = scFactory

	return vmf.finalizeSystemVMCreation(systemEI)
}

// createSystemVMForGenesis will create the same VMExecutionHandler structure used when the mainnet genesis was created
func (vmf *vmContainerFactory) createSystemVMForGenesis() (vmcommon.VMExecutionHandler, error) {
	scFactory, systemEI, err := vmf.createSystemVMFactoryAndEEI()
	if err != nil {
		return nil, err
	}

	vmf.systemContracts, err = scFactory.CreateForGenesis()
	if err != nil {
		return nil, err
	}

	return vmf.finalizeSystemVMCreation(systemEI)
}

// BlockChainHookImpl returns the created blockChainHookImpl
func (vmf *vmContainerFactory) BlockChainHookImpl() process.BlockChainHookWithAccountsAdapter {
	return vmf.blockChainHook
}

// SystemSmartContractContainer return the created system smart contracts
func (vmf *vmContainerFactory) SystemSmartContractContainer() vm.SystemSCContainer {
	return vmf.systemContracts
}

// SystemSmartContractContainerFactory returns the system smart contract container factory
func (vmf *vmContainerFactory) SystemSmartContractContainerFactory() vm.SystemSCContainerFactory {
	return vmf.scFactory
}

// IsInterfaceNil returns true if there is no value under the interface
func (vmf *vmContainerFactory) IsInterfaceNil() bool {
	return vmf == nil
}

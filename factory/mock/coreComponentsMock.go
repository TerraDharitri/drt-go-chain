package mock

import (
	"sync"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data/endProcess"
	"github.com/TerraDharitri/drt-go-chain-core/data/typeConverters"
	"github.com/TerraDharitri/drt-go-chain-core/hashing"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/consensus"
	"github.com/TerraDharitri/drt-go-chain/factory"
	"github.com/TerraDharitri/drt-go-chain/ntp"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/sharding"
	"github.com/TerraDharitri/drt-go-chain/sharding/nodesCoordinator"
	"github.com/TerraDharitri/drt-go-chain/storage"
)

// CoreComponentsMock -
type CoreComponentsMock struct {
	IntMarsh                     marshal.Marshalizer
	TxMarsh                      marshal.Marshalizer
	VmMarsh                      marshal.Marshalizer
	Hash                         hashing.Hasher
	TxSignHasherField            hashing.Hasher
	UInt64ByteSliceConv          typeConverters.Uint64ByteSliceConverter
	AddrPubKeyConv               core.PubkeyConverter
	ValPubKeyConv                core.PubkeyConverter
	PathHdl                      storage.PathManagerHandler
	WatchdogTimer                core.WatchdogTimer
	AlarmSch                     core.TimersScheduler
	NtpSyncTimer                 ntp.SyncTimer
	GenesisBlockTime             time.Time
	ChainIdCalled                func() string
	MinTransactionVersionCalled  func() uint32
	mutIntMarshalizer            sync.RWMutex
	RoundHandlerField            consensus.RoundHandler
	EconomicsHandler             process.EconomicsDataHandler
	APIEconomicsHandler          process.EconomicsDataHandler
	RatingsConfig                process.RatingsInfoHandler
	RatingHandler                sharding.PeerAccountListAndRatingHandler
	NodesConfig                  sharding.GenesisNodesSetupHandler
	Shuffler                     nodesCoordinator.NodesShuffler
	EpochChangeNotifier          process.EpochNotifier
	RoundChangeNotifier          process.RoundNotifier
	EnableRoundsHandlerField     process.EnableRoundsHandler
	EpochNotifierWithConfirm     factory.EpochStartNotifierWithConfirm
	TxVersionCheckHandler        process.TxVersionCheckerHandler
	ChanStopProcess              chan endProcess.ArgEndProcess
	StartTime                    time.Time
	NodeTypeProviderField        core.NodeTypeProviderHandler
	WasmVMChangeLockerInternal   common.Locker
	ProcessStatusHandlerInternal common.ProcessStatusHandler
	HardforkTriggerPubKeyField   []byte
	EnableEpochsHandlerField     common.EnableEpochsHandler
}

// InternalMarshalizer -
func (ccm *CoreComponentsMock) InternalMarshalizer() marshal.Marshalizer {
	ccm.mutIntMarshalizer.RLock()
	defer ccm.mutIntMarshalizer.RUnlock()

	return ccm.IntMarsh
}

// SetInternalMarshalizer -
func (ccm *CoreComponentsMock) SetInternalMarshalizer(m marshal.Marshalizer) error {
	ccm.mutIntMarshalizer.Lock()
	ccm.IntMarsh = m
	ccm.mutIntMarshalizer.Unlock()

	return nil
}

// TxMarshalizer -
func (ccm *CoreComponentsMock) TxMarshalizer() marshal.Marshalizer {
	return ccm.TxMarsh
}

// VmMarshalizer -
func (ccm *CoreComponentsMock) VmMarshalizer() marshal.Marshalizer {
	return ccm.VmMarsh
}

// Hasher -
func (ccm *CoreComponentsMock) Hasher() hashing.Hasher {
	return ccm.Hash
}

// TxSignHasher -
func (ccm *CoreComponentsMock) TxSignHasher() hashing.Hasher {
	return ccm.TxSignHasherField
}

// Uint64ByteSliceConverter -
func (ccm *CoreComponentsMock) Uint64ByteSliceConverter() typeConverters.Uint64ByteSliceConverter {
	return ccm.UInt64ByteSliceConv
}

// AddressPubKeyConverter -
func (ccm *CoreComponentsMock) AddressPubKeyConverter() core.PubkeyConverter {
	return ccm.AddrPubKeyConv
}

// ValidatorPubKeyConverter -
func (ccm *CoreComponentsMock) ValidatorPubKeyConverter() core.PubkeyConverter {
	return ccm.ValPubKeyConv
}

// PathHandler -
func (ccm *CoreComponentsMock) PathHandler() storage.PathManagerHandler {
	return ccm.PathHdl
}

// Watchdog -
func (ccm *CoreComponentsMock) Watchdog() core.WatchdogTimer {
	return ccm.WatchdogTimer
}

// AlarmScheduler -
func (ccm *CoreComponentsMock) AlarmScheduler() core.TimersScheduler {
	return ccm.AlarmSch
}

// SyncTimer -
func (ccm *CoreComponentsMock) SyncTimer() ntp.SyncTimer {
	return ccm.NtpSyncTimer
}

// GenesisTime -
func (ccm *CoreComponentsMock) GenesisTime() time.Time {
	return ccm.GenesisBlockTime
}

// ChainID -
func (ccm *CoreComponentsMock) ChainID() string {
	if ccm.ChainIdCalled != nil {
		return ccm.ChainIdCalled()
	}
	return "undefined"
}

// MinTransactionVersion -
func (ccm *CoreComponentsMock) MinTransactionVersion() uint32 {
	if ccm.MinTransactionVersionCalled != nil {
		return ccm.MinTransactionVersionCalled()
	}
	return 1
}

// TxVersionChecker -
func (ccm *CoreComponentsMock) TxVersionChecker() process.TxVersionCheckerHandler {
	return ccm.TxVersionCheckHandler
}

// EncodedAddressLen -
func (ccm *CoreComponentsMock) EncodedAddressLen() uint32 {
	return uint32(ccm.AddressPubKeyConverter().Len() * 2)
}

// RoundHandler -
func (ccm *CoreComponentsMock) RoundHandler() consensus.RoundHandler {
	return ccm.RoundHandlerField
}

// EconomicsData -
func (ccm *CoreComponentsMock) EconomicsData() process.EconomicsDataHandler {
	return ccm.EconomicsHandler
}

// APIEconomicsData -
func (ccm *CoreComponentsMock) APIEconomicsData() process.EconomicsDataHandler {
	return ccm.APIEconomicsHandler
}

// RatingsData -
func (ccm *CoreComponentsMock) RatingsData() process.RatingsInfoHandler {
	return ccm.RatingsConfig
}

// Rater -
func (ccm *CoreComponentsMock) Rater() sharding.PeerAccountListAndRatingHandler {
	return ccm.RatingHandler
}

// GenesisNodesSetup -
func (ccm *CoreComponentsMock) GenesisNodesSetup() sharding.GenesisNodesSetupHandler {
	return ccm.NodesConfig
}

// NodesShuffler -
func (ccm *CoreComponentsMock) NodesShuffler() nodesCoordinator.NodesShuffler {
	return ccm.Shuffler
}

// EpochNotifier -
func (ccm *CoreComponentsMock) EpochNotifier() process.EpochNotifier {
	return ccm.EpochChangeNotifier
}

// RoundNotifier -
func (ccm *CoreComponentsMock) RoundNotifier() process.RoundNotifier {
	return ccm.RoundChangeNotifier
}

// EnableRoundsHandler -
func (ccm *CoreComponentsMock) EnableRoundsHandler() process.EnableRoundsHandler {
	return ccm.EnableRoundsHandlerField
}

// EpochStartNotifierWithConfirm -
func (ccm *CoreComponentsMock) EpochStartNotifierWithConfirm() factory.EpochStartNotifierWithConfirm {
	return ccm.EpochNotifierWithConfirm
}

// ChanStopNodeProcess -
func (ccm *CoreComponentsMock) ChanStopNodeProcess() chan endProcess.ArgEndProcess {
	return ccm.ChanStopProcess
}

// NodeTypeProvider -
func (ccm *CoreComponentsMock) NodeTypeProvider() core.NodeTypeProviderHandler {
	return ccm.NodeTypeProviderField
}

// WasmVMChangeLocker -
func (ccm *CoreComponentsMock) WasmVMChangeLocker() common.Locker {
	return ccm.WasmVMChangeLockerInternal
}

// ProcessStatusHandler -
func (ccm *CoreComponentsMock) ProcessStatusHandler() common.ProcessStatusHandler {
	return ccm.ProcessStatusHandlerInternal
}

// HardforkTriggerPubKey -
func (ccm *CoreComponentsMock) HardforkTriggerPubKey() []byte {
	return ccm.HardforkTriggerPubKeyField
}

// EnableEpochsHandler -
func (ccm *CoreComponentsMock) EnableEpochsHandler() common.EnableEpochsHandler {
	return ccm.EnableEpochsHandlerField
}

// IsInterfaceNil -
func (ccm *CoreComponentsMock) IsInterfaceNil() bool {
	return ccm == nil
}

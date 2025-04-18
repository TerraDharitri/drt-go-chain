package state

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/storage/storageEpochChange"
	"github.com/TerraDharitri/drt-go-chain/trie/storageMarker"
	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/atomic"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
)

// ArgsNewSnapshotsManager are the args needed for creating a new snapshots manager
type ArgsNewSnapshotsManager struct {
	ShouldSerializeSnapshots bool
	ProcessingMode           common.NodeProcessingMode
	Marshaller               marshal.Marshalizer
	AddressConverter         core.PubkeyConverter
	ProcessStatusHandler     common.ProcessStatusHandler
	StateMetrics             StateMetrics
	AccountFactory           AccountFactory
	ChannelsProvider         IteratorChannelsProvider
	StateStatsHandler        StateStatsHandler
	LastSnapshotMarker       LastSnapshotMarker
}

type snapshotsManager struct {
	isSnapshotInProgress     atomic.Flag
	lastSnapshot             *snapshotInfo
	shouldSerializeSnapshots bool
	processingMode           common.NodeProcessingMode

	stateMetrics         StateMetrics
	lastSnapshotMarker   LastSnapshotMarker
	marshaller           marshal.Marshalizer
	addressConverter     core.PubkeyConverter
	trieSyncer           AccountsDBSyncer
	processStatusHandler common.ProcessStatusHandler
	channelsProvider     IteratorChannelsProvider
	accountFactory       AccountFactory
	stateStatsHandler    StateStatsHandler
	mutex                sync.RWMutex
}

// NewSnapshotsManager creates a new snapshots manager
func NewSnapshotsManager(args ArgsNewSnapshotsManager) (*snapshotsManager, error) {
	if check.IfNil(args.Marshaller) {
		return nil, ErrNilMarshalizer
	}
	if check.IfNil(args.ProcessStatusHandler) {
		return nil, ErrNilProcessStatusHandler
	}
	if check.IfNil(args.AddressConverter) {
		return nil, ErrNilAddressConverter
	}
	if check.IfNil(args.StateMetrics) {
		return nil, ErrNilStateMetrics
	}
	if check.IfNil(args.ChannelsProvider) {
		return nil, ErrNilChannelsProvider
	}
	if check.IfNil(args.AccountFactory) {
		return nil, ErrNilAccountFactory
	}
	if check.IfNil(args.StateStatsHandler) {
		return nil, ErrNilStatsHandler
	}
	if check.IfNil(args.LastSnapshotMarker) {
		return nil, ErrNilLastSnapshotMarker
	}

	return &snapshotsManager{
		isSnapshotInProgress:     atomic.Flag{},
		lastSnapshot:             &snapshotInfo{},
		shouldSerializeSnapshots: args.ShouldSerializeSnapshots,
		processingMode:           args.ProcessingMode,
		stateMetrics:             args.StateMetrics,
		marshaller:               args.Marshaller,
		addressConverter:         args.AddressConverter,
		trieSyncer:               nil, //this will be set later
		processStatusHandler:     args.ProcessStatusHandler,
		channelsProvider:         args.ChannelsProvider,
		mutex:                    sync.RWMutex{},
		accountFactory:           args.AccountFactory,
		stateStatsHandler:        args.StateStatsHandler,
		lastSnapshotMarker:       args.LastSnapshotMarker,
	}, nil
}

// SetSyncer sets the given syncer as the syncer for the underlying trie
func (sm *snapshotsManager) SetSyncer(syncer AccountsDBSyncer) error {
	if check.IfNil(syncer) {
		return ErrNilTrieSyncer
	}

	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sm.trieSyncer = syncer
	return nil
}

// IsSnapshotInProgress returns true if a snapshot is in progress
func (sm *snapshotsManager) IsSnapshotInProgress() bool {
	return sm.isSnapshotInProgress.IsSet()
}

func (sm *snapshotsManager) getTrieSyncer() AccountsDBSyncer {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	return sm.trieSyncer
}

// StartSnapshotAfterRestartIfNeeded starts a snapshot after restart if needed
func (sm *snapshotsManager) StartSnapshotAfterRestartIfNeeded(trieStorageManager common.StorageManager) error {
	if check.IfNil(sm.getTrieSyncer()) {
		return ErrNilTrieSyncer
	}

	if !trieStorageManager.ShouldTakeSnapshot() {
		return nil
	}

	rootHash, epoch, err := sm.getSnapshotRootHashAndEpoch(trieStorageManager)
	if err != nil {
		log.Warn("could not retrieve snapshot info", "error", err)
		return nil
	}

	log.Debug("startSnapshotAfterRestart", "rootHash", rootHash, "epoch", epoch)
	sm.SnapshotState(rootHash, epoch, trieStorageManager)

	return nil
}

func (sm *snapshotsManager) getSnapshotRootHashAndEpoch(trieStorageManager common.StorageManager) ([]byte, uint32, error) {
	rootHash, err := sm.lastSnapshotMarker.GetMarkerInfo(trieStorageManager)
	if err != nil {
		return nil, 0, err
	}

	epoch, err := trieStorageManager.GetLatestStorageEpoch()
	if err != nil {
		return nil, 0, err
	}

	return rootHash, epoch, nil
}

// SnapshotState starts a state snapshot
func (sm *snapshotsManager) SnapshotState(
	rootHash []byte,
	epoch uint32,
	trieStorageManager common.StorageManager,
) {
	if check.IfNil(trieStorageManager) {
		return
	}
	if !trieStorageManager.IsSnapshotSupported() {
		log.Debug("skipping snapshot as the snapshot is not supported by the current trieStorageManager",
			"trieStorageManager type", fmt.Sprintf("%T", trieStorageManager))
		return
	}

	sm.mutex.Lock()

	stats, skipSnapshot := sm.prepareSnapshot(rootHash, epoch, trieStorageManager)
	if skipSnapshot {
		log.Debug("skipping snapshot",
			"last snapshot rootHash", sm.lastSnapshot.rootHash,
			"rootHash", rootHash,
			"last snapshot epoch", sm.lastSnapshot.epoch,
			"epoch", epoch,
			"isSnapshotInProgress", sm.isSnapshotInProgress.IsSet(),
		)

		sm.mutex.Unlock()
		return
	}
	sm.mutex.Unlock()

	log.Info("starting snapshot", "type", sm.stateMetrics.GetSnapshotMessage(), "rootHash", rootHash, "epoch", epoch)

	go sm.snapshotState(rootHash, epoch, trieStorageManager, stats)

	sm.waitForCompletionIfAppropriate(stats)
}

func (sm *snapshotsManager) prepareSnapshot(rootHash []byte, epoch uint32, trieStorageManager common.StorageManager) (*snapshotStatistics, bool) {
	snapshotAlreadyTaken := bytes.Equal(sm.lastSnapshot.rootHash, rootHash) && sm.lastSnapshot.epoch == epoch
	if snapshotAlreadyTaken {
		return nil, true
	}

	if sm.processingMode != common.ImportDb {
		go sm.lastSnapshotMarker.AddMarker(trieStorageManager, epoch, rootHash)
	}

	if sm.isSnapshotInProgress.IsSet() {
		return nil, true
	}

	sm.isSnapshotInProgress.SetValue(true)
	sm.lastSnapshot.rootHash = rootHash
	sm.lastSnapshot.epoch = epoch
	trieStorageManager.EnterPruningBufferingMode()
	stats := newSnapshotStatistics(1, 1)

	sm.stateStatsHandler.ResetSnapshot()

	return stats, false
}

func (sm *snapshotsManager) snapshotState(
	rootHash []byte,
	epoch uint32,
	trieStorageManager common.StorageManager,
	stats *snapshotStatistics,
) {
	if sm.processingMode != common.ImportDb {
		err := storageEpochChange.WaitForStorageEpochChange(storageEpochChange.StorageEpochChangeWaitArgs{
			TrieStorageManager:            trieStorageManager,
			Epoch:                         epoch,
			WaitTimeForSnapshotEpochCheck: storageEpochChange.WaitTimeForSnapshotEpochCheck,
			SnapshotWaitTimeout:           storageEpochChange.SnapshotWaitTimeout,
		})
		if err != nil {
			log.Error("error waiting for storage epoch change", "err", err)
			sm.earlySnapshotCompletion(stats, trieStorageManager)
			return
		}
	}

	if !trieStorageManager.ShouldTakeSnapshot() {
		log.Debug("skipping snapshot", "rootHash", rootHash, "epoch", epoch)

		sm.earlySnapshotCompletion(stats, trieStorageManager)
		return
	}

	missingNodesChannel := make(chan []byte, missingNodesChannelSize)
	iteratorChannels := sm.channelsProvider.GetIteratorChannels()

	sm.stateMetrics.UpdateMetricsOnSnapshotStart()

	go func() {
		stats.NewSnapshotStarted()

		trieStorageManager.TakeSnapshot("", rootHash, rootHash, iteratorChannels, missingNodesChannel, stats, epoch)
		sm.snapshotUserAccountDataTrie(rootHash, iteratorChannels, missingNodesChannel, stats, epoch, trieStorageManager)

		stats.SnapshotFinished()
	}()

	go sm.syncMissingNodes(missingNodesChannel, iteratorChannels.ErrChan, stats, sm.getTrieSyncer())

	go sm.processSnapshotCompletion(stats, trieStorageManager, missingNodesChannel, iteratorChannels.ErrChan, rootHash, epoch)
}

func (sm *snapshotsManager) earlySnapshotCompletion(stats *snapshotStatistics, trieStorageManager common.StorageManager) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	stats.SnapshotFinished()
	sm.isSnapshotInProgress.Reset()
	trieStorageManager.ExitPruningBufferingMode()
}

func (sm *snapshotsManager) snapshotUserAccountDataTrie(
	mainTrieRootHash []byte,
	iteratorChannels *common.TrieIteratorChannels,
	missingNodesChannel chan []byte,
	stats common.SnapshotStatisticsHandler,
	epoch uint32,
	trieStorageManager common.StorageManager,
) {
	if iteratorChannels.LeavesChan == nil {
		return
	}

	for leaf := range iteratorChannels.LeavesChan {
		userAccount, skipAccount, err := getUserAccountFromBytes(sm.accountFactory, sm.marshaller, leaf.Key(), leaf.Value())
		if err != nil {
			iteratorChannels.ErrChan.WriteInChanNonBlocking(err)
			return
		}
		if skipAccount {
			continue
		}

		if len(userAccount.GetRootHash()) == 0 {
			continue
		}

		stats.NewSnapshotStarted()

		iteratorChannelsForDataTries := &common.TrieIteratorChannels{
			LeavesChan: nil,
			ErrChan:    iteratorChannels.ErrChan,
		}

		address := sm.addressConverter.SilentEncode(userAccount.AddressBytes(), log)
		trieStorageManager.TakeSnapshot(address, userAccount.GetRootHash(), mainTrieRootHash, iteratorChannelsForDataTries, missingNodesChannel, stats, epoch)
	}
}

func (sm *snapshotsManager) syncMissingNodes(missingNodesChan chan []byte, errChan common.BufferedErrChan, stats *snapshotStatistics, syncer AccountsDBSyncer) {
	defer stats.SyncFinished()

	if check.IfNil(syncer) {
		log.Error("can not sync missing nodes", "error", ErrNilTrieSyncer.Error())
		for missingNode := range missingNodesChan {
			log.Warn("could not sync node", "hash", missingNode)
		}
		errChan.WriteInChanNonBlocking(ErrNilTrieSyncer)
		return
	}

	for missingNode := range missingNodesChan {
		err := syncer.SyncAccounts(missingNode, storageMarker.NewDisabledStorageMarker())
		if err != nil {
			log.Error("could not sync missing node",
				"missing node hash", missingNode,
				"error", err,
			)
			errChan.WriteInChanNonBlocking(err)
		}
	}
}

func (sm *snapshotsManager) processSnapshotCompletion(
	stats *snapshotStatistics,
	trieStorageManager common.StorageManager,
	missingNodesCh chan []byte,
	errChan common.BufferedErrChan,
	rootHash []byte,
	epoch uint32,
) {
	sm.finishSnapshotOperation(rootHash, stats, missingNodesCh, sm.stateMetrics.GetSnapshotMessage(), trieStorageManager)

	defer func() {
		sm.isSnapshotInProgress.Reset()
		sm.stateMetrics.UpdateMetricsOnSnapshotCompletion(stats)
		sm.printStorageStatistics()
		errChan.Close()
	}()

	errorDuringSnapshot := errChan.ReadFromChanNonBlocking()
	shouldNotMarkActive := trieStorageManager.IsClosed() || errorDuringSnapshot != nil
	if shouldNotMarkActive {
		log.Debug("will not set activeDB in epoch as the snapshot might be incomplete",
			"epoch", epoch, "trie storage manager closed", trieStorageManager.IsClosed(),
			"errors during snapshot found", errorDuringSnapshot)
		return
	}

	sm.lastSnapshotMarker.RemoveMarker(trieStorageManager, epoch, rootHash)

	log.Debug("set activeDB in epoch", "epoch", epoch)
	errPut := trieStorageManager.PutInEpochWithoutCache([]byte(common.ActiveDBKey), []byte(common.ActiveDBVal), epoch)
	handleLoggingWhenError("error while putting active DB value into main storer", errPut)
}

func (sm *snapshotsManager) printStorageStatistics() {
	stats := sm.stateStatsHandler.SnapshotStats()
	if stats != nil {
		log.Debug("snapshot storage statistics",
			"stats", stats,
		)
	}
}

func (sm *snapshotsManager) finishSnapshotOperation(
	rootHash []byte,
	stats *snapshotStatistics,
	missingNodesCh chan []byte,
	message string,
	trieStorageManager common.StorageManager,
) {
	stats.WaitForSnapshotsToFinish()
	close(missingNodesCh)
	stats.WaitForSyncToFinish()

	trieStorageManager.ExitPruningBufferingMode()

	stats.PrintStats(message, rootHash)
}

func (sm *snapshotsManager) waitForCompletionIfAppropriate(stats common.SnapshotStatisticsHandler) {
	shouldSerializeSnapshots := sm.shouldSerializeSnapshots || sm.processingMode == common.ImportDb
	if !shouldSerializeSnapshots {
		return
	}

	log.Debug("manually setting idle on the process status handler in order to be able to start & complete the snapshotting/checkpointing process")
	sm.processStatusHandler.SetIdle()

	stats.WaitForSnapshotsToFinish()
}

// IsInterfaceNil returns true if there is no value under the interface
func (sm *snapshotsManager) IsInterfaceNil() bool {
	return sm == nil
}

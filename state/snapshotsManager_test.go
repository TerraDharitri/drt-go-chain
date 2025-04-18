package state_test

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/common/statistics/disabled"
	"github.com/TerraDharitri/drt-go-chain/process/mock"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain/state/iteratorChannelsProvider"
	"github.com/TerraDharitri/drt-go-chain/state/lastSnapshotMarker"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/marshallerMock"
	stateTest "github.com/TerraDharitri/drt-go-chain/testscommon/state"
	"github.com/TerraDharitri/drt-go-chain/testscommon/storageManager"
	"github.com/TerraDharitri/drt-go-chain-core/core/atomic"
	"github.com/stretchr/testify/assert"
)

func getDefaultSnapshotManagerArgs() state.ArgsNewSnapshotsManager {
	return state.ArgsNewSnapshotsManager{
		ShouldSerializeSnapshots: false,
		ProcessingMode:           common.Normal,
		Marshaller:               &marshallerMock.MarshalizerMock{},
		AddressConverter:         &testscommon.PubkeyConverterMock{},
		ProcessStatusHandler:     &testscommon.ProcessStatusHandlerStub{},
		StateMetrics:             &stateTest.StateMetricsStub{},
		AccountFactory:           &stateTest.AccountsFactoryStub{},
		ChannelsProvider:         iteratorChannelsProvider.NewUserStateIteratorChannelsProvider(),
		StateStatsHandler:        disabled.NewStateStatistics(),
		LastSnapshotMarker:       lastSnapshotMarker.NewLastSnapshotMarker(),
	}
}

func TestNewSnapshotsManager(t *testing.T) {
	t.Parallel()

	t.Run("nil marshaller", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.Marshaller = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilMarshalizer, err)
	})
	t.Run("nil process status handler", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.ProcessStatusHandler = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilProcessStatusHandler, err)
	})
	t.Run("nil address converter", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.AddressConverter = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilAddressConverter, err)
	})
	t.Run("nil state metrics", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.StateMetrics = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilStateMetrics, err)
	})
	t.Run("nil channels provider", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.ChannelsProvider = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilChannelsProvider, err)
	})
	t.Run("nil account factory", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.AccountFactory = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilAccountFactory, err)
	})
	t.Run("nil stats handler", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.StateStatsHandler = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilStatsHandler, err)
	})
	t.Run("nil last snapshot marker", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.LastSnapshotMarker = nil

		sm, err := state.NewSnapshotsManager(args)
		assert.Nil(t, sm)
		assert.Equal(t, state.ErrNilLastSnapshotMarker, err)
	})
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		sm, err := state.NewSnapshotsManager(args)
		assert.NotNil(t, sm)
		assert.NoError(t, err)
	})
}

func TestSnapshotsManager_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	instance := state.NewNilSnapshotsManager()
	assert.True(t, instance.IsInterfaceNil())

	instance, _ = state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
	assert.False(t, instance.IsInterfaceNil())
}

func TestSnapshotsManager_SetSyncer(t *testing.T) {
	t.Parallel()

	t.Run("nil syncer", func(t *testing.T) {
		t.Parallel()

		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		assert.Equal(t, state.ErrNilTrieSyncer, sm.SetSyncer(nil))
	})
	t.Run("set syncer concurrently", func(t *testing.T) {
		t.Parallel()

		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())

		numConcurrentCalls := 10
		wg := sync.WaitGroup{}
		wg.Add(numConcurrentCalls)
		for i := 0; i < numConcurrentCalls; i++ {
			go func() {
				time.Sleep(10 * time.Millisecond)
				err := sm.SetSyncer(&mock.AccountsDBSyncerStub{})
				assert.Nil(t, err)
				wg.Done()
			}()
		}
		wg.Wait()
	})
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		ts := &mock.AccountsDBSyncerStub{}
		err := sm.SetSyncer(ts)
		assert.Nil(t, err)
	})
}

func TestSnapshotsManager_StartSnapshotAfterRestartIfNeeded(t *testing.T) {
	t.Parallel()

	t.Run("nil syncer", func(t *testing.T) {
		t.Parallel()

		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		assert.Equal(t, state.ErrNilTrieSyncer, sm.StartSnapshotAfterRestartIfNeeded(&storageManager.StorageManagerStub{}))
	})
	t.Run("tsm should not take snapshot", func(t *testing.T) {
		t.Parallel()

		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		ts := &mock.AccountsDBSyncerStub{}
		_ = sm.SetSyncer(ts)

		tsm := &storageManager.StorageManagerStub{
			ShouldTakeSnapshotCalled: func() bool {
				return false
			},
			GetFromCurrentEpochCalled: func(bytes []byte) ([]byte, error) {
				assert.Fail(t, "should not have been called")
				return nil, nil
			},
		}

		err := sm.StartSnapshotAfterRestartIfNeeded(tsm)
		assert.Nil(t, err)
	})
	t.Run("tsm get from current epoch error", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("expected error")
		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		ts := &mock.AccountsDBSyncerStub{}
		_ = sm.SetSyncer(ts)

		tsm := &storageManager.StorageManagerStub{
			ShouldTakeSnapshotCalled: func() bool {
				return true
			},
			GetFromCurrentEpochCalled: func(bytes []byte) ([]byte, error) {
				return nil, expectedErr
			},
			GetLatestStorageEpochCalled: func() (uint32, error) {
				assert.Fail(t, "should not have been called")
				return 0, nil
			},
		}

		err := sm.StartSnapshotAfterRestartIfNeeded(tsm)
		assert.Nil(t, err)
	})
	t.Run("tsm get latest storage epoch error", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("expected error")
		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		ts := &mock.AccountsDBSyncerStub{}
		_ = sm.SetSyncer(ts)

		tsm := &storageManager.StorageManagerStub{
			ShouldTakeSnapshotCalled: func() bool {
				return true
			},
			GetFromCurrentEpochCalled: func(bytes []byte) ([]byte, error) {
				return nil, nil
			},
			GetLatestStorageEpochCalled: func() (uint32, error) {
				return 0, expectedErr
			},
		}

		err := sm.StartSnapshotAfterRestartIfNeeded(tsm)
		assert.Nil(t, err)
	})
}

func TestSnapshotsManager_SnapshotState(t *testing.T) {
	t.Parallel()

	rootHash := []byte("rootHash")
	epoch := uint32(5)

	t.Run("nil snapshots manager should not panic", func(t *testing.T) {
		t.Parallel()

		defer func() {
			r := recover()
			if r != nil {
				assert.Fail(t, fmt.Sprintf("should have not panicked %v", r))
			}
		}()

		args := getDefaultSnapshotManagerArgs()
		sm, _ := state.NewSnapshotsManager(args)
		sm.SnapshotState(rootHash, epoch, nil)
	})
	t.Run("storage manager does not support snapshots should not initiate snapshotting", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.StateMetrics = &stateTest.StateMetricsStub{
			GetSnapshotMessageCalled: func() string {
				assert.Fail(t, "should have not called GetSnapshotMessage")
				return ""
			},
		}
		sm, _ := state.NewSnapshotsManager(args)
		tsm := &storageManager.StorageManagerStub{
			PutInEpochCalled: func(key []byte, val []byte, e uint32) error {
				assert.Fail(t, "should have not called put in epoch")
				return nil
			},
			EnterPruningBufferingModeCalled: func() {
				assert.Fail(t, "should have not called enter pruning buffering mode")
			},
			IsSnapshotSupportedCalled: func() bool {
				return false
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)

		lastRootHash, lastEpoch := sm.GetLastSnapshotInfo()
		assert.Nil(t, lastRootHash)
		assert.Zero(t, lastEpoch)
	})
	t.Run("should not start snapshot for same rootHash in same epoch, and lastSnapshot should not be rewritten", func(t *testing.T) {
		t.Parallel()

		args := getDefaultSnapshotManagerArgs()
		args.StateMetrics = &stateTest.StateMetricsStub{
			GetSnapshotMessageCalled: func() string {
				assert.Fail(t, "snapshot must not start in this case, this func should not have been called")
				return ""
			},
		}
		sm, _ := state.NewSnapshotsManager(args)
		sm.SetLastSnapshotInfo(rootHash, epoch)
		tsm := &storageManager.StorageManagerStub{
			PutInEpochCalled: func(_ []byte, _ []byte, _ uint32) error {
				assert.Fail(t, "lastSnapshot should not be rewritten in this case")
				return nil
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)
	})
	t.Run("should not start snapshot if another snapshot is in progress, lastSnapshot should be saved", func(t *testing.T) {
		t.Parallel()

		putInEpochCalled := atomic.Flag{}

		args := getDefaultSnapshotManagerArgs()
		args.StateMetrics = &stateTest.StateMetricsStub{
			GetSnapshotMessageCalled: func() string {
				assert.Fail(t, "snapshot must not start in this case, this func should not have been called")
				return ""
			},
		}
		sm, _ := state.NewSnapshotsManager(args)
		sm.SetSnapshotInProgress()
		tsm := &storageManager.StorageManagerStub{
			PutInEpochCalled: func(key []byte, val []byte, e uint32) error {
				assert.Equal(t, []byte("lastSnapshot"), key)
				assert.Equal(t, rootHash, val)
				assert.Equal(t, epoch, e)
				putInEpochCalled.SetValue(true)
				return nil
			},
			EnterPruningBufferingModeCalled: func() {
				assert.Fail(t, "the func should have returned before this is called")
			},
			GetLatestStorageEpochCalled: func() (uint32, error) {
				return epoch, nil
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)
		for !putInEpochCalled.IsSet() {
			time.Sleep(10 * time.Millisecond)
		}
	})
	t.Run("starting snapshot sets some parameters", func(t *testing.T) {
		t.Parallel()

		putInEpochCalled := atomic.Flag{}
		enterPruningBufferingModeCalled := false
		getSnapshotMessageCalled := atomic.Flag{}

		args := getDefaultSnapshotManagerArgs()
		args.StateMetrics = &stateTest.StateMetricsStub{
			GetSnapshotMessageCalled: func() string {
				getSnapshotMessageCalled.SetValue(true)
				return ""
			},
		}
		sm, _ := state.NewSnapshotsManager(args)
		tsm := &storageManager.StorageManagerStub{
			PutInEpochCalled: func(key []byte, val []byte, e uint32) error {
				assert.Equal(t, []byte("lastSnapshot"), key)
				assert.Equal(t, rootHash, val)
				assert.Equal(t, epoch, e)
				putInEpochCalled.SetValue(true)
				return nil
			},
			EnterPruningBufferingModeCalled: func() {
				enterPruningBufferingModeCalled = true
				for !putInEpochCalled.IsSet() {
					time.Sleep(10 * time.Millisecond)
				}
			},
			GetLatestStorageEpochCalled: func() (uint32, error) {
				return epoch, nil
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)
		assert.True(t, getSnapshotMessageCalled.IsSet())
		assert.True(t, putInEpochCalled.IsSet())
		assert.True(t, enterPruningBufferingModeCalled)
		assert.True(t, sm.IsSnapshotInProgress())

		lastRootHash, lastEpoch := sm.GetLastSnapshotInfo()
		assert.Equal(t, rootHash, lastRootHash)
		assert.Equal(t, epoch, lastEpoch)
	})
	t.Run("waiting for the correct storage epoch fails", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")
		getLatestStorageEpochCalled := atomic.Flag{}

		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		enterPruningBufferingModeCalled := atomic.Flag{}
		exitPruningBufferingModeCalled := atomic.Flag{}
		tsm := &storageManager.StorageManagerStub{
			GetLatestStorageEpochCalled: func() (uint32, error) {
				for !sm.IsSnapshotInProgress() {
					time.Sleep(10 * time.Millisecond)
				}
				getLatestStorageEpochCalled.SetValue(true)
				return 0, expectedErr
			},
			ShouldTakeSnapshotCalled: func() bool {
				assert.Fail(t, "the func should have returned before this is called")
				return false
			},
			EnterPruningBufferingModeCalled: func() {
				enterPruningBufferingModeCalled.SetValue(true)
			},
			ExitPruningBufferingModeCalled: func() {
				exitPruningBufferingModeCalled.SetValue(true)
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)
		for sm.IsSnapshotInProgress() {
			time.Sleep(10 * time.Millisecond)
		}

		assert.True(t, getLatestStorageEpochCalled.IsSet())
		assert.True(t, enterPruningBufferingModeCalled.IsSet())
		assert.True(t, exitPruningBufferingModeCalled.IsSet())
	})
	t.Run("tsm signals that a snapshot should not be taken", func(t *testing.T) {
		t.Parallel()

		shouldTakeSnapshotCalled := false

		args := getDefaultSnapshotManagerArgs()
		args.StateMetrics = &stateTest.StateMetricsStub{
			UpdateMetricsOnSnapshotStartCalled: func() {
				assert.Fail(t, "the func should have returned before this is called")
			},
		}
		sm, _ := state.NewSnapshotsManager(args)
		enterPruningBufferingModeCalled := atomic.Flag{}
		exitPruningBufferingModeCalled := atomic.Flag{}
		tsm := &storageManager.StorageManagerStub{
			GetLatestStorageEpochCalled: func() (uint32, error) {
				return 5, nil
			},
			ShouldTakeSnapshotCalled: func() bool {
				shouldTakeSnapshotCalled = true
				assert.True(t, sm.IsSnapshotInProgress())
				return false
			},
			EnterPruningBufferingModeCalled: func() {
				enterPruningBufferingModeCalled.SetValue(true)
			},
			ExitPruningBufferingModeCalled: func() {
				exitPruningBufferingModeCalled.SetValue(true)
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)
		for sm.IsSnapshotInProgress() {
			time.Sleep(10 * time.Millisecond)
		}

		assert.True(t, shouldTakeSnapshotCalled)
		assert.True(t, enterPruningBufferingModeCalled.IsSet())
		assert.True(t, exitPruningBufferingModeCalled.IsSet())
	})
	t.Run("snapshot with errors does not mark active and does not remove lastSnapshot", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		sm, _ := state.NewSnapshotsManager(getDefaultSnapshotManagerArgs())
		tsm := &storageManager.StorageManagerStub{
			GetLatestStorageEpochCalled: func() (uint32, error) {
				return 5, nil
			},
			ShouldTakeSnapshotCalled: func() bool {
				return true
			},
			TakeSnapshotCalled: func(_ string, _ []byte, _ []byte, channels *common.TrieIteratorChannels, _ chan []byte, stats common.SnapshotStatisticsHandler, _ uint32) {
				stats.SnapshotFinished()
				close(channels.LeavesChan)
				channels.ErrChan.WriteInChanNonBlocking(expectedErr)
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)
		for sm.IsSnapshotInProgress() {
			time.Sleep(10 * time.Millisecond)
		}
	})
	t.Run("snapshot ok should remove lastSnapshot from all active storers and mark db as complete", func(t *testing.T) {
		t.Parallel()

		putInEpochWithoutCacheCalled := false
		removeFromAllActiveEpochsCalled := false

		args := getDefaultSnapshotManagerArgs()
		args.ChannelsProvider = iteratorChannelsProvider.NewUserStateIteratorChannelsProvider()
		sm, _ := state.NewSnapshotsManager(args)
		_ = sm.SetSyncer(&mock.AccountsDBSyncerStub{})
		tsm := &storageManager.StorageManagerStub{
			GetLatestStorageEpochCalled: func() (uint32, error) {
				return 5, nil
			},
			ShouldTakeSnapshotCalled: func() bool {
				return true
			},
			TakeSnapshotCalled: func(_ string, _ []byte, _ []byte, channels *common.TrieIteratorChannels, _ chan []byte, stats common.SnapshotStatisticsHandler, u uint32) {
				stats.SnapshotFinished()
				close(channels.LeavesChan)
			},
			RemoveFromAllActiveEpochsCalled: func(hash []byte) error {
				assert.True(t, sm.IsSnapshotInProgress())
				assert.Equal(t, []byte("lastSnapshot"), hash)
				removeFromAllActiveEpochsCalled = true
				return nil
			},
			PutInEpochWithoutCacheCalled: func(key []byte, val []byte, e uint32) error {
				assert.Equal(t, []byte(common.ActiveDBKey), key)
				assert.Equal(t, []byte(common.ActiveDBVal), val)
				assert.Equal(t, epoch, e)
				putInEpochWithoutCacheCalled = true
				return nil
			},
		}

		sm.SnapshotState(rootHash, epoch, tsm)
		for sm.IsSnapshotInProgress() {
			time.Sleep(10 * time.Millisecond)
		}

		assert.True(t, putInEpochWithoutCacheCalled)
		assert.True(t, removeFromAllActiveEpochsCalled)
	})
}

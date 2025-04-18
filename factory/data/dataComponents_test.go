package data_test

import (
	"errors"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain/config"
	errorsDrt "github.com/TerraDharitri/drt-go-chain/errors"
	dataComp "github.com/TerraDharitri/drt-go-chain/factory/data"
	"github.com/TerraDharitri/drt-go-chain/factory/mock"
	componentsMock "github.com/TerraDharitri/drt-go-chain/testscommon/components"
	"github.com/TerraDharitri/drt-go-chain/testscommon/factory"
	"github.com/stretchr/testify/require"
)

func TestNewDataComponentsFactory(t *testing.T) {
	t.Parallel()

	t.Run("nil shard coordinator should error", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.ShardCoordinator = nil

		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.Nil(t, dcf)
		require.Equal(t, errorsDrt.ErrNilShardCoordinator, err)
	})
	t.Run("nil core components should error", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		args := componentsMock.GetDataArgs(nil, shardCoordinator)
		args.Core = nil

		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.Nil(t, dcf)
		require.Equal(t, errorsDrt.ErrNilCoreComponents, err)
	})
	t.Run("nil status core components should error", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.StatusCore = nil

		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.Nil(t, dcf)
		require.Equal(t, errorsDrt.ErrNilStatusCoreComponents, err)
	})
	t.Run("nil crypto components should error", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.Crypto = nil

		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.Nil(t, dcf)
		require.Equal(t, errorsDrt.ErrNilCryptoComponents, err)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)
		require.NotNil(t, dcf)
	})
}

func TestDataComponentsFactory_Create(t *testing.T) {
	t.Parallel()

	t.Run("NewBlockChain returns error for shard", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.StatusCore = &factory.StatusCoreComponentsStub{
			AppStatusHandlerField: nil,
		}
		args.Config.ShardHdrNonceHashStorage = config.StorageConfig{}
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)

		dc, err := dcf.Create()
		require.Error(t, err)
		require.Nil(t, dc)
	})
	t.Run("NewBlockChain returns error for meta", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		shardCoordinator.CurrentShard = core.MetachainShardId
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.StatusCore = &factory.StatusCoreComponentsStub{
			AppStatusHandlerField: nil,
		}
		args.Config.ShardHdrNonceHashStorage = config.StorageConfig{}
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)

		dc, err := dcf.Create()
		require.Error(t, err)
		require.Nil(t, dc)
	})
	t.Run("createBlockChainFromConfig returns error", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		shardCoordinator.CurrentShard = 12345
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.Config.ShardHdrNonceHashStorage = config.StorageConfig{}
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)

		dc, err := dcf.Create()
		require.Equal(t, errorsDrt.ErrBlockchainCreation, err)
		require.Nil(t, dc)
	})
	t.Run("NewStorageServiceFactory returns error", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.Config.StoragePruning.NumActivePersisters = 0
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)

		dc, err := dcf.Create()
		require.Error(t, err)
		require.Nil(t, dc)
	})
	t.Run("createDataStoreFromConfig fails for shard due to bad config", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.Config.ShardHdrNonceHashStorage = config.StorageConfig{}
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)

		dc, err := dcf.Create()
		require.Error(t, err)
		require.Nil(t, dc)
	})
	t.Run("createDataStoreFromConfig fails, invalid shard", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		cnt := 0
		shardCoordinator.SelfIDCalled = func() uint32 {
			cnt++
			if cnt > 1 {
				return 12345
			}
			return 0
		}
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.Config.ShardHdrNonceHashStorage = config.StorageConfig{}
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)

		dc, err := dcf.Create()
		require.Equal(t, errorsDrt.ErrDataStoreCreation, err)
		require.Nil(t, dc)
	})
	t.Run("NewDataPoolFromConfig fails should error", func(t *testing.T) {
		t.Parallel()

		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		coreComponents := componentsMock.GetCoreComponents()
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		args.Config.TxBlockBodyDataPool.Type = "invalid"
		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)

		dc, err := dcf.Create()
		require.True(t, errors.Is(err, errorsDrt.ErrDataPoolCreation))
		require.Nil(t, dc)
	})
	t.Run("should work for shard", func(t *testing.T) {
		t.Parallel()

		coreComponents := componentsMock.GetCoreComponents()
		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
		dcf, err := dataComp.NewDataComponentsFactory(args)

		require.NoError(t, err)
		dc, err := dcf.Create()
		require.NoError(t, err)
		require.NotNil(t, dc)
	})
	t.Run("should work for meta", func(t *testing.T) {
		t.Parallel()

		coreComponents := componentsMock.GetCoreComponents()
		shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
		shardCoordinator.CurrentShard = core.MetachainShardId
		args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)

		dcf, err := dataComp.NewDataComponentsFactory(args)
		require.NoError(t, err)
		dc, err := dcf.Create()
		require.NoError(t, err)
		require.NotNil(t, dc)
	})
}

func TestManagedDataComponents_CloseShouldWork(t *testing.T) {
	t.Parallel()

	coreComponents := componentsMock.GetCoreComponents()
	shardCoordinator := mock.NewMultiShardsCoordinatorMock(2)
	args := componentsMock.GetDataArgs(coreComponents, shardCoordinator)
	dcf, _ := dataComp.NewDataComponentsFactory(args)

	dc, _ := dcf.Create()

	err := dc.Close()
	require.NoError(t, err)
}

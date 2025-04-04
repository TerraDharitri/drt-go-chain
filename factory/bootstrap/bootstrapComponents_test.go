package bootstrap_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/config"
	errorsDrt "github.com/TerraDharitri/drt-go-chain/errors"
	"github.com/TerraDharitri/drt-go-chain/factory/bootstrap"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	componentsMock "github.com/TerraDharitri/drt-go-chain/testscommon/components"
	"github.com/TerraDharitri/drt-go-chain/testscommon/factory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBootstrapComponentsFactory(t *testing.T) {
	t.Parallel()

	args := componentsMock.GetBootStrapFactoryArgs()
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		bcf, err := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)
		require.Nil(t, err)
	})
	t.Run("nil core components should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.CoreComponents = nil
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrNilCoreComponentsHolder, err)
	})
	t.Run("nil enable epochs handler should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.CoreComponents = &factory.CoreComponentsHolderStub{
			EnableEpochsHandlerCalled: func() common.EnableEpochsHandler {
				return nil
			},
		}
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrNilEnableEpochsHandler, err)
	})
	t.Run("nil crypto components should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.CryptoComponents = nil
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrNilCryptoComponentsHolder, err)
	})
	t.Run("nil network components should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.NetworkComponents = nil
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrNilNetworkComponentsHolder, err)
	})
	t.Run("nil status core components should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.StatusCoreComponents = nil
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrNilStatusCoreComponents, err)
	})
	t.Run("nil trie sync statistics should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.StatusCoreComponents = &factory.StatusCoreComponentsStub{
			TrieSyncStatisticsField: nil,
		}
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrNilTrieSyncStatistics, err)
	})
	t.Run("nil app status handler should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.StatusCoreComponents = &factory.StatusCoreComponentsStub{
			AppStatusHandlerField:   nil,
			TrieSyncStatisticsField: &testscommon.SizeSyncStatisticsHandlerStub{},
		}
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrNilAppStatusHandler, err)
	})
	t.Run("empty working dir should error", func(t *testing.T) {
		t.Parallel()

		argsCopy := args
		argsCopy.WorkingDir = ""
		bcf, err := bootstrap.NewBootstrapComponentsFactory(argsCopy)
		require.Nil(t, bcf)
		require.Equal(t, errorsDrt.ErrInvalidWorkingDir, err)
	})
}

func TestBootstrapComponentsFactory_Create(t *testing.T) {
	t.Parallel()

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, err)
		require.NotNil(t, bc)
	})
	t.Run("ProcessDestinationShardAsObserver fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		args.PrefConfig.Preferences.DestinationShardAsObserver = ""
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.True(t, strings.Contains(err.Error(), "DestinationShardAsObserver"))
	})
	t.Run("NewCache fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		args.Config.Versions.Cache = config.CacheConfig{
			Type:        "LRU",
			SizeInBytes: 1,
		}
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.True(t, strings.Contains(err.Error(), "LRU"))
	})
	t.Run("NewHeaderVersionHandler fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		args.Config.Versions.DefaultVersion = string(bytes.Repeat([]byte("a"), 20))
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.NotNil(t, err)
	})
	t.Run("NewHeaderIntegrityVerifier fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		coreComponents := componentsMock.GetDefaultCoreComponents()
		coreComponents.ChainIdCalled = func() string {
			return ""
		}
		args.CoreComponents = coreComponents
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.NotNil(t, err)
	})
	t.Run("CreateShardCoordinator fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		coreComponents := componentsMock.GetDefaultCoreComponents()
		coreComponents.NodesConfig = nil
		args.CoreComponents = coreComponents
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.NotNil(t, err)
	})
	t.Run("NewBootstrapDataProvider fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		coreComponents := componentsMock.GetDefaultCoreComponents()
		args.CoreComponents = coreComponents
		coreComponents.IntMarsh = nil
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.True(t, errors.Is(err, errorsDrt.ErrNewBootstrapDataProvider))
	})
	t.Run("import db mode - NewStorageEpochStartBootstrap fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		coreComponents := componentsMock.GetDefaultCoreComponents()
		args.CoreComponents = coreComponents
		coreComponents.RatingHandler = nil
		args.ImportDbConfig.IsImportDBMode = true
		bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.True(t, errors.Is(err, errorsDrt.ErrNewStorageEpochStartBootstrap))
	})
	t.Run("NewStorageEpochStartBootstrap fails should error", func(t *testing.T) {
		t.Parallel()

		args := componentsMock.GetBootStrapFactoryArgs()
		coreComponents := componentsMock.GetDefaultCoreComponents()
		args.CoreComponents = coreComponents
		coreComponents.RatingHandler = nil
		bcf, err := bootstrap.NewBootstrapComponentsFactory(args)
		require.Nil(t, err)
		require.NotNil(t, bcf)

		bc, err := bcf.Create()
		require.Nil(t, bc)
		require.True(t, errors.Is(err, errorsDrt.ErrNewEpochStartBootstrap))
	})
}
func TestBootstrapComponents(t *testing.T) {
	t.Parallel()

	args := componentsMock.GetBootStrapFactoryArgs()
	bcf, _ := bootstrap.NewBootstrapComponentsFactory(args)
	require.NotNil(t, bcf)

	bc, err := bcf.Create()
	require.Nil(t, err)
	require.NotNil(t, bc)

	assert.Equal(t, core.NodeTypeObserver, bc.NodeType())
	assert.False(t, check.IfNil(bc.ShardCoordinator()))
	assert.False(t, check.IfNil(bc.HeaderVersionHandler()))
	assert.False(t, check.IfNil(bc.VersionedHeaderFactory()))
	assert.False(t, check.IfNil(bc.HeaderIntegrityVerifier()))

	require.Nil(t, bc.Close())
}

package genesis

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/dataRetriever"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/enableEpochsHandlerMock"
	"github.com/TerraDharitri/drt-go-chain/testscommon/hashingMocks"
	"github.com/TerraDharitri/drt-go-chain/testscommon/storageManager"
	"github.com/TerraDharitri/drt-go-chain/update"
	"github.com/TerraDharitri/drt-go-chain/update/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//TODO increase code coverage

func TestNewStateImport(t *testing.T) {
	trieStorageManagers := make(map[string]common.StorageManager)
	trieStorageManagers[dataRetriever.UserAccountsUnit.String()] = &storageManager.StorageManagerStub{}
	tests := []struct {
		name    string
		args    ArgsNewStateImport
		exError error
	}{
		{
			name: "NilHarforkStorer",
			args: ArgsNewStateImport{
				HardforkStorer:      nil,
				Marshalizer:         &mock.MarshalizerMock{},
				Hasher:              &mock.HasherStub{},
				TrieStorageManagers: trieStorageManagers,
				AddressConverter:    &testscommon.PubkeyConverterMock{},
			},
			exError: update.ErrNilHardforkStorer,
		},
		{
			name: "NilMarshalizer",
			args: ArgsNewStateImport{
				HardforkStorer:      &mock.HardforkStorerStub{},
				Marshalizer:         nil,
				Hasher:              &mock.HasherStub{},
				TrieStorageManagers: trieStorageManagers,
				AddressConverter:    &testscommon.PubkeyConverterMock{},
			},
			exError: update.ErrNilMarshalizer,
		},
		{
			name: "NilHasher",
			args: ArgsNewStateImport{
				HardforkStorer:      &mock.HardforkStorerStub{},
				Marshalizer:         &mock.MarshalizerMock{},
				Hasher:              nil,
				TrieStorageManagers: trieStorageManagers,
				AddressConverter:    &testscommon.PubkeyConverterMock{},
			},
			exError: update.ErrNilHasher,
		},
		{
			name: "Ok",
			args: ArgsNewStateImport{
				HardforkStorer:      &mock.HardforkStorerStub{},
				Marshalizer:         &mock.MarshalizerMock{},
				Hasher:              &mock.HasherStub{},
				TrieStorageManagers: trieStorageManagers,
				AddressConverter:    &testscommon.PubkeyConverterMock{},
				EnableEpochsHandler: &enableEpochsHandlerMock.EnableEpochsHandlerStub{},
			},
			exError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewStateImport(tt.args)
			require.Equal(t, tt.exError, err)
		})
	}
}

func TestImportAll(t *testing.T) {
	t.Parallel()

	trieStorageManagers := make(map[string]common.StorageManager)
	trieStorageManagers[dataRetriever.UserAccountsUnit.String()] = &storageManager.StorageManagerStub{}
	trieStorageManagers[dataRetriever.PeerAccountsUnit.String()] = &storageManager.StorageManagerStub{}

	args := ArgsNewStateImport{
		HardforkStorer:      &mock.HardforkStorerStub{},
		Hasher:              &hashingMocks.HasherMock{},
		Marshalizer:         &mock.MarshalizerMock{},
		TrieStorageManagers: trieStorageManagers,
		ShardID:             0,
		StorageConfig:       config.StorageConfig{},
		AddressConverter:    &testscommon.PubkeyConverterMock{},
		EnableEpochsHandler: &enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	}

	importState, _ := NewStateImport(args)
	require.False(t, check.IfNil(importState))

	err := importState.ImportAll()
	require.Nil(t, err)
}

func TestStateImport_ImportUnFinishedMetaBlocksShouldWork(t *testing.T) {
	t.Parallel()

	trieStorageManagers := make(map[string]common.StorageManager)
	trieStorageManagers[dataRetriever.UserAccountsUnit.String()] = &storageManager.StorageManagerStub{}

	hasher := &hashingMocks.HasherMock{}
	marshahlizer := &mock.MarshalizerMock{}
	metaBlock := &block.MetaBlock{
		Round:   1,
		ChainID: []byte("chainId"),
	}
	metaBlockHash, _ := core.CalculateHash(marshahlizer, hasher, metaBlock)

	args := ArgsNewStateImport{
		HardforkStorer: &mock.HardforkStorerStub{
			GetCalled: func(identifier string, key []byte) ([]byte, error) {
				return marshahlizer.Marshal(metaBlock)
			},
		},
		Hasher:              hasher,
		Marshalizer:         marshahlizer,
		TrieStorageManagers: trieStorageManagers,
		ShardID:             0,
		StorageConfig:       config.StorageConfig{},
		AddressConverter:    &testscommon.PubkeyConverterMock{},
		EnableEpochsHandler: &enableEpochsHandlerMock.EnableEpochsHandlerStub{},
	}

	importState, _ := NewStateImport(args)
	require.False(t, check.IfNil(importState))

	key := fmt.Sprintf("meta@chainId@%s", hex.EncodeToString(metaBlockHash))
	err := importState.importUnFinishedMetaBlocks(UnFinishedMetaBlocksIdentifier, [][]byte{
		[]byte(key),
	})

	require.Nil(t, err)
	assert.Equal(t, importState.importedUnFinishedMetaBlocks[string(metaBlockHash)], metaBlock)
}

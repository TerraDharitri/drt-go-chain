package mock

import (
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain/common"
)

// EpochStartTriesSyncHandlerMock -
type EpochStartTriesSyncHandlerMock struct {
	SyncTriesFromCalled func(meta data.MetaHeaderHandler) error
	GetTriesCalled      func() (map[string]common.Trie, error)
}

// SyncTriesFrom -
func (es *EpochStartTriesSyncHandlerMock) SyncTriesFrom(meta data.MetaHeaderHandler) error {
	if es.SyncTriesFromCalled != nil {
		return es.SyncTriesFromCalled(meta)
	}
	return nil
}

// GetTries -
func (es *EpochStartTriesSyncHandlerMock) GetTries() (map[string]common.Trie, error) {
	if es.GetTriesCalled != nil {
		return es.GetTriesCalled()
	}
	return nil, nil
}

// IsInterfaceNil -
func (es *EpochStartTriesSyncHandlerMock) IsInterfaceNil() bool {
	return es == nil
}

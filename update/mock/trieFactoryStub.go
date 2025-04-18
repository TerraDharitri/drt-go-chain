package mock

import (
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/config"
)

// TrieFactoryStub -
type TrieFactoryStub struct {
	CreateCalled func(config config.StorageConfig, s string, b bool) (common.StorageManager, common.Trie, error)
}

// Create -
func (t *TrieFactoryStub) Create(config config.StorageConfig, s string, b bool) (common.StorageManager, common.Trie, error) {
	if t.CreateCalled != nil {
		return t.CreateCalled(config, s, b)
	}
	return nil, nil, nil
}

// IsInterfaceNil -
func (t *TrieFactoryStub) IsInterfaceNil() bool {
	return t == nil
}

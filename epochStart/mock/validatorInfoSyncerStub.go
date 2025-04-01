package mock

import (
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain-core/data"
)

// ValidatorInfoSyncerStub -
type ValidatorInfoSyncerStub struct {
}

// SyncMiniBlocks -
func (vip *ValidatorInfoSyncerStub) SyncMiniBlocks(_ data.HeaderHandler) ([][]byte, data.BodyHandler, error) {
	return nil, nil, nil
}

// SyncValidatorsInfo -
func (vip *ValidatorInfoSyncerStub) SyncValidatorsInfo(_ data.BodyHandler) ([][]byte, map[string]*state.ShardValidatorInfo, error) {
	return nil, nil, nil
}

// IsInterfaceNil -
func (vip *ValidatorInfoSyncerStub) IsInterfaceNil() bool {
	return vip == nil
}

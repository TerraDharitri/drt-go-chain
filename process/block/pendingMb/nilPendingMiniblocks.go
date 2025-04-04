package pendingMb

import (
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain-core/data"
)

var _ process.PendingMiniBlocksHandler = (*nilPendingMiniBlocks)(nil)

type nilPendingMiniBlocks struct {
}

// NewNilPendingMiniBlocks will create a new pendingMiniBlocks object
func NewNilPendingMiniBlocks() (*nilPendingMiniBlocks, error) {
	return &nilPendingMiniBlocks{}, nil
}

// AddProcessedHeader will add in pending list all miniblocks hashes from a given metablock
func (p *nilPendingMiniBlocks) AddProcessedHeader(_ data.HeaderHandler) error {
	return nil
}

// RevertHeader will remove from pending list all miniblocks hashes from a given metablock
func (p *nilPendingMiniBlocks) RevertHeader(_ data.HeaderHandler) error {
	return nil
}

// GetPendingMiniBlocks will return the pending miniblocks hashes for a given shard
func (p *nilPendingMiniBlocks) GetPendingMiniBlocks(_ uint32) [][]byte {
	return make([][]byte, 0)
}

// SetPendingMiniBlocks will set the pending miniblocks hashes for a given shard
func (p *nilPendingMiniBlocks) SetPendingMiniBlocks(_ uint32, _ [][]byte) {
}

// IsInterfaceNil returns true if there is no value under the interface
func (p *nilPendingMiniBlocks) IsInterfaceNil() bool {
	return p == nil
}

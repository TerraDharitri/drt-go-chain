package trieNodeData

import (
	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain/common"
)

type intermediaryNodeData struct {
	*baseNodeData
}

// NewIntermediaryNodeData creates a new intermediary node data
func NewIntermediaryNodeData(key common.KeyBuilder, data []byte) (*intermediaryNodeData, error) {
	if check.IfNil(key) {
		return nil, ErrNilKeyBuilder
	}

	return &intermediaryNodeData{
		baseNodeData: &baseNodeData{
			keyBuilder: key,
			data:       data,
		},
	}, nil
}

// IsLeaf returns false
func (ind *intermediaryNodeData) IsLeaf() bool {
	return false
}

// GetVersion returns NotSpecified
func (ind *intermediaryNodeData) GetVersion() core.TrieNodeVersion {
	return core.NotSpecified
}

// IsInterfaceNil returns true if there is no value under the interface
func (ind *intermediaryNodeData) IsInterfaceNil() bool {
	return ind == nil
}

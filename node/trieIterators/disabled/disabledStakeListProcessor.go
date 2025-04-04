package disabled

import (
	"context"
	"errors"

	"github.com/TerraDharitri/drt-go-chain-core/data/api"
)

var errCannotReturnDirectStakedListFromShardNode = errors.New("direct staked list can not be returned by a shard node")

type directStakedListProcessor struct{}

// NewDisabledDirectStakedListProcessor returns a disabled implementation to be used on shard nodes
func NewDisabledDirectStakedListProcessor() *directStakedListProcessor {
	return &directStakedListProcessor{}
}

// GetDirectStakedList returns the errCannotReturnDirectStakedListFromShardNode error
func (dslp *directStakedListProcessor) GetDirectStakedList(_ context.Context) ([]*api.DirectStakedValue, error) {
	return nil, errCannotReturnDirectStakedListFromShardNode
}

// IsInterfaceNil returns true if there is no value under the interface
func (dslp *directStakedListProcessor) IsInterfaceNil() bool {
	return dslp == nil
}

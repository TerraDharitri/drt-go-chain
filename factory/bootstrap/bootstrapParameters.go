package bootstrap

import (
	"github.com/TerraDharitri/drt-go-chain/epochStart/bootstrap"
	"github.com/TerraDharitri/drt-go-chain/sharding/nodesCoordinator"
)

type bootstrapParams struct {
	bootstrapParams bootstrap.Parameters
}

// Epoch returns the epoch number after bootstrap
func (bph *bootstrapParams) Epoch() uint32 {
	return bph.bootstrapParams.Epoch
}

// SelfShardID returns the self shard ID after bootstrap
func (bph *bootstrapParams) SelfShardID() uint32 {
	return bph.bootstrapParams.SelfShardId
}

// NumOfShards returns the number of shards after bootstrap
func (bph *bootstrapParams) NumOfShards() uint32 {
	return bph.bootstrapParams.NumOfShards
}

// NodesConfig returns the nodes coordinator config after bootstrap
func (bph *bootstrapParams) NodesConfig() nodesCoordinator.NodesCoordinatorRegistryHandler {
	return bph.bootstrapParams.NodesConfig
}

// IsInterfaceNil returns true if the underlying object is nil
func (bph *bootstrapParams) IsInterfaceNil() bool {
	return bph == nil
}

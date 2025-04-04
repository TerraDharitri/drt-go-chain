package bootstrapMocks

import "github.com/TerraDharitri/drt-go-chain/sharding/nodesCoordinator"

// BootstrapParamsHandlerMock -
type BootstrapParamsHandlerMock struct {
	EpochCalled       func() uint32
	SelfShardIDCalled func() uint32
	NumOfShardsCalled func() uint32
	NodesConfigCalled func() nodesCoordinator.NodesCoordinatorRegistryHandler
}

// Epoch -
func (bphm *BootstrapParamsHandlerMock) Epoch() uint32 {
	if bphm.EpochCalled != nil {
		return bphm.EpochCalled()
	}

	return 0
}

// SelfShardID -
func (bphm *BootstrapParamsHandlerMock) SelfShardID() uint32 {
	if bphm.SelfShardIDCalled != nil {
		return bphm.SelfShardIDCalled()
	}
	return 0
}

// NumOfShards -
func (bphm *BootstrapParamsHandlerMock) NumOfShards() uint32 {
	if bphm.NumOfShardsCalled != nil {
		return bphm.NumOfShardsCalled()
	}
	return 1
}

// NodesConfig -
func (bphm *BootstrapParamsHandlerMock) NodesConfig() nodesCoordinator.NodesCoordinatorRegistryHandler {
	if bphm.NodesConfigCalled != nil {
		return bphm.NodesConfigCalled()
	}
	return nil
}

// IsInterfaceNil -
func (bphm *BootstrapParamsHandlerMock) IsInterfaceNil() bool {
	return bphm == nil
}

package bootstrapMocks

import (
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/epochStart/bootstrap"
)

// EpochStartBootstrapperStub -
type EpochStartBootstrapperStub struct {
	TrieHolder      common.TriesHolder
	StorageManagers map[string]common.StorageManager
	BootstrapCalled func() (bootstrap.Parameters, error)
}

// Bootstrap -
func (esbs *EpochStartBootstrapperStub) Bootstrap() (bootstrap.Parameters, error) {
	if esbs.BootstrapCalled != nil {
		return esbs.BootstrapCalled()
	}

	return bootstrap.Parameters{}, nil
}

// IsInterfaceNil -
func (esbs *EpochStartBootstrapperStub) IsInterfaceNil() bool {
	return esbs == nil
}

// Close -
func (esbs *EpochStartBootstrapperStub) Close() error {
	return nil
}

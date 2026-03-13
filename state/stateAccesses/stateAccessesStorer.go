package stateAccesses

import (
	"fmt"

	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	data "github.com/TerraDharitri/drt-go-chain-core/data/stateChange"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain/storage"
)

type stateAccessesStorer struct {
	storer     storage.Storer
	marshaller marshal.Marshalizer
}

// NewStateAccessesStorer creates a new state accesses storer
func NewStateAccessesStorer(storer storage.Storer, marshaller marshal.Marshalizer) (state.StateAccessesStorer, error) {
	if check.IfNil(storer) {
		return nil, state.ErrNilStateAccessesStorer
	}
	if check.IfNil(marshaller) {
		return nil, state.ErrNilMarshalizer
	}

	return &stateAccessesStorer{
		storer:     storer,
		marshaller: marshaller,
	}, nil
}

// Store saves the state accesses for the given transactions in the given storer.
func (sas *stateAccessesStorer) Store(stateAccessesForTxs map[string]*data.StateAccesses) error {
	for txHash, stateAccess := range stateAccessesForTxs {
		marshalledData, err := sas.marshaller.Marshal(stateAccess)
		if err != nil {
			return fmt.Errorf("failed to marshal state accesses: %w", err)
		}

		err = sas.storer.Put([]byte(txHash), marshalledData)
		if err != nil {
			return fmt.Errorf("failed to store marshalled data: %w", err)
		}
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (sas *stateAccessesStorer) IsInterfaceNil() bool {
	return sas == nil
}

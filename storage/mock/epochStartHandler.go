package mock

import (
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain/epochStart"
)

// EpochStartNotifierStub -
type EpochStartNotifierStub struct {
	RegisterHandlerCalled   func(handler epochStart.ActionHandler)
	UnregisterHandlerCalled func(handler epochStart.ActionHandler)
	NotifyAllCalled         func(hdr data.HeaderHandler)
}

// RegisterHandler -
func (esnm *EpochStartNotifierStub) RegisterHandler(handler epochStart.ActionHandler) {
	if esnm.RegisterHandlerCalled != nil {
		esnm.RegisterHandlerCalled(handler)
	}
}

// UnregisterHandler -
func (esnm *EpochStartNotifierStub) UnregisterHandler(handler epochStart.ActionHandler) {
	if esnm.UnregisterHandlerCalled != nil {
		esnm.UnregisterHandlerCalled(handler)
	}
}

// NotifyAll -
func (esnm *EpochStartNotifierStub) NotifyAll(hdr data.HeaderHandler) {
	if esnm.NotifyAllCalled != nil {
		esnm.NotifyAllCalled(hdr)
	}
}

// IsInterfaceNil -
func (esnm *EpochStartNotifierStub) IsInterfaceNil() bool {
	return esnm == nil
}

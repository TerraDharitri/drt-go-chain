package epochNotifier

import (
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/epochStart"
)

// EpochNotifierStub -
type EpochNotifierStub struct {
	NewEpochCalled              func(epoch uint32, timestamp uint64)
	RegisterHandlerCalled       func(handler epochStart.ActionHandler)
	CheckEpochCalled            func(header data.HeaderHandler)
	CurrentEpochCalled          func() uint32
	RegisterNotifyHandlerCalled func(handler vmcommon.EpochSubscriberHandler)
}

// NewEpoch -
func (ens *EpochNotifierStub) NewEpoch(epoch uint32, timestamp uint64) {
	if ens.NewEpochCalled != nil {
		ens.NewEpochCalled(epoch, timestamp)
	}
}

// RegisterHandler -
func (ens *EpochNotifierStub) RegisterHandler(handler epochStart.ActionHandler) {
	if ens.RegisterHandlerCalled != nil {
		ens.RegisterHandlerCalled(handler)
	}
}

// CheckEpoch -
func (ens *EpochNotifierStub) CheckEpoch(header data.HeaderHandler) {
	if ens.CheckEpochCalled != nil {
		ens.CheckEpochCalled(header)
	}
}

// RegisterNotifyHandler -
func (ens *EpochNotifierStub) RegisterNotifyHandler(handler vmcommon.EpochSubscriberHandler) {
	if ens.RegisterNotifyHandlerCalled != nil {
		ens.RegisterNotifyHandlerCalled(handler)
	} else {
		if !check.IfNil(handler) {
			handler.EpochConfirmed(0, 0)
		}
	}
}

// CurrentEpoch -
func (ens *EpochNotifierStub) CurrentEpoch() uint32 {
	if ens.CurrentEpochCalled != nil {
		return ens.CurrentEpochCalled()
	}

	return 0
}

// IsInterfaceNil -
func (ens *EpochNotifierStub) IsInterfaceNil() bool {
	return ens == nil
}

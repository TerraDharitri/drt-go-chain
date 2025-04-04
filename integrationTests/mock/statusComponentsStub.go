package mock

import (
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/common/statistics"
	"github.com/TerraDharitri/drt-go-chain/outport"
)

// StatusComponentsStub -
type StatusComponentsStub struct {
	Outport                  outport.OutportHandler
	SoftwareVersionCheck     statistics.SoftwareVersionChecker
	ManagedPeersMonitorField common.ManagedPeersMonitor
}

// Create -
func (scs *StatusComponentsStub) Create() error {
	return nil
}

// Close -
func (scs *StatusComponentsStub) Close() error {
	return nil
}

// CheckSubcomponents -
func (scs *StatusComponentsStub) CheckSubcomponents() error {
	return nil
}

// OutportHandler -
func (scs *StatusComponentsStub) OutportHandler() outport.OutportHandler {
	return scs.Outport
}

// SoftwareVersionChecker -
func (scs *StatusComponentsStub) SoftwareVersionChecker() statistics.SoftwareVersionChecker {
	return scs.SoftwareVersionCheck
}

// ManagedPeersMonitor -
func (scs *StatusComponentsStub) ManagedPeersMonitor() common.ManagedPeersMonitor {
	return scs.ManagedPeersMonitorField
}

// IsInterfaceNil -
func (scs *StatusComponentsStub) IsInterfaceNil() bool {
	return scs == nil
}

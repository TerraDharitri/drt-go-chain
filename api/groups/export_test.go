package groups

import "github.com/TerraDharitri/drt-go-chain/process"

// ExecManualTrigger -
const ExecManualTrigger = execManualTrigger

// ExecBroadcastTrigger -
const ExecBroadcastTrigger = execBroadcastTrigger

// CreateSCQuery -
func (vvg *vmValuesGroup) CreateSCQuery(request *VMValueRequest) (*process.SCQuery, error) {
	return vvg.createSCQuery(request)
}

// VmValuesFacadeHandler exported the vm values facade handler interface for testing purposes
type VmValuesFacadeHandler interface {
	vmValuesFacadeHandler
}

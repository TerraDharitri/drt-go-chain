package forking

import vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"

// Handlers -
func (gen *genericEpochNotifier) Handlers() []vmcommon.EpochSubscriberHandler {
	gen.mutHandler.RLock()
	defer gen.mutHandler.RUnlock()

	return gen.handlers
}

// CurrentTimestamp -
func (gen *genericEpochNotifier) CurrentTimestamp() uint64 {
	_, timestamp := gen.getEpochTimestamp()

	return timestamp
}

// Handlers -
func (grn *genericRoundNotifier) Handlers() []vmcommon.RoundSubscriberHandler {
	grn.mutHandler.RLock()
	defer grn.mutHandler.RUnlock()

	return grn.handlers
}

// CurrentTimestamp -
func (grn *genericRoundNotifier) CurrentTimestamp() uint64 {
	_, timestamp := grn.getRoundTimestamp()

	return timestamp
}

package mock

import (
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
)

// StatusComputerStub -
type StatusComputerStub struct {
}

// ComputeStatusWhenInStorageKnowingMiniblock -
func (scs *StatusComputerStub) ComputeStatusWhenInStorageKnowingMiniblock(_ block.Type, _ *transaction.ApiTransactionResult) (transaction.TxStatus, error) {
	return "", nil
}

// ComputeStatusWhenInStorageNotKnowingMiniblock -
func (scs *StatusComputerStub) ComputeStatusWhenInStorageNotKnowingMiniblock(_ uint32, _ *transaction.ApiTransactionResult) (transaction.TxStatus, error) {
	return "", nil
}

// SetStatusIfIsRewardReverted -
func (scs *StatusComputerStub) SetStatusIfIsRewardReverted(_ *transaction.ApiTransactionResult, _ block.Type, _ uint64, _ []byte) (bool, error) {
	return false, nil
}

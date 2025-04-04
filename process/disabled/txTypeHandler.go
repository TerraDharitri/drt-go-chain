package disabled

import (
	"github.com/TerraDharitri/drt-go-chain-core/data"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-chain/process"
)

var log = logger.GetOrCreate("disabledTxTypeHandler")

type txTypeHandler struct {
}

// NewTxTypeHandler returns a new instance of disabled txTypeHandler
func NewTxTypeHandler() *txTypeHandler {
	return &txTypeHandler{}
}

// ComputeTransactionType always returns invalid transaction as it is disabled
func (handler *txTypeHandler) ComputeTransactionType(_ data.TransactionHandler) (process.TransactionType, process.TransactionType, bool) {
	log.Warn("disabled txTypeHandler ComputeTransactionType always returns invalid transaction")
	return process.InvalidTransaction, process.InvalidTransaction, false
}

// IsInterfaceNil returns true if there is no value under the interface
func (handler *txTypeHandler) IsInterfaceNil() bool {
	return handler == nil
}

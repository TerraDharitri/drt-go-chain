package testscommon

import (
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/state"
)

// TxProcessorStub -
type TxProcessorStub struct {
	ProcessTransactionCalled func(transaction *transaction.Transaction) (vmcommon.ReturnCode, error)
	VerifyTransactionCalled  func(tx *transaction.Transaction) error
	VerifyGuardianCalled     func(tx *transaction.Transaction, account state.UserAccountHandler) error
}

// ProcessTransaction -
func (tps *TxProcessorStub) ProcessTransaction(transaction *transaction.Transaction) (vmcommon.ReturnCode, error) {
	if tps.ProcessTransactionCalled != nil {
		return tps.ProcessTransactionCalled(transaction)
	}

	return 0, nil
}

// VerifyTransaction -
func (tps *TxProcessorStub) VerifyTransaction(tx *transaction.Transaction) error {
	if tps.VerifyTransactionCalled != nil {
		return tps.VerifyTransactionCalled(tx)
	}

	return nil
}

// VerifyGuardian -
func (tps *TxProcessorStub) VerifyGuardian(tx *transaction.Transaction, account state.UserAccountHandler) error {
	if tps.VerifyGuardianCalled != nil {
		return tps.VerifyGuardianCalled(tx, account)
	}

	return nil
}

// IsInterfaceNil -
func (tps *TxProcessorStub) IsInterfaceNil() bool {
	return tps == nil
}

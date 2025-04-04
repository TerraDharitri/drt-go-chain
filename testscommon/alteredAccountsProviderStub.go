package testscommon

import (
	"github.com/TerraDharitri/drt-go-chain-core/data/alteredAccount"
	"github.com/TerraDharitri/drt-go-chain-core/data/outport"
	"github.com/TerraDharitri/drt-go-chain/outport/process/alteredaccounts/shared"
)

// AlteredAccountsProviderStub -
type AlteredAccountsProviderStub struct {
	ExtractAlteredAccountsFromPoolCalled func(txPool *outport.TransactionPool, options shared.AlteredAccountsOptions) (map[string]*alteredAccount.AlteredAccount, error)
}

// ExtractAlteredAccountsFromPool -
func (a *AlteredAccountsProviderStub) ExtractAlteredAccountsFromPool(txPool *outport.TransactionPool, options shared.AlteredAccountsOptions) (map[string]*alteredAccount.AlteredAccount, error) {
	if a.ExtractAlteredAccountsFromPoolCalled != nil {
		return a.ExtractAlteredAccountsFromPoolCalled(txPool, options)
	}

	return nil, nil
}

// IsInterfaceNil -
func (a *AlteredAccountsProviderStub) IsInterfaceNil() bool {
	return a == nil
}

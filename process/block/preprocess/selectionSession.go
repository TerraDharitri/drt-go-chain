package preprocess

import (
	"errors"

	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain/storage/txcache"
)

type selectionSession struct {
	accountsAdapter       state.AccountsAdapter
	transactionsProcessor process.TransactionProcessor

	// Cache of accounts, held in the scope of a single selection session.
	// Not concurrency-safe, but never accessed concurrently.
	ephemeralAccountsCache map[string]vmcommon.AccountHandler
}

// ArgsSelectionSession holds the arguments for creating a new selection session.
type ArgsSelectionSession struct {
	AccountsAdapter       state.AccountsAdapter
	TransactionsProcessor process.TransactionProcessor
}

// NewSelectionSession creates a new selection session.
func NewSelectionSession(args ArgsSelectionSession) (*selectionSession, error) {
	if check.IfNil(args.AccountsAdapter) {
		return nil, process.ErrNilAccountsAdapter
	}
	if check.IfNil(args.TransactionsProcessor) {
		return nil, process.ErrNilTxProcessor
	}

	return &selectionSession{
		accountsAdapter:        args.AccountsAdapter,
		transactionsProcessor:  args.TransactionsProcessor,
		ephemeralAccountsCache: make(map[string]vmcommon.AccountHandler),
	}, nil
}

// GetAccountState returns the state of an account.
// Will be called by mempool during transaction selection.
func (session *selectionSession) GetAccountState(address []byte) (*txcache.AccountState, error) {
	account, err := session.getExistingAccount(address)
	if err != nil {
		return nil, err
	}

	userAccount, ok := account.(state.UserAccountHandler)
	if !ok {
		return nil, process.ErrWrongTypeAssertion
	}

	return &txcache.AccountState{
		Nonce:   userAccount.GetNonce(),
		Balance: userAccount.GetBalance(),
	}, nil
}

func (session *selectionSession) getExistingAccount(address []byte) (vmcommon.AccountHandler, error) {
	account, ok := session.ephemeralAccountsCache[string(address)]
	if ok {
		return account, nil
	}

	account, err := session.accountsAdapter.GetExistingAccount(address)
	if err != nil {
		return nil, err
	}

	session.ephemeralAccountsCache[string(address)] = account
	return account, nil
}

// IsIncorrectlyGuarded checks if a transaction is incorrectly guarded (not executable).
// Will be called by mempool during transaction selection.
func (session *selectionSession) IsIncorrectlyGuarded(tx data.TransactionHandler) bool {
	address := tx.GetSndAddr()
	account, err := session.getExistingAccount(address)
	if err != nil {
		return false
	}

	userAccount, ok := account.(state.UserAccountHandler)
	if !ok {
		// On this branch, we are (approximately) mirroring the behavior of "transactionsProcessor.VerifyGuardian()".
		return true
	}

	txTyped, ok := tx.(*transaction.Transaction)
	if !ok {
		return false
	}

	err = session.transactionsProcessor.VerifyGuardian(txTyped, userAccount)
	return errors.Is(err, process.ErrTransactionNotExecutable)
}

// IsInterfaceNil returns true if there is no value under the interface
func (session *selectionSession) IsInterfaceNil() bool {
	return session == nil
}

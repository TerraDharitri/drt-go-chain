package mock

import (
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/TerraDharitri/drt-go-chain/state"
)

// ImportHandlerStub -
type ImportHandlerStub struct {
	ImportAllCalled               func() error
	GetValidatorAccountsDBCalled  func() state.AccountsAdapter
	GetMiniBlocksCalled           func() map[string]*block.MiniBlock
	GetHardForkMetaBlockCalled    func() data.MetaHeaderHandler
	GetUnFinishedMetaBlocksCalled func() map[string]data.MetaHeaderHandler
	GetTransactionsCalled         func() map[string]data.TransactionHandler
	GetAccountsDBForShardCalled   func(shardID uint32) state.AccountsAdapter
	CloseCalled                   func() error
}

// ImportAll -
func (ihs *ImportHandlerStub) ImportAll() error {
	if ihs.ImportAllCalled != nil {
		return ihs.ImportAllCalled()
	}
	return nil
}

// GetValidatorAccountsDB -
func (ihs *ImportHandlerStub) GetValidatorAccountsDB() state.AccountsAdapter {
	if ihs.GetValidatorAccountsDBCalled != nil {
		return ihs.GetValidatorAccountsDBCalled()
	}
	return nil
}

// GetMiniBlocks -
func (ihs *ImportHandlerStub) GetMiniBlocks() map[string]*block.MiniBlock {
	if ihs.GetMiniBlocksCalled != nil {
		return ihs.GetMiniBlocksCalled()
	}
	return nil
}

// GetHardForkMetaBlock -
func (ihs *ImportHandlerStub) GetHardForkMetaBlock() data.MetaHeaderHandler {
	if ihs.GetHardForkMetaBlockCalled != nil {
		return ihs.GetHardForkMetaBlockCalled()
	}
	return nil
}

// GetUnFinishedMetaBlocks -
func (ihs *ImportHandlerStub) GetUnFinishedMetaBlocks() map[string]data.MetaHeaderHandler {
	if ihs.GetUnFinishedMetaBlocksCalled != nil {
		return ihs.GetUnFinishedMetaBlocksCalled()
	}
	return nil
}

// GetTransactions -
func (ihs *ImportHandlerStub) GetTransactions() map[string]data.TransactionHandler {
	if ihs.GetTransactionsCalled != nil {
		return ihs.GetTransactionsCalled()
	}
	return nil
}

// GetAccountsDBForShard -
func (ihs *ImportHandlerStub) GetAccountsDBForShard(shardID uint32) state.AccountsAdapter {
	if ihs.GetAccountsDBForShardCalled != nil {
		return ihs.GetAccountsDBForShardCalled(shardID)
	}
	return nil
}

// Close -
func (ihs *ImportHandlerStub) Close() error {
	if ihs.CloseCalled != nil {
		return ihs.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (ihs *ImportHandlerStub) IsInterfaceNil() bool {
	return ihs == nil
}

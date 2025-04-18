package dataValidators

import (
	"fmt"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/sharding"
	"github.com/TerraDharitri/drt-go-chain/state"
)

var _ process.TxValidator = (*txValidator)(nil)

var log = logger.GetOrCreate("process/dataValidators")

// txValidator represents a tx handler validator that doesn't check the validity of provided txHandler
type txValidator struct {
	accounts             state.AccountsAdapter
	shardCoordinator     sharding.Coordinator
	whiteListHandler     process.WhiteListHandler
	pubKeyConverter      core.PubkeyConverter
	txVersionChecker     process.TxVersionCheckerHandler
	maxNonceDeltaAllowed int
}

// NewTxValidator creates a new nil tx handler validator instance
func NewTxValidator(
	accounts state.AccountsAdapter,
	shardCoordinator sharding.Coordinator,
	whiteListHandler process.WhiteListHandler,
	pubKeyConverter core.PubkeyConverter,
	txVersionChecker process.TxVersionCheckerHandler,
	maxNonceDeltaAllowed int,
) (*txValidator, error) {
	if check.IfNil(accounts) {
		return nil, process.ErrNilAccountsAdapter
	}
	if check.IfNil(shardCoordinator) {
		return nil, process.ErrNilShardCoordinator
	}
	if check.IfNil(whiteListHandler) {
		return nil, process.ErrNilWhiteListHandler
	}
	if check.IfNil(pubKeyConverter) {
		return nil, fmt.Errorf("%w in NewTxValidator", process.ErrNilPubkeyConverter)
	}
	if check.IfNil(txVersionChecker) {
		return nil, process.ErrNilTransactionVersionChecker
	}

	return &txValidator{
		accounts:             accounts,
		shardCoordinator:     shardCoordinator,
		whiteListHandler:     whiteListHandler,
		maxNonceDeltaAllowed: maxNonceDeltaAllowed,
		pubKeyConverter:      pubKeyConverter,
		txVersionChecker:     txVersionChecker,
	}, nil
}

// CheckTxValidity will filter transactions that needs to be added in pools
func (txv *txValidator) CheckTxValidity(interceptedTx process.InterceptedTransactionHandler) error {
	interceptedData, ok := interceptedTx.(process.InterceptedData)
	if ok {
		if txv.whiteListHandler.IsWhiteListed(interceptedData) {
			return nil
		}
	}

	if txv.isSenderInDifferentShard(interceptedTx) {
		return nil
	}

	// for relayed v3, we allow sender accounts that do not exist
	isRelayedV3 := common.IsRelayedTxV3(interceptedTx.Transaction())
	hasValue := hasTxValue(interceptedTx)
	shouldAllowMissingSenderAccount := isRelayedV3 && !hasValue
	accountHandler, err := txv.getSenderAccount(interceptedTx)
	if err != nil && !shouldAllowMissingSenderAccount {
		return err
	}

	return txv.checkAccount(interceptedTx, accountHandler)
}

func hasTxValue(interceptedTx process.InterceptedTransactionHandler) bool {
	txValue := interceptedTx.Transaction().GetValue()
	if check.IfNilReflect(txValue) {
		return false
	}

	return txValue.Sign() > 0
}

func (txv *txValidator) checkAccount(
	interceptedTx process.InterceptedTransactionHandler,
	accountHandler vmcommon.AccountHandler,
) error {
	err := txv.checkNonce(interceptedTx, accountHandler)
	if err != nil {
		return err
	}

	feePayerAccount, err := txv.getFeePayerAccount(interceptedTx, accountHandler)
	if err != nil {
		return err
	}

	return txv.checkBalance(interceptedTx, feePayerAccount)
}

func (txv *txValidator) getFeePayerAccount(
	interceptedTx process.InterceptedTransactionHandler,
	accountHandler vmcommon.AccountHandler,
) (state.UserAccountHandler, error) {
	payerAddress := interceptedTx.SenderAddress()
	payerAccount := accountHandler

	tx := interceptedTx.Transaction()
	if common.IsRelayedTxV3(tx) {
		relayedTx := tx.(data.RelayedTransactionHandler)
		payerAddress = relayedTx.GetRelayerAddr()
		relayerAccount, err := txv.accounts.GetExistingAccount(payerAddress)
		if err != nil {
			return nil, fmt.Errorf("%w for address %s and shard %d, err: %s",
				process.ErrAccountNotFound,
				txv.pubKeyConverter.SilentEncode(payerAddress, log),
				txv.shardCoordinator.SelfId(),
				err.Error(),
			)
		}
		payerAccount = relayerAccount
	}
	account, ok := payerAccount.(state.UserAccountHandler)
	if !ok {
		return nil, fmt.Errorf("%w, account is not of type *state.Account, address: %s",
			process.ErrWrongTypeAssertion,
			txv.pubKeyConverter.SilentEncode(payerAddress, log),
		)
	}
	return account, nil
}

func (txv *txValidator) checkBalance(interceptedTx process.InterceptedTransactionHandler, account state.UserAccountHandler) error {
	accountBalance := account.GetBalance()
	txFee := interceptedTx.Fee()
	if accountBalance.Cmp(txFee) < 0 {
		return fmt.Errorf("%w, for address: %s, wanted %v, have %v",
			process.ErrInsufficientFunds,
			txv.pubKeyConverter.SilentEncode(account.AddressBytes(), log),
			txFee,
			accountBalance,
		)
	}

	return nil
}

func (txv *txValidator) checkNonce(interceptedTx process.InterceptedTransactionHandler, accountHandler vmcommon.AccountHandler) error {
	accountNonce := uint64(0)
	if !check.IfNil(accountHandler) {
		accountNonce = accountHandler.GetNonce()
	}

	txNonce := interceptedTx.Nonce()
	lowerNonceInTx := txNonce < accountNonce
	veryHighNonceInTx := txNonce > accountNonce+uint64(txv.maxNonceDeltaAllowed)
	if lowerNonceInTx || veryHighNonceInTx {
		return fmt.Errorf("%w lowerNonceInTx: %v, veryHighNonceInTx: %v",
			process.ErrWrongTransaction,
			lowerNonceInTx,
			veryHighNonceInTx,
		)
	}
	return nil
}

func (txv *txValidator) isSenderInDifferentShard(interceptedTx process.InterceptedTransactionHandler) bool {
	shardID := txv.shardCoordinator.SelfId()
	txShardID := interceptedTx.SenderShardId()
	return shardID != txShardID
}

func (txv *txValidator) getSenderAccount(interceptedTx process.InterceptedTransactionHandler) (vmcommon.AccountHandler, error) {
	senderAddress := interceptedTx.SenderAddress()
	accountHandler, err := txv.accounts.GetExistingAccount(senderAddress)
	if err != nil {
		return nil, fmt.Errorf("%w for address %s and shard %d, err: %s",
			process.ErrAccountNotFound,
			txv.pubKeyConverter.SilentEncode(senderAddress, log),
			txv.shardCoordinator.SelfId(),
			err.Error(),
		)
	}

	return accountHandler, nil
}

func getTxData(interceptedTx process.InterceptedTransactionHandler) ([]byte, error) {
	tx := interceptedTx.Transaction()
	if tx == nil {
		return nil, process.ErrNilTransaction
	}

	return tx.GetData(), nil
}

// CheckTxWhiteList will check if the cross shard transactions are whitelisted and could be added in pools
func (txv *txValidator) CheckTxWhiteList(data process.InterceptedData) error {
	interceptedTx, ok := data.(process.InterceptedTransactionHandler)
	if !ok {
		return process.ErrWrongTypeAssertion
	}

	isTxCrossShardDestMe := interceptedTx.SenderShardId() != txv.shardCoordinator.SelfId() &&
		interceptedTx.ReceiverShardId() == txv.shardCoordinator.SelfId()
	if !isTxCrossShardDestMe {
		return nil
	}

	if txv.whiteListHandler.IsWhiteListed(data) {
		return nil
	}

	return process.ErrTransactionIsNotWhitelisted
}

// IsInterfaceNil returns true if there is no value under the interface
func (txv *txValidator) IsInterfaceNil() bool {
	return txv == nil
}

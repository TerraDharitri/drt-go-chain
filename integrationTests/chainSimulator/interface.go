package chainSimulator

import (
	"math/big"

	"github.com/TerraDharitri/drt-go-chain/node/chainSimulator/dtos"
	"github.com/TerraDharitri/drt-go-chain/node/chainSimulator/process"

	"github.com/TerraDharitri/drt-go-chain-core/data/api"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	crypto "github.com/TerraDharitri/drt-go-chain-crypto"
)

// ChainSimulator defines the operations for an entity that can simulate operations of a chain
type ChainSimulator interface {
	GenerateBlocks(numOfBlocks int) error
	GenerateBlocksUntilEpochIsReached(targetEpoch int32) error
	AddValidatorKeys(validatorsPrivateKeys [][]byte) error
	GetNodeHandler(shardID uint32) process.NodeHandler
	RemoveAccounts(addresses []string) error
	SendTxAndGenerateBlockTilTxIsExecuted(txToSend *transaction.Transaction, maxNumOfBlockToGenerateWhenExecutingTx int) (*transaction.ApiTransactionResult, error)
	SendTxsAndGenerateBlocksTilAreExecuted(txsToSend []*transaction.Transaction, maxNumOfBlocksToGenerateWhenExecutingTx int) ([]*transaction.ApiTransactionResult, error)
	SetStateMultiple(stateSlice []*dtos.AddressState) error
	GenerateAndMintWalletAddress(targetShardID uint32, value *big.Int) (dtos.WalletAddress, error)
	GetInitialWalletKeys() *dtos.InitialWalletKeys
	GetAccount(address dtos.WalletAddress) (api.AccountResponse, error)
	ForceResetValidatorStatisticsCache() error
	GetValidatorPrivateKeys() []crypto.PrivateKey
	SetKeyValueForAddress(address string, keyValueMap map[string]string) error
	Close()
}

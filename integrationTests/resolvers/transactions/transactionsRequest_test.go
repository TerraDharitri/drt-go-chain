package transactions

import (
	"math/big"
	"sync"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing"
	ed255192 "github.com/TerraDharitri/drt-go-chain-crypto/signing/ed25519"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/integrationTests"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/resolvers"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/state"
)

func TestTransactionsRequestsShouldWorkForHigherMaxTxNonceDeltaAllowed(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	numTxs := common.MaxTxNonceDeltaAllowed * 3
	mutMap := sync.Mutex{}
	txHashesMap := make(map[string]struct{})

	rm := resolvers.NewReceiverMonitor(t)
	shardIdResolver := uint32(0)
	shardIdRequester := uint32(0)
	nResolver, nRequester := resolvers.CreateResolverRequester(shardIdResolver, shardIdRequester)
	defer func() {
		nRequester.Close()
		nResolver.Close()
	}()

	nRequester.DataPool.Transactions().RegisterOnAdded(func(key []byte, value interface{}) {
		hash := string(key)

		mutMap.Lock()
		txHashesMap[hash] = struct{}{}
		if len(txHashesMap) == numTxs {
			rm.Done()
		}
		mutMap.Unlock()
	})

	txHashes := make([][]byte, 0, numTxs)
	txSuite := ed255192.NewEd25519()
	txKeyGen := signing.NewKeyGenerator(txSuite)
	sk, pk := txKeyGen.GeneratePair()
	senderBytes, _ := pk.ToByteArray()
	for nResolver.ShardCoordinator.ComputeId(senderBytes) != shardIdResolver {
		sk, pk = txKeyGen.GeneratePair()
		senderBytes, _ = pk.ToByteArray()
	}

	cacheId := process.ShardCacherIdentifier(shardIdRequester, shardIdResolver)
	for i := 0; i < numTxs; i++ {
		tx := integrationTests.GenerateTransferTx(
			uint64(i),
			sk,
			pk,
			big.NewInt(0),
			integrationTests.MinTxGasPrice,
			integrationTests.MinTxGasLimit,
			integrationTests.ChainID,
			1,
		)

		txHash, _ := core.CalculateHash(integrationTests.TestMarshalizer, integrationTests.TestHasher, tx)
		nResolver.DataPool.Transactions().AddData(txHash, tx, 0, cacheId)
		txHashes = append(txHashes, txHash)
	}

	account, _ := nRequester.AccntState.LoadAccount(senderBytes)
	userAccount := account.(state.UserAccountHandler)
	_ = userAccount.AddToBalance(big.NewInt(1000))
	_ = nRequester.AccntState.SaveAccount(account)
	_, _ = nRequester.AccntState.Commit()

	nRequester.RequestHandler.RequestTransaction(shardIdResolver, txHashes)

	rm.WaitWithTimeout()
}

package track

import (
	"github.com/TerraDharitri/drt-go-chain/dataRetriever"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/sharding"
	"github.com/TerraDharitri/drt-go-chain/storage"
	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
)

type miniBlockTrack struct {
	blockTransactionsPool    dataRetriever.ShardedDataCacherNotifier
	rewardTransactionsPool   dataRetriever.ShardedDataCacherNotifier
	unsignedTransactionsPool dataRetriever.ShardedDataCacherNotifier
	miniBlocksPool           storage.Cacher
	shardCoordinator         sharding.Coordinator
	whitelistHandler         process.WhiteListHandler
}

// NewMiniBlockTrack creates an object for tracking the received mini blocks
func NewMiniBlockTrack(
	dataPool dataRetriever.PoolsHolder,
	shardCoordinator sharding.Coordinator,
	whitelistHandler process.WhiteListHandler,
) (*miniBlockTrack, error) {

	if check.IfNil(dataPool) {
		return nil, process.ErrNilPoolsHolder
	}
	if check.IfNil(dataPool.Transactions()) {
		return nil, process.ErrNilTransactionPool
	}
	if check.IfNil(dataPool.RewardTransactions()) {
		return nil, process.ErrNilRewardTxDataPool
	}
	if check.IfNil(dataPool.UnsignedTransactions()) {
		return nil, process.ErrNilUnsignedTxDataPool
	}
	if check.IfNil(dataPool.MiniBlocks()) {
		return nil, process.ErrNilMiniBlockPool
	}
	if check.IfNil(shardCoordinator) {
		return nil, process.ErrNilShardCoordinator
	}
	if check.IfNil(whitelistHandler) {
		return nil, process.ErrNilWhiteListHandler
	}

	mbt := miniBlockTrack{
		blockTransactionsPool:    dataPool.Transactions(),
		rewardTransactionsPool:   dataPool.RewardTransactions(),
		unsignedTransactionsPool: dataPool.UnsignedTransactions(),
		miniBlocksPool:           dataPool.MiniBlocks(),
		shardCoordinator:         shardCoordinator,
		whitelistHandler:         whitelistHandler,
	}

	mbt.miniBlocksPool.RegisterHandler(mbt.receivedMiniBlock, core.UniqueIdentifier())

	return &mbt, nil
}

func (mbt *miniBlockTrack) receivedMiniBlock(key []byte, value interface{}) {
	if key == nil {
		return
	}

	miniBlock, ok := value.(*block.MiniBlock)
	if !ok {
		log.Warn("miniBlockTrack.receivedMiniBlock", "error", process.ErrWrongTypeAssertion)
		return
	}

	log.Debug("received miniblock from network in block tracker",
		"hash", key,
		"sender", miniBlock.SenderShardID,
		"receiver", miniBlock.ReceiverShardID,
		"type", miniBlock.Type,
		"num txs", len(miniBlock.TxHashes))

	if miniBlock.SenderShardID == mbt.shardCoordinator.SelfId() {
		return
	}

	// TODO - stop reusing miniBlock.TxHashes for peer changes, add new fields
	transactionPool := mbt.getTransactionPool(miniBlock.Type)
	if check.IfNil(transactionPool) {
		return
	}

	mbt.whitelistHandler.Add(miniBlock.TxHashes)

	strCache := process.ShardCacherIdentifier(miniBlock.SenderShardID, miniBlock.ReceiverShardID)
	transactionPool.ImmunizeSetOfDataAgainstEviction(miniBlock.TxHashes, strCache)
}

func (mbt *miniBlockTrack) getTransactionPool(mbType block.Type) dataRetriever.ShardedDataCacherNotifier {
	switch mbType {
	case block.TxBlock:
		return mbt.blockTransactionsPool
	case block.RewardsBlock:
		return mbt.rewardTransactionsPool
	case block.SmartContractResultBlock:
		return mbt.unsignedTransactionsPool
	}

	return nil
}

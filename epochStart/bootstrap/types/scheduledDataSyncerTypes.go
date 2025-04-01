package types

import (
	"github.com/TerraDharitri/drt-go-chain/epochStart"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/update"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
)

// ScheduledDataSyncerCreateArgs holds the arguments to create a scheduled data syncer factory
type ScheduledDataSyncerCreateArgs struct {
	ScheduledTxsHandler  process.ScheduledTxsExecutionHandler
	HeadersSyncer        epochStart.HeadersByHashSyncer
	MiniBlocksSyncer     epochStart.PendingMiniBlocksSyncHandler
	TxSyncer             update.TransactionsSyncHandler
	ScheduledEnableEpoch uint32
}

// ScheduledDataSyncerCreator is the interface implemented by the scheduled data syncer factory allowing to create
// scheduled data syncer instances
type ScheduledDataSyncerCreator interface {
	Create(args *ScheduledDataSyncerCreateArgs) (ScheduledDataSyncer, error)
	IsInterfaceNil() bool
}

// ScheduledDataSyncer interface allows to synchronize the correct headers and root hash with or without scheduled sc calls feature activated.
type ScheduledDataSyncer interface {
	UpdateSyncDataIfNeeded(notarizedShardHeader data.ShardHeaderHandler) (data.ShardHeaderHandler, map[string]data.HeaderHandler, map[string]*block.MiniBlock, error)
	GetRootHashToSync(notarizedShardHeader data.ShardHeaderHandler) []byte
	IsInterfaceNil() bool
}

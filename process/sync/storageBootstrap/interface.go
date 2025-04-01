package storageBootstrap

import (
	"github.com/TerraDharitri/drt-go-chain/process/block/bootstrapStorage"
	"github.com/TerraDharitri/drt-go-chain-core/data"
)

// StorageBootstrapper is the main interface for bootstrap from storage execution engine
type storageBootstrapperHandler interface {
	getHeader(hash []byte) (data.HeaderHandler, error)
	getHeaderWithNonce(nonce uint64, shardID uint32) (data.HeaderHandler, []byte, error)
	applyCrossNotarizedHeaders(crossNotarizedHeaders []bootstrapStorage.BootstrapHeaderInfo) error
	applyNumPendingMiniBlocks(pendingMiniBlocks []bootstrapStorage.PendingMiniBlocksInfo)
	applySelfNotarizedHeaders(selfNotarizedHeaders []bootstrapStorage.BootstrapHeaderInfo) ([]data.HeaderHandler, [][]byte, error)
	cleanupNotarizedStorage(hash []byte)
	cleanupNotarizedStorageForHigherNoncesIfExist(crossNotarizedHeaders []bootstrapStorage.BootstrapHeaderInfo)
	getRootHash(hash []byte) []byte
	IsInterfaceNil() bool
}

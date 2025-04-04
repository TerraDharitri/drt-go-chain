package outport

import (
	outportcore "github.com/TerraDharitri/drt-go-chain-core/data/outport"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	"github.com/TerraDharitri/drt-go-chain/outport/process"
)

// Driver is an interface for saving node specific data to other storage.
// This could be an elastic search index, a MySql database or any other external services.
type Driver interface {
	SaveBlock(outportBlock *outportcore.OutportBlock) error
	RevertIndexedBlock(blockData *outportcore.BlockData) error
	SaveRoundsInfo(roundsInfos *outportcore.RoundsInfo) error
	SaveValidatorsPubKeys(validatorsPubKeys *outportcore.ValidatorsPubKeys) error
	SaveValidatorsRating(validatorsRating *outportcore.ValidatorsRating) error
	SaveAccounts(accounts *outportcore.Accounts) error
	FinalizedBlock(finalizedBlock *outportcore.FinalizedBlock) error
	GetMarshaller() marshal.Marshalizer
	SetCurrentSettings(config outportcore.OutportConfig) error
	RegisterHandler(handlerFunction func() error, topic string) error
	Close() error
	IsInterfaceNil() bool
}

// OutportHandler is interface that defines what a proxy implementation should be able to do
// The node is able to talk only with this interface
type OutportHandler interface {
	SaveBlock(outportBlock *outportcore.OutportBlockWithHeaderAndBody) error
	RevertIndexedBlock(blockData *outportcore.HeaderDataWithBody) error
	SaveRoundsInfo(roundsInfos *outportcore.RoundsInfo)
	SaveValidatorsPubKeys(validatorsPubKeys *outportcore.ValidatorsPubKeys)
	SaveValidatorsRating(validatorsRating *outportcore.ValidatorsRating)
	SaveAccounts(accounts *outportcore.Accounts)
	FinalizedBlock(finalizedBlock *outportcore.FinalizedBlock)
	SubscribeDriver(driver Driver) error
	HasDrivers() bool
	Close() error
	IsInterfaceNil() bool
}

// DataProviderOutport is an interface that defines what an implementation of data provider outport should be able to do
type DataProviderOutport interface {
	PrepareOutportSaveBlockData(arg process.ArgPrepareOutportSaveBlockData) (*outportcore.OutportBlockWithHeaderAndBody, error)
	IsInterfaceNil() bool
}

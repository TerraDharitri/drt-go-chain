package block

import (
	"sync"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/TerraDharitri/drt-go-chain-core/data/scheduled"
	"github.com/TerraDharitri/drt-go-chain-core/hashing"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"

	"github.com/TerraDharitri/drt-go-chain/dataRetriever"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/process/block/bootstrapStorage"
	"github.com/TerraDharitri/drt-go-chain/process/block/processedMb"
	"github.com/TerraDharitri/drt-go-chain/process/mock"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/dblookupext"
	"github.com/TerraDharitri/drt-go-chain/testscommon/enableEpochsHandlerMock"
	"github.com/TerraDharitri/drt-go-chain/testscommon/epochNotifier"
	"github.com/TerraDharitri/drt-go-chain/testscommon/factory"
	"github.com/TerraDharitri/drt-go-chain/testscommon/hashingMocks"
	"github.com/TerraDharitri/drt-go-chain/testscommon/outport"
	"github.com/TerraDharitri/drt-go-chain/testscommon/shardingMocks"
	stateMock "github.com/TerraDharitri/drt-go-chain/testscommon/state"
	statusHandlerMock "github.com/TerraDharitri/drt-go-chain/testscommon/statusHandler"
	storageStubs "github.com/TerraDharitri/drt-go-chain/testscommon/storage"
)

func (bp *baseProcessor) ComputeHeaderHash(hdr data.HeaderHandler) ([]byte, error) {
	return core.CalculateHash(bp.marshalizer, bp.hasher, hdr)
}

func (bp *baseProcessor) VerifyStateRoot(rootHash []byte) bool {
	return bp.verifyStateRoot(rootHash)
}

func (bp *baseProcessor) CheckBlockValidity(
	headerHandler data.HeaderHandler,
	bodyHandler data.BodyHandler,
) error {
	return bp.checkBlockValidity(headerHandler, bodyHandler)
}

func (bp *baseProcessor) RemoveHeadersBehindNonceFromPools(
	shouldRemoveBlockBody bool,
	shardId uint32,
	nonce uint64,
) {
	bp.removeHeadersBehindNonceFromPools(shouldRemoveBlockBody, shardId, nonce)
}

func (bp *baseProcessor) GetPruningHandler(finalHeaderNonce uint64) state.PruningHandler {
	return bp.getPruningHandler(finalHeaderNonce)
}

func (bp *baseProcessor) SetLastRestartNonce(lastRestartNonce uint64) {
	bp.lastRestartNonce = lastRestartNonce
}

func (bp *baseProcessor) CommitTrieEpochRootHashIfNeeded(metaBlock *block.MetaBlock, rootHash []byte) error {
	return bp.commitTrieEpochRootHashIfNeeded(metaBlock, rootHash)
}

func (sp *shardProcessor) ReceivedMetaBlock(header data.HeaderHandler, metaBlockHash []byte) {
	sp.receivedMetaBlock(header, metaBlockHash)
}

func (sp *shardProcessor) CreateMiniBlocks(haveTime func() bool) (*block.Body, map[string]*processedMb.ProcessedMiniBlockInfo, error) {
	return sp.createMiniBlocks(haveTime, []byte("random"))
}

func (sp *shardProcessor) GetOrderedProcessedMetaBlocksFromHeader(header data.HeaderHandler) ([]data.HeaderHandler, error) {
	return sp.getOrderedProcessedMetaBlocksFromHeader(header)
}

func (sp *shardProcessor) UpdateCrossShardInfo(processedMetaHdrs []data.HeaderHandler) error {
	return sp.updateCrossShardInfo(processedMetaHdrs)
}

func (sp *shardProcessor) UpdateStateStorage(finalHeaders []data.HeaderHandler, currentHeader data.HeaderHandler) {
	currShardHeader, ok := currentHeader.(data.ShardHeaderHandler)
	if !ok {
		return
	}
	sp.updateState(finalHeaders, currShardHeader)
}

func NewShardProcessorEmptyWith3shards(
	tdp dataRetriever.PoolsHolder,
	genesisBlocks map[uint32]data.HeaderHandler,
	blockChain data.ChainHandler,
) (*shardProcessor, error) {
	shardCoordinator := mock.NewMultiShardsCoordinatorMock(3)
	nodesCoordinator := shardingMocks.NewNodesCoordinatorMock()

	argsHeaderValidator := ArgsHeaderValidator{
		Hasher:      &hashingMocks.HasherMock{},
		Marshalizer: &mock.MarshalizerMock{},
	}
	hdrValidator, _ := NewHeaderValidator(argsHeaderValidator)

	accountsDb := make(map[state.AccountsDbIdentifier]state.AccountsAdapter)
	accountsDb[state.UserAccountsState] = &stateMock.AccountsStub{}

	coreComponents := &mock.CoreComponentsMock{
		IntMarsh:                  &mock.MarshalizerMock{},
		Hash:                      &hashingMocks.HasherMock{},
		UInt64ByteSliceConv:       &mock.Uint64ByteSliceConverterMock{},
		StatusField:               &statusHandlerMock.AppStatusHandlerStub{},
		RoundField:                &mock.RoundHandlerMock{},
		ProcessStatusHandlerField: &testscommon.ProcessStatusHandlerStub{},
		EpochNotifierField:        &epochNotifier.EpochNotifierStub{},
		EnableEpochsHandlerField:  enableEpochsHandlerMock.NewEnableEpochsHandlerStub(),
		RoundNotifierField:        &epochNotifier.RoundNotifierStub{},
		EnableRoundsHandlerField:  &testscommon.EnableRoundsHandlerStub{},
	}
	dataComponents := &mock.DataComponentsMock{
		Storage:    &storageStubs.ChainStorerStub{},
		DataPool:   tdp,
		BlockChain: blockChain,
	}
	boostrapComponents := &mock.BootstrapComponentsMock{
		Coordinator:          shardCoordinator,
		HdrIntegrityVerifier: &mock.HeaderIntegrityVerifierStub{},
		VersionedHdrFactory:  &testscommon.VersionedHeaderFactoryStub{},
	}
	statusComponents := &mock.StatusComponentsMock{
		Outport: &outport.OutportStub{},
	}
	statusCoreComponents := &factory.StatusCoreComponentsStub{
		AppStatusHandlerField: &statusHandlerMock.AppStatusHandlerStub{},
	}

	arguments := ArgShardProcessor{
		ArgBaseProcessor: ArgBaseProcessor{
			CoreComponents:       coreComponents,
			DataComponents:       dataComponents,
			BootstrapComponents:  boostrapComponents,
			StatusComponents:     statusComponents,
			StatusCoreComponents: statusCoreComponents,
			AccountsDB:           accountsDb,
			ForkDetector:         &mock.ForkDetectorMock{},
			NodesCoordinator:     nodesCoordinator,
			FeeHandler:           &mock.FeeAccumulatorStub{},
			RequestHandler:       &testscommon.RequestHandlerStub{},
			BlockChainHook:       &testscommon.BlockChainHookStub{},
			TxCoordinator:        &testscommon.TransactionCoordinatorMock{},
			EpochStartTrigger:    &mock.EpochStartTriggerStub{},
			HeaderValidator:      hdrValidator,
			BootStorer: &mock.BoostrapStorerMock{
				PutCalled: func(round int64, bootData bootstrapStorage.BootstrapData) error {
					return nil
				},
			},
			BlockTracker:                 mock.NewBlockTrackerMock(shardCoordinator, genesisBlocks),
			BlockSizeThrottler:           &mock.BlockSizeThrottlerStub{},
			Version:                      "softwareVersion",
			HistoryRepository:            &dblookupext.HistoryRepositoryStub{},
			GasHandler:                   &mock.GasHandlerMock{},
			OutportDataProvider:          &outport.OutportDataProviderStub{},
			ScheduledTxsExecutionHandler: &testscommon.ScheduledTxsExecutionStub{},
			ProcessedMiniBlocksTracker:   &testscommon.ProcessedMiniBlocksTrackerStub{},
			ReceiptsRepository:           &testscommon.ReceiptsRepositoryStub{},
			BlockProcessingCutoffHandler: &testscommon.BlockProcessingCutoffStub{},
			ManagedPeersHolder:           &testscommon.ManagedPeersHolderStub{},
			SentSignaturesTracker:        &testscommon.SentSignatureTrackerStub{},
		},
	}
	shardProc, err := NewShardProcessor(arguments)
	return shardProc, err
}

func (mp *metaProcessor) RequestBlockHeaders(header *block.MetaBlock) (uint32, uint32) {
	return mp.requestShardHeaders(header)
}

func (mp *metaProcessor) ReceivedShardHeader(header data.HeaderHandler, shardHeaderHash []byte) {
	mp.receivedShardHeader(header, shardHeaderHash)
}

func (mp *metaProcessor) GetDataPool() dataRetriever.PoolsHolder {
	return mp.dataPool
}

func (mp *metaProcessor) AddHdrHashToRequestedList(hdr data.HeaderHandler, hdrHash []byte) {
	mp.hdrsForCurrBlock.mutHdrsForBlock.Lock()
	defer mp.hdrsForCurrBlock.mutHdrsForBlock.Unlock()

	if mp.hdrsForCurrBlock.hdrHashAndInfo == nil {
		mp.hdrsForCurrBlock.hdrHashAndInfo = make(map[string]*hdrInfo)
	}

	if mp.hdrsForCurrBlock.highestHdrNonce == nil {
		mp.hdrsForCurrBlock.highestHdrNonce = make(map[uint32]uint64, mp.shardCoordinator.NumberOfShards())
	}

	mp.hdrsForCurrBlock.hdrHashAndInfo[string(hdrHash)] = &hdrInfo{hdr: hdr, usedInBlock: true}
	mp.hdrsForCurrBlock.missingHdrs++
}

func (mp *metaProcessor) IsHdrMissing(hdrHash []byte) bool {
	mp.hdrsForCurrBlock.mutHdrsForBlock.RLock()
	defer mp.hdrsForCurrBlock.mutHdrsForBlock.RUnlock()

	hdrInfoValue, ok := mp.hdrsForCurrBlock.hdrHashAndInfo[string(hdrHash)]
	if !ok {
		return true
	}

	return check.IfNil(hdrInfoValue.hdr)
}

func (mp *metaProcessor) CreateShardInfo() ([]data.ShardDataHandler, error) {
	return mp.createShardInfo()
}

func (mp *metaProcessor) RequestMissingFinalityAttestingShardHeaders() uint32 {
	mp.hdrsForCurrBlock.mutHdrsForBlock.Lock()
	defer mp.hdrsForCurrBlock.mutHdrsForBlock.Unlock()

	return mp.requestMissingFinalityAttestingShardHeaders()
}

func (mp *metaProcessor) SaveMetricCrossCheckBlockHeight() {
	mp.saveMetricCrossCheckBlockHeight()
}

func (bp *baseProcessor) NotarizedHdrs() map[uint32][]data.HeaderHandler {
	lastCrossNotarizedHeaders := make(map[uint32][]data.HeaderHandler)
	for shardID := uint32(0); shardID < bp.shardCoordinator.NumberOfShards(); shardID++ {
		lastCrossNotarizedHeaderForShard := bp.LastNotarizedHdrForShard(shardID)
		if !check.IfNil(lastCrossNotarizedHeaderForShard) {
			lastCrossNotarizedHeaders[shardID] = append(lastCrossNotarizedHeaders[shardID], lastCrossNotarizedHeaderForShard)
		}
	}

	lastCrossNotarizedHeaderForShard := bp.LastNotarizedHdrForShard(core.MetachainShardId)
	if !check.IfNil(lastCrossNotarizedHeaderForShard) {
		lastCrossNotarizedHeaders[core.MetachainShardId] = append(lastCrossNotarizedHeaders[core.MetachainShardId], lastCrossNotarizedHeaderForShard)
	}

	return lastCrossNotarizedHeaders
}

func (bp *baseProcessor) LastNotarizedHdrForShard(shardID uint32) data.HeaderHandler {
	lastCrossNotarizedHeaderForShard, _, _ := bp.blockTracker.GetLastCrossNotarizedHeader(shardID)
	if check.IfNil(lastCrossNotarizedHeaderForShard) {
		return nil
	}

	return lastCrossNotarizedHeaderForShard
}

func (bp *baseProcessor) SetMarshalizer(marshal marshal.Marshalizer) {
	bp.marshalizer = marshal
}

func (bp *baseProcessor) SetHasher(hasher hashing.Hasher) {
	bp.hasher = hasher
}

func (bp *baseProcessor) SetHeaderValidator(validator process.HeaderConstructionValidator) {
	bp.headerValidator = validator
}

func (bp *baseProcessor) RequestHeadersIfMissing(sortedHdrs []data.HeaderHandler, shardId uint32) error {
	return bp.requestHeadersIfMissing(sortedHdrs, shardId)
}

func (mp *metaProcessor) SetShardBlockFinality(val uint32) {
	mp.hdrsForCurrBlock.mutHdrsForBlock.Lock()
	mp.shardBlockFinality = val
	mp.hdrsForCurrBlock.mutHdrsForBlock.Unlock()
}

func (mp *metaProcessor) SaveLastNotarizedHeader(header *block.MetaBlock) error {
	return mp.saveLastNotarizedHeader(header)
}

func (mp *metaProcessor) CheckShardHeadersValidity(header *block.MetaBlock) (map[uint32]data.HeaderHandler, error) {
	return mp.checkShardHeadersValidity(header)
}

func (mp *metaProcessor) CheckShardHeadersFinality(highestNonceHdrs map[uint32]data.HeaderHandler) error {
	return mp.checkShardHeadersFinality(highestNonceHdrs)
}

func (mp *metaProcessor) CheckHeaderBodyCorrelation(hdr data.HeaderHandler, body *block.Body) error {
	return mp.checkHeaderBodyCorrelation(hdr.GetMiniBlockHeaderHandlers(), body)
}

func (bp *baseProcessor) IsHdrConstructionValid(currHdr, prevHdr data.HeaderHandler) error {
	return bp.headerValidator.IsHeaderConstructionValid(currHdr, prevHdr)
}

func (mp *metaProcessor) ChRcvAllHdrs() chan bool {
	return mp.chRcvAllHdrs
}

func (mp *metaProcessor) UpdateShardsHeadersNonce(key uint32, value uint64) {
	mp.updateShardHeadersNonce(key, value)
}

func (mp *metaProcessor) GetShardsHeadersNonce() *sync.Map {
	return mp.shardsHeadersNonce
}

func (sp *shardProcessor) SaveLastNotarizedHeader(shardId uint32, processedHdrs []data.HeaderHandler) error {
	return sp.saveLastNotarizedHeader(shardId, processedHdrs)
}

func (sp *shardProcessor) CheckHeaderBodyCorrelation(hdr data.HeaderHandler, body *block.Body) error {
	return sp.checkHeaderBodyCorrelation(hdr.GetMiniBlockHeaderHandlers(), body)
}

func (sp *shardProcessor) CheckAndRequestIfMetaHeadersMissing() {
	sp.checkAndRequestIfMetaHeadersMissing()
}

func (sp *shardProcessor) GetHashAndHdrStruct(header data.HeaderHandler, hash []byte) *hashAndHdr {
	return &hashAndHdr{header, hash}
}

func (sp *shardProcessor) RequestMissingFinalityAttestingHeaders() uint32 {
	sp.hdrsForCurrBlock.mutHdrsForBlock.Lock()
	defer sp.hdrsForCurrBlock.mutHdrsForBlock.Unlock()

	return sp.requestMissingFinalityAttestingHeaders(
		core.MetachainShardId,
		sp.metaBlockFinality,
	)
}

func (sp *shardProcessor) CheckMetaHeadersValidityAndFinality() error {
	return sp.checkMetaHeadersValidityAndFinality()
}

func (sp *shardProcessor) CreateAndProcessMiniBlocksDstMe(
	haveTime func() bool,
) (block.MiniBlockSlice, uint32, uint32, error) {
	createAndProcessInfo, err := sp.createAndProcessMiniBlocksDstMe(haveTime)
	return createAndProcessInfo.miniBlocks, createAndProcessInfo.numHdrsAdded, createAndProcessInfo.numTxsAdded, err
}

func (sp *shardProcessor) DisplayLogInfo(
	header data.HeaderHandler,
	body *block.Body,
	headerHash []byte,
	numShards uint32,
	selfId uint32,
	dataPool dataRetriever.PoolsHolder,
	blockTracker process.BlockTracker,
) {
	sp.txCounter.displayLogInfo(header, body, headerHash, numShards, selfId, dataPool, blockTracker)
}

func (sp *shardProcessor) GetHighestHdrForOwnShardFromMetachain(processedHdrs []data.HeaderHandler) ([]data.HeaderHandler, [][]byte, error) {
	return sp.getHighestHdrForOwnShardFromMetachain(processedHdrs)
}

func (sp *shardProcessor) RestoreMetaBlockIntoPool(
	miniBlockHashes map[string]uint32,
	metaBlockHashes [][]byte,
	headerHandler data.HeaderHandler,
) error {
	return sp.restoreMetaBlockIntoPool(headerHandler, miniBlockHashes, metaBlockHashes)
}

func (sp *shardProcessor) GetAllMiniBlockDstMeFromMeta(
	header data.ShardHeaderHandler,
) (map[string][]byte, error) {
	return sp.getAllMiniBlockDstMeFromMeta(header)
}

func (bp *baseProcessor) SetHdrForCurrentBlock(headerHash []byte, headerHandler data.HeaderHandler, usedInBlock bool) {
	bp.hdrsForCurrBlock.mutHdrsForBlock.Lock()
	bp.hdrsForCurrBlock.hdrHashAndInfo[string(headerHash)] = &hdrInfo{hdr: headerHandler, usedInBlock: usedInBlock}
	bp.hdrsForCurrBlock.mutHdrsForBlock.Unlock()
}

func (bp *baseProcessor) SetHighestHdrNonceForCurrentBlock(shardId uint32, value uint64) {
	bp.hdrsForCurrBlock.mutHdrsForBlock.Lock()
	bp.hdrsForCurrBlock.highestHdrNonce[shardId] = value
	bp.hdrsForCurrBlock.mutHdrsForBlock.Unlock()
}

func (bp *baseProcessor) CreateBlockStarted() error {
	return bp.createBlockStarted()
}

func (sp *shardProcessor) AddProcessedCrossMiniBlocksFromHeader(header data.HeaderHandler) error {
	return sp.addProcessedCrossMiniBlocksFromHeader(header)
}

func (mp *metaProcessor) VerifyCrossShardMiniBlockDstMe(header *block.MetaBlock) error {
	return mp.verifyCrossShardMiniBlockDstMe(header)
}

func (mp *metaProcessor) ApplyBodyToHeader(metaHdr data.MetaHeaderHandler, body *block.Body) (data.BodyHandler, error) {
	return mp.applyBodyToHeader(metaHdr, body)
}

func (sp *shardProcessor) ApplyBodyToHeader(shardHdr data.ShardHeaderHandler, body *block.Body, processedMiniBlocksDestMeInfo map[string]*processedMb.ProcessedMiniBlockInfo) (*block.Body, error) {
	return sp.applyBodyToHeader(shardHdr, body, processedMiniBlocksDestMeInfo)
}

func (mp *metaProcessor) CreateBlockBody(metaBlock data.HeaderHandler, haveTime func() bool) (data.BodyHandler, error) {
	return mp.createBlockBody(metaBlock, haveTime)
}

func (sp *shardProcessor) CreateBlockBody(shardHdr data.HeaderHandler, haveTime func() bool) (data.BodyHandler, map[string]*processedMb.ProcessedMiniBlockInfo, error) {
	return sp.createBlockBody(shardHdr, haveTime)
}

func (sp *shardProcessor) CheckEpochCorrectnessCrossChain() error {
	return sp.checkEpochCorrectnessCrossChain()
}

func (sp *shardProcessor) CheckEpochCorrectness(header *block.Header) error {
	return sp.checkEpochCorrectness(header)
}

func (sp *shardProcessor) GetBootstrapHeadersInfo(
	selfNotarizedHeaders []data.HeaderHandler,
	selfNotarizedHeadersHashes [][]byte,
) []bootstrapStorage.BootstrapHeaderInfo {
	return sp.getBootstrapHeadersInfo(selfNotarizedHeaders, selfNotarizedHeadersHashes)
}

func (sp *shardProcessor) RequestMetaHeadersIfNeeded(hdrsAdded uint32, lastMetaHdr data.HeaderHandler) {
	sp.requestMetaHeadersIfNeeded(hdrsAdded, lastMetaHdr)
}

func (mp *metaProcessor) RequestShardHeadersIfNeeded(hdrsAddedForShard map[uint32]uint32, lastShardHdr map[uint32]data.HeaderHandler) {
	mp.requestShardHeadersIfNeeded(hdrsAddedForShard, lastShardHdr)
}

func (bp *baseProcessor) AddHeaderIntoTrackerPool(nonce uint64, shardID uint32) {
	bp.addHeaderIntoTrackerPool(nonce, shardID)
}

func (bp *baseProcessor) UpdateState(
	finalHeader data.HeaderHandler,
	rootHash []byte,
	prevRootHash []byte,
	accounts state.AccountsAdapter,
) {
	bp.updateStateStorage(finalHeader, rootHash, prevRootHash, accounts)
}

func GasAndFeesDelta(initialGasAndFees, finalGasAndFees scheduled.GasAndFees) scheduled.GasAndFees {
	return gasAndFeesDelta(initialGasAndFees, finalGasAndFees)
}

func (sp *shardProcessor) RequestEpochStartInfo(header data.ShardHeaderHandler, haveTime func() time.Duration) error {
	return sp.requestEpochStartInfo(header, haveTime)
}

func (mp *metaProcessor) ProcessEpochStartMetaBlock(
	header *block.MetaBlock,
	body *block.Body,
) error {
	return mp.processEpochStartMetaBlock(header, body)
}

func (mp *metaProcessor) UpdateEpochStartHeader(metaHdr *block.MetaBlock) error {
	return mp.updateEpochStartHeader(metaHdr)
}

func (mp *metaProcessor) CreateEpochStartBody(metaBlock *block.MetaBlock) (data.BodyHandler, error) {
	return mp.createEpochStartBody(metaBlock)
}

func (bp *baseProcessor) GetIndexOfFirstMiniBlockToBeExecuted(header data.HeaderHandler) int {
	return bp.getIndexOfFirstMiniBlockToBeExecuted(header)
}

func (bp *baseProcessor) GetFinalMiniBlocks(header data.HeaderHandler, body *block.Body) (*block.Body, error) {
	return bp.getFinalMiniBlocks(header, body)
}

func GetScheduledMiniBlocksFromMe(headerHandler data.HeaderHandler, bodyHandler data.BodyHandler) (block.MiniBlockSlice, error) {
	return getScheduledMiniBlocksFromMe(headerHandler, bodyHandler)
}

func (bp *baseProcessor) CheckScheduledMiniBlocksValidity(headerHandler data.HeaderHandler) error {
	return bp.checkScheduledMiniBlocksValidity(headerHandler)
}

func (bp *baseProcessor) SetMiniBlockHeaderReservedField(
	miniBlock *block.MiniBlock,
	miniBlockHeaderHandler data.MiniBlockHeaderHandler,
	processedMiniBlocksDestMeInfo map[string]*processedMb.ProcessedMiniBlockInfo,
) error {
	return bp.setMiniBlockHeaderReservedField(miniBlock, miniBlockHeaderHandler, processedMiniBlocksDestMeInfo)
}

func (mp *metaProcessor) GetFinalMiniBlockHeaders(miniBlockHeaderHandlers []data.MiniBlockHeaderHandler) []data.MiniBlockHeaderHandler {
	return mp.getFinalMiniBlockHeaders(miniBlockHeaderHandlers)
}

func CheckProcessorNilParameters(arguments ArgBaseProcessor) error {
	return checkProcessorParameters(arguments)
}

func (bp *baseProcessor) SetIndexOfFirstTxProcessed(miniBlockHeaderHandler data.MiniBlockHeaderHandler) error {
	return bp.setIndexOfFirstTxProcessed(miniBlockHeaderHandler)
}

func (bp *baseProcessor) SetIndexOfLastTxProcessed(
	miniBlockHeaderHandler data.MiniBlockHeaderHandler,
	processedMiniBlocksDestMeInfo map[string]*processedMb.ProcessedMiniBlockInfo,
) error {
	return bp.setIndexOfLastTxProcessed(miniBlockHeaderHandler, processedMiniBlocksDestMeInfo)
}

func (bp *baseProcessor) GetProcessedMiniBlocksTracker() process.ProcessedMiniBlocksTracker {
	return bp.processedMiniBlocksTracker
}

func (bp *baseProcessor) SetProcessingTypeAndConstructionStateForScheduledMb(
	miniBlockHeaderHandler data.MiniBlockHeaderHandler,
	processedMiniBlocksDestMeInfo map[string]*processedMb.ProcessedMiniBlockInfo,
) error {
	return bp.setProcessingTypeAndConstructionStateForScheduledMb(miniBlockHeaderHandler, processedMiniBlocksDestMeInfo)
}

func (bp *baseProcessor) SetProcessingTypeAndConstructionStateForNormalMb(
	miniBlockHeaderHandler data.MiniBlockHeaderHandler,
	processedMiniBlocksDestMeInfo map[string]*processedMb.ProcessedMiniBlockInfo,
) error {
	return bp.setProcessingTypeAndConstructionStateForNormalMb(miniBlockHeaderHandler, processedMiniBlocksDestMeInfo)
}

func (sp *shardProcessor) RollBackProcessedMiniBlockInfo(miniBlockHeader data.MiniBlockHeaderHandler, miniBlockHash []byte) {
	sp.rollBackProcessedMiniBlockInfo(miniBlockHeader, miniBlockHash)
}

func (sp *shardProcessor) SetProcessedMiniBlocksInfo(miniBlockHashes [][]byte, metaBlockHash string, metaBlock *block.MetaBlock) {
	sp.setProcessedMiniBlocksInfo(miniBlockHashes, metaBlockHash, metaBlock)
}

func (sp *shardProcessor) GetIndexOfLastTxProcessedInMiniBlock(miniBlockHash []byte, metaBlock *block.MetaBlock) int32 {
	return getIndexOfLastTxProcessedInMiniBlock(miniBlockHash, metaBlock)
}

func (sp *shardProcessor) RollBackProcessedMiniBlocksInfo(headerHandler data.HeaderHandler, mapMiniBlockHashes map[string]uint32) {
	sp.rollBackProcessedMiniBlocksInfo(headerHandler, mapMiniBlockHashes)
}

func (bp *baseProcessor) CheckConstructionStateAndIndexesCorrectness(mbh data.MiniBlockHeaderHandler) error {
	return checkConstructionStateAndIndexesCorrectness(mbh)
}

func (mp *metaProcessor) GetAllMarshalledTxs(body *block.Body) map[string][][]byte {
	return mp.getAllMarshalledTxs(body)
}

// SetNonceOfFirstCommittedBlock -
func (bp *baseProcessor) SetNonceOfFirstCommittedBlock(nonce uint64) {
	bp.setNonceOfFirstCommittedBlock(nonce)
}

// CheckSentSignaturesAtCommitTime -
func (bp *baseProcessor) CheckSentSignaturesAtCommitTime(header data.HeaderHandler) error {
	return bp.checkSentSignaturesAtCommitTime(header)
}

// GetHdrForBlock -
func (mp *metaProcessor) GetHdrForBlock() *hdrForBlock {
	return mp.hdrsForCurrBlock
}

// ChannelReceiveAllHeaders -
func (mp *metaProcessor) ChannelReceiveAllHeaders() chan bool {
	return mp.chRcvAllHdrs
}

// ComputeExistingAndRequestMissingShardHeaders -
func (mp *metaProcessor) ComputeExistingAndRequestMissingShardHeaders(metaBlock *block.MetaBlock) (uint32, uint32) {
	return mp.computeExistingAndRequestMissingShardHeaders(metaBlock)
}

// ComputeExistingAndRequestMissingMetaHeaders -
func (sp *shardProcessor) ComputeExistingAndRequestMissingMetaHeaders(header data.ShardHeaderHandler) (uint32, uint32) {
	return sp.computeExistingAndRequestMissingMetaHeaders(header)
}

// GetHdrForBlock -
func (sp *shardProcessor) GetHdrForBlock() *hdrForBlock {
	return sp.hdrsForCurrBlock
}

// ChannelReceiveAllHeaders -
func (sp *shardProcessor) ChannelReceiveAllHeaders() chan bool {
	return sp.chRcvAllMetaHdrs
}

// InitMaps -
func (hfb *hdrForBlock) InitMaps() {
	hfb.initMaps()
	hfb.resetMissingHdrs()
}

// Clone -
func (hfb *hdrForBlock) Clone() *hdrForBlock {
	return hfb
}

// SetNumMissingHdrs -
func (hfb *hdrForBlock) SetNumMissingHdrs(num uint32) {
	hfb.mutHdrsForBlock.Lock()
	hfb.missingHdrs = num
	hfb.mutHdrsForBlock.Unlock()
}

// SetNumMissingFinalityAttestingHdrs -
func (hfb *hdrForBlock) SetNumMissingFinalityAttestingHdrs(num uint32) {
	hfb.mutHdrsForBlock.Lock()
	hfb.missingFinalityAttestingHdrs = num
	hfb.mutHdrsForBlock.Unlock()
}

// SetHighestHdrNonce -
func (hfb *hdrForBlock) SetHighestHdrNonce(shardId uint32, nonce uint64) {
	hfb.mutHdrsForBlock.Lock()
	hfb.highestHdrNonce[shardId] = nonce
	hfb.mutHdrsForBlock.Unlock()
}

// HdrInfo -
type HdrInfo struct {
	UsedInBlock bool
	Hdr         data.HeaderHandler
}

// SetHdrHashAndInfo -
func (hfb *hdrForBlock) SetHdrHashAndInfo(hash string, info *HdrInfo) {
	hfb.mutHdrsForBlock.Lock()
	hfb.hdrHashAndInfo[hash] = &hdrInfo{
		hdr:         info.Hdr,
		usedInBlock: info.UsedInBlock,
	}
	hfb.mutHdrsForBlock.Unlock()
}

// GetHdrHashMap -
func (hfb *hdrForBlock) GetHdrHashMap() map[string]data.HeaderHandler {
	m := make(map[string]data.HeaderHandler)

	hfb.mutHdrsForBlock.RLock()
	for hash, hi := range hfb.hdrHashAndInfo {
		m[hash] = hi.hdr
	}
	hfb.mutHdrsForBlock.RUnlock()

	return m
}

// GetHighestHdrNonce -
func (hfb *hdrForBlock) GetHighestHdrNonce() map[uint32]uint64 {
	m := make(map[uint32]uint64)

	hfb.mutHdrsForBlock.RLock()
	for shardId, nonce := range hfb.highestHdrNonce {
		m[shardId] = nonce
	}
	hfb.mutHdrsForBlock.RUnlock()

	return m
}

// GetMissingHdrs -
func (hfb *hdrForBlock) GetMissingHdrs() uint32 {
	hfb.mutHdrsForBlock.RLock()
	defer hfb.mutHdrsForBlock.RUnlock()

	return hfb.missingHdrs
}

// GetMissingFinalityAttestingHdrs -
func (hfb *hdrForBlock) GetMissingFinalityAttestingHdrs() uint32 {
	hfb.mutHdrsForBlock.RLock()
	defer hfb.mutHdrsForBlock.RUnlock()

	return hfb.missingFinalityAttestingHdrs
}

// GetHdrHashAndInfo -
func (hfb *hdrForBlock) GetHdrHashAndInfo() map[string]*HdrInfo {
	hfb.mutHdrsForBlock.RLock()
	defer hfb.mutHdrsForBlock.RUnlock()

	m := make(map[string]*HdrInfo)
	for hash, hi := range hfb.hdrHashAndInfo {
		m[hash] = &HdrInfo{
			UsedInBlock: hi.usedInBlock,
			Hdr:         hi.hdr,
		}
	}

	return m
}

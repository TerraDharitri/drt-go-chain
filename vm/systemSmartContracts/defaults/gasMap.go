package defaults

import (
	"github.com/TerraDharitri/drt-go-chain/common"
)

// FillGasMapInternal -
func FillGasMapInternal(gasMap map[string]map[string]uint64, value uint64) map[string]map[string]uint64 {
	gasMap[common.BaseOperationCost] = FillGasMapBaseOperationCosts(value)
	gasMap[common.BuiltInCost] = FillGasMapBuiltInCosts(value)
	gasMap[common.MetaChainSystemSCsCost] = FillGasMapMetaChainSystemSCsCosts(value)

	gasMap[common.MaxPerTransaction] = FillMaxPerTransaction()

	return gasMap
}

// FillGasMapBaseOperationCosts -
func FillGasMapBaseOperationCosts(value uint64) map[string]uint64 {
	gasMap := make(map[string]uint64)
	gasMap["StorePerByte"] = value
	gasMap["DataCopyPerByte"] = value
	gasMap["ReleasePerByte"] = value
	gasMap["PersistPerByte"] = value
	gasMap["CompilePerByte"] = value
	gasMap["AoTPreparePerByte"] = value
	gasMap["GetCode"] = value
	return gasMap
}

// FillGasMapBuiltInCosts -
func FillGasMapBuiltInCosts(value uint64) map[string]uint64 {
	gasMap := make(map[string]uint64)
	gasMap["ClaimDeveloperRewards"] = value
	gasMap["ChangeOwnerAddress"] = value
	gasMap["SaveUserName"] = value
	gasMap["SaveKeyValue"] = value
	gasMap["DCDTTransfer"] = value
	gasMap["DCDTBurn"] = value
	gasMap["DCDTLocalMint"] = value
	gasMap["DCDTLocalBurn"] = value
	gasMap["DCDTNFTCreate"] = value
	gasMap["DCDTNFTAddQuantity"] = value
	gasMap["DCDTNFTBurn"] = value
	gasMap["DCDTNFTTransfer"] = value
	gasMap["DCDTNFTChangeCreateOwner"] = value
	gasMap["DCDTNFTAddUri"] = value
	gasMap["DCDTNFTUpdateAttributes"] = value
	gasMap["DCDTNFTMultiTransfer"] = value
	gasMap["DCDTModifyRoyalties"] = value
	gasMap["DCDTModifyCreator"] = value
	gasMap["DCDTNFTRecreate"] = value
	gasMap["DCDTNFTUpdate"] = value
	gasMap["DCDTNFTSetNewURIs"] = value
	gasMap["SetGuardian"] = value
	gasMap["GuardAccount"] = value
	gasMap["UnGuardAccount"] = value
	gasMap["TrieLoadPerNode"] = value
	gasMap["TrieStorePerNode"] = value

	return gasMap
}

// FillGasMapMetaChainSystemSCsCosts -
func FillGasMapMetaChainSystemSCsCosts(value uint64) map[string]uint64 {
	gasMap := make(map[string]uint64)
	gasMap["Stake"] = value
	gasMap["UnStake"] = value
	gasMap["UnBond"] = value
	gasMap["Claim"] = value
	gasMap["Get"] = value
	gasMap["ChangeRewardAddress"] = value
	gasMap["ChangeValidatorKeys"] = value
	gasMap["UnJail"] = value
	gasMap["DCDTIssue"] = value
	gasMap["DCDTOperations"] = value
	gasMap["Proposal"] = value
	gasMap["Vote"] = value
	gasMap["DelegateVote"] = value
	gasMap["RevokeVote"] = value
	gasMap["CloseProposal"] = value
	gasMap["DelegationOps"] = value
	gasMap["UnStakeTokens"] = value
	gasMap["UnBondTokens"] = value
	gasMap["DelegationMgrOps"] = value
	gasMap["GetAllNodeStates"] = value
	gasMap["ValidatorToDelegation"] = value
	gasMap["GetActiveFund"] = value
	gasMap["FixWaitingListSize"] = value

	return gasMap
}

// FillMaxPerTransaction -
func FillMaxPerTransaction() map[string]uint64 {
	gasMap := make(map[string]uint64)
	gasMap["MaxBuiltInCallsPerTx"] = 100
	gasMap["MaxNumberOfTransfersPerTx"] = 100
	gasMap["MaxNumberOfTrieReadsPerTx"] = 100000
	return gasMap
}

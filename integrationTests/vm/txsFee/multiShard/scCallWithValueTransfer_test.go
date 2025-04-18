package multiShard

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm"
	"github.com/TerraDharitri/drt-go-chain/integrationTests/vm/txsFee/utils"
	"github.com/stretchr/testify/require"
)

func TestDeployContractAndTransferValueSCProcessorV1(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testDeployContractAndTransferValue(t, 1000)
}

func TestDeployContractAndTransferValueSCProcessorV2(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testDeployContractAndTransferValue(t, 0)
}

func testDeployContractAndTransferValue(t *testing.T, scProcessorV2EnabledEpoch uint32) {
	testContextSource, err := vm.CreatePreparedTxProcessorWithVMsMultiShard(0, config.EnableEpochs{}, 1)
	require.Nil(t, err)
	defer testContextSource.Close()

	configEnabledEpochs := config.EnableEpochs{}
	configEnabledEpochs.SCProcessorV2EnableEpoch = scProcessorV2EnabledEpoch

	testContextDst, err := vm.CreatePreparedTxProcessorWithVMsMultiShard(1, configEnabledEpochs, 1)
	require.Nil(t, err)
	defer testContextDst.Close()

	pathToContract := "../testdata/transferValueContract/transferValue.wasm"

	// deploy contract on the source context
	scAddrShardSource, ownerAddress := utils.DoDeployWithCustomParams(t, testContextSource, pathToContract, big.NewInt(100000000000), 20000, nil)
	utils.CleanAccumulatedIntermediateTransactions(t, testContextSource)

	// deploy the same contract on the destination context (in this test we need a payable contract)
	scAddressShardDestination, _ := utils.DoDeployWithCustomParams(t, testContextDst, pathToContract, big.NewInt(100000000000), 20000, nil)
	utils.CleanAccumulatedIntermediateTransactions(t, testContextDst)
	testContextDst.TxsLogsProcessor.Clean()

	receiverHex := hex.EncodeToString(scAddressShardDestination)
	tx := vm.CreateTransaction(1, big.NewInt(1000000), ownerAddress, scAddrShardSource, 10, 20000, []byte("send_rewa@"+receiverHex))
	retCode, err := testContextSource.TxProcessor.ProcessTransaction(tx)
	require.Equal(t, vmcommon.Ok, retCode)
	require.Nil(t, err)

	intermediateTxs := testContextSource.GetIntermediateTransactions(t)
	require.NotNil(t, intermediateTxs)

	utils.ProcessSCRResult(t, testContextDst, intermediateTxs[0], vmcommon.Ok, nil)
	logs := testContextDst.TxsLogsProcessor.GetAllCurrentLogs()
	require.Len(t, logs, 1)

	generatedLogIdentifier := string(logs[0].LogHandler.GetLogEvents()[0].GetIdentifier())
	require.Equal(t, generatedLogIdentifier, core.CompletedTxEventIdentifier)
}

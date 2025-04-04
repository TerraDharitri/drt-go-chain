package interceptedRequestTxBlockBody

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-chain/dataRetriever/requestHandlers"
	"github.com/TerraDharitri/drt-go-chain/integrationTests"
	"github.com/TerraDharitri/drt-go-chain/process/factory"
	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNode_GenerateSendInterceptTxBlockBodyWithNetMessenger(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	hasher := integrationTests.TestHasher
	marshalizer := integrationTests.TestMarshalizer
	var numOfShards uint32 = 1
	var shardID uint32 = 0
	var txSignPrivKeyShardId uint32 = 0

	fmt.Println("Requester:	")
	nRequester := integrationTests.NewTestProcessorNode(integrationTests.ArgTestProcessorNode{
		MaxShards:            numOfShards,
		NodeShardId:          shardID,
		TxSignPrivKeyShardId: txSignPrivKeyShardId,
	})

	fmt.Println("Resolver:")
	nResolver := integrationTests.NewTestProcessorNode(integrationTests.ArgTestProcessorNode{
		MaxShards:            numOfShards,
		NodeShardId:          shardID,
		TxSignPrivKeyShardId: txSignPrivKeyShardId,
	})

	defer func() {
		nRequester.Close()
		nResolver.Close()
	}()

	//connect messengers together
	time.Sleep(time.Second)
	err := nRequester.ConnectOnMain(nResolver)
	require.Nil(t, err)

	time.Sleep(time.Second)

	//Step 1. Generate a block body
	body := block.Body{
		MiniBlocks: []*block.MiniBlock{
			{
				ReceiverShardID: 0,
				SenderShardID:   0,
				TxHashes: [][]byte{
					hasher.Compute("tx1"),
				},
			},
		},
	}

	miniBlock := body.MiniBlocks[0]
	miniBlockHashes := make([][]byte, 1)

	txBlockBodyBuff, _ := marshalizer.Marshal(miniBlock)
	txBlockBodyHash := hasher.Compute(string(txBlockBodyBuff))

	//Step 2. resolver has the tx block body
	nResolver.DataPool.MiniBlocks().HasOrAdd(txBlockBodyHash, miniBlock, miniBlock.Size())
	fmt.Printf("Added %s to dPoolResolver\n", base64.StdEncoding.EncodeToString(txBlockBodyHash))

	//Step 3. wire up a received handler
	chanDone := make(chan bool)

	nRequester.DataPool.MiniBlocks().RegisterHandler(func(key []byte, value interface{}) {
		txBlockBodyStored, _ := nRequester.DataPool.MiniBlocks().Get(key)

		if reflect.DeepEqual(txBlockBodyStored, miniBlock) {
			chanDone <- true
		}

		assert.Equal(t, txBlockBodyStored, miniBlock)
	}, core.UniqueIdentifier())

	//Step 4. request tx block body
	txBlockBodyRequester, _ := nRequester.RequestersFinder.IntraShardRequester(factory.MiniBlocksTopic)
	miniBlockRequester := txBlockBodyRequester.(requestHandlers.HashSliceRequester)
	miniBlockHashes[0] = txBlockBodyHash
	nRequester.WhiteListHandler.Add(miniBlockHashes)
	_ = miniBlockRequester.RequestDataFromHashArray(miniBlockHashes, 0)

	select {
	case <-chanDone:
	case <-time.After(time.Second * 10):
		assert.Fail(t, "timeout")
	}
}

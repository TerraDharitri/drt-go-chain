package rewards

import (
	"bytes"
	"testing"

	"github.com/TerraDharitri/drt-go-chain/integrationTests/resolvers"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/process/factory"
	"github.com/TerraDharitri/drt-go-chain-core/core"
)

func TestRequestResolveRewardsByHashRequestingShardResolvingOtherShard(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	rm := resolvers.NewReceiverMonitor(t)
	shardIdResolver := core.MetachainShardId
	shardIdRequester := uint32(0)
	nResolver, nRequester := resolvers.CreateResolverRequester(shardIdResolver, shardIdRequester)
	defer func() {
		nRequester.Close()
		nResolver.Close()
	}()
	headerNonce := uint64(0)
	reward, hash := resolvers.CreateReward(headerNonce)

	//add reward with round 0 in pool
	cacheId := process.ShardCacherIdentifier(shardIdRequester, core.MetachainShardId)
	nResolver.DataPool.RewardTransactions().AddData(hash, reward, reward.Size(), cacheId)

	//setup header received event
	nRequester.DataPool.RewardTransactions().RegisterOnAdded(
		func(key []byte, value interface{}) {
			if bytes.Equal(key, hash) {
				resolvers.Log.Info("received reward tx", "hash", key)
				rm.Done()
			}
		},
	)

	//request by hash should work
	requester, err := nRequester.RequestersFinder.CrossShardRequester(factory.RewardsTransactionTopic, core.MetachainShardId)
	resolvers.Log.LogIfError(err)
	err = requester.RequestDataFromHash(hash, 0)
	resolvers.Log.LogIfError(err)

	rm.WaitWithTimeout()
}

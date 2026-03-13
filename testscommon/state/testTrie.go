package state

import (
	"fmt"

	"github.com/TerraDharitri/drt-go-chain-core/hashing"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	"github.com/TerraDharitri/drt-go-chain/common"
	disabled2 "github.com/TerraDharitri/drt-go-chain/common/disabled"
	"github.com/TerraDharitri/drt-go-chain/common/statistics/disabled"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/enableEpochsHandlerMock"
	"github.com/TerraDharitri/drt-go-chain/testscommon/hashingMocks"
	"github.com/TerraDharitri/drt-go-chain/testscommon/marshallerMock"
	"github.com/TerraDharitri/drt-go-chain/trie"
)

// GetDefaultTrieParameters -
func GetDefaultTrieParameters() (common.StorageManager, marshal.Marshalizer, hashing.Hasher) {
	db := testscommon.NewMemDbMock()
	hasher := &hashingMocks.HasherMock{}
	marshaller := &marshallerMock.MarshalizerMock{}

	tsmArgs := trie.NewTrieStorageManagerArgs{
		MainStorer:  db,
		Marshalizer: marshaller,
		Hasher:      hasher,
		GeneralConfig: config.TrieStorageManagerConfig{
			SnapshotsGoroutineNum: 5,
		},
		IdleProvider:   disabled2.NewProcessStatusHandler(),
		Identifier:     "identifier",
		StatsCollector: disabled.NewStateStatistics(),
	}
	tsm, _ := trie.NewTrieStorageManager(tsmArgs)
	return tsm, marshaller, hasher
}

// GetNewTrie -
func GetNewTrie() common.Trie {
	tsm, marshaller, hasher := GetDefaultTrieParameters()
	tr, _ := trie.NewTrie(tsm, marshaller, hasher, &enableEpochsHandlerMock.EnableEpochsHandlerStub{}, 5)

	return tr
}

// AddDataToTrie -
func AddDataToTrie(tr common.Trie, numLeaves int) {
	for i := 0; i < numLeaves; i++ {
		val := fmt.Sprintf("value%v", i)
		_ = tr.Update([]byte(val), []byte(val))
	}
	_ = tr.Commit()
}

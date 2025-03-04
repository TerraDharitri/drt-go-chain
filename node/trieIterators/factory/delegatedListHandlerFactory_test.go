package factory

import (
	"fmt"
	"sync"
	"testing"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain/node/mock"
	"github.com/TerraDharitri/drt-go-chain/node/trieIterators"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	stateMock "github.com/TerraDharitri/drt-go-chain/testscommon/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDelegatedListHandlerHandler_Disabled(t *testing.T) {
	t.Parallel()

	args := trieIterators.ArgTrieIteratorProcessor{
		ShardID: 0,
	}

	delegatedListHandler, err := CreateDelegatedListHandler(args)
	require.Nil(t, err)
	assert.Equal(t, "*disabled.delegatedListProcessor", fmt.Sprintf("%T", delegatedListHandler))
}

func TestCreateDelegatedListHandlerHandler_DelegatedListProcessor(t *testing.T) {
	t.Parallel()

	args := trieIterators.ArgTrieIteratorProcessor{
		ShardID: core.MetachainShardId,
		Accounts: &trieIterators.AccountsWrapper{
			Mutex:           &sync.Mutex{},
			AccountsAdapter: &stateMock.AccountsStub{},
		},
		PublicKeyConverter: &testscommon.PubkeyConverterMock{},
		QueryService:       &mock.SCQueryServiceStub{},
	}

	delegatedListHandler, err := CreateDelegatedListHandler(args)
	require.Nil(t, err)
	assert.Equal(t, "*trieIterators.delegatedListProcessor", fmt.Sprintf("%T", delegatedListHandler))
}

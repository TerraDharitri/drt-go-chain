package trieIterators

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core/keyValStorage"
	"github.com/TerraDharitri/drt-go-chain-core/data/api"
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/node/mock"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/state"
	"github.com/TerraDharitri/drt-go-chain/state/accounts"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	stateMock "github.com/TerraDharitri/drt-go-chain/testscommon/state"
	trieMock "github.com/TerraDharitri/drt-go-chain/testscommon/trie"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDirectStakedListProcessor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		argsFunc func() ArgTrieIteratorProcessor
		exError  error
	}{
		{
			name: "NilAccounts",
			argsFunc: func() ArgTrieIteratorProcessor {
				arg := createMockArgs()
				arg.Accounts = nil

				return arg
			},
			exError: ErrNilAccountsAdapter,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dslp, err := NewDirectStakedListProcessor(tt.argsFunc())
			require.True(t, errors.Is(err, tt.exError))
			require.Nil(t, dslp)
		})
	}

	dslp, err := NewDirectStakedListProcessor(createMockArgs())
	require.NotNil(t, dslp)
	require.Nil(t, err)
}

func TestDirectStakedListProc_GetDelegatorsListContextShouldTimeout(t *testing.T) {
	t.Parallel()

	validators := [][]byte{[]byte("validator1"), []byte("validator2")}

	arg := createMockArgs()
	arg.PublicKeyConverter = testscommon.NewPubkeyConverterMock(10)
	arg.QueryService = &mock.SCQueryServiceStub{
		ExecuteQueryCalled: func(query *process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error) {
			return nil, nil, fmt.Errorf("not an expected call")
		},
	}
	arg.Accounts.AccountsAdapter = &stateMock.AccountsStub{
		GetExistingAccountCalled: func(addressContainer []byte) (vmcommon.AccountHandler, error) {
			return createScAccount(addressContainer, validators, addressContainer, time.Second), nil
		},
		RecreateTrieCalled: func(rootHash common.RootHashHolder) error {
			return nil
		},
	}
	dslp, _ := NewDirectStakedListProcessor(arg)

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	directStakedList, err := dslp.GetDirectStakedList(ctxWithTimeout)
	require.Equal(t, ErrTrieOperationsTimeout, err)
	require.Nil(t, directStakedList)
}

func TestDirectStakedListProc_GetDelegatorsListShouldWork(t *testing.T) {
	t.Parallel()

	validators := [][]byte{[]byte("validator1"), []byte("validator2")}

	arg := createMockArgs()
	arg.PublicKeyConverter = testscommon.NewPubkeyConverterMock(10)
	arg.QueryService = &mock.SCQueryServiceStub{
		ExecuteQueryCalled: func(query *process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error) {
			switch query.FuncName {
			case "getTotalStakedTopUpStakedBlsKeys":
				for index, validator := range validators {
					if bytes.Equal(validator, query.Arguments[0]) {
						topUpValue := big.NewInt(int64(index + 1))
						totalStakedValue := big.NewInt(int64(index+1) * 10)

						return &vmcommon.VMOutput{
							ReturnData: [][]byte{topUpValue.Bytes(), totalStakedValue.Bytes(), make([]byte, 0)},
						}, nil, nil
					}
				}
			}

			return nil, nil, fmt.Errorf("not an expected call")
		},
	}
	arg.Accounts.AccountsAdapter = &stateMock.AccountsStub{
		GetExistingAccountCalled: func(addressContainer []byte) (vmcommon.AccountHandler, error) {
			return createScAccount(addressContainer, validators, addressContainer, 0), nil
		},
		RecreateTrieCalled: func(rootHash common.RootHashHolder) error {
			return nil
		},
	}
	dslp, _ := NewDirectStakedListProcessor(arg)

	directStakedList, err := dslp.GetDirectStakedList(context.Background())
	require.Nil(t, err)
	require.Equal(t, 2, len(directStakedList))
	encodedValidator0PubKey, err := arg.PublicKeyConverter.Encode(validators[0])
	require.Nil(t, err)
	encodedValidator1PubKey, err := arg.PublicKeyConverter.Encode(validators[1])
	require.Nil(t, err)

	expectedDirectStake1 := api.DirectStakedValue{
		Address:    encodedValidator0PubKey,
		BaseStaked: "9",
		TopUp:      "1",
		Total:      "10",
	}
	expectedDirectStake2 := api.DirectStakedValue{
		Address:    encodedValidator1PubKey,
		BaseStaked: "18",
		TopUp:      "2",
		Total:      "20",
	}

	assert.Equal(t, []*api.DirectStakedValue{&expectedDirectStake1, &expectedDirectStake2}, directStakedList)
}

func createScAccount(address []byte, leaves [][]byte, rootHash []byte, timeSleep time.Duration) state.UserAccountHandler {
	dtt := &trieMock.DataTrieTrackerStub{
		DataTrieCalled: func() common.Trie {
			return &trieMock.TrieStub{
				RootCalled: func() ([]byte, error) {
					return rootHash, nil
				},
				GetAllLeavesOnChannelCalled: func(leavesChannels *common.TrieIteratorChannels, ctx context.Context, rootHash []byte, _ common.KeyBuilder, _ common.TrieLeafParser) error {
					go func() {
						time.Sleep(timeSleep)
						for _, leafBuff := range leaves {
							leaf := keyValStorage.NewKeyValStorage(leafBuff, nil)
							leavesChannels.LeavesChan <- leaf
						}

						close(leavesChannels.LeavesChan)
						leavesChannels.ErrChan.Close()
					}()

					return nil
				},
			}
		},
	}

	acc, _ := accounts.NewUserAccount(address, dtt, &trieMock.TrieLeafParserStub{})

	return acc
}

func TestDirectStakedListProcessor_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var dslp *directStakedListProcessor
	require.True(t, dslp.IsInterfaceNil())

	dslp, _ = NewDirectStakedListProcessor(createMockArgs())
	require.False(t, dslp.IsInterfaceNil())
}

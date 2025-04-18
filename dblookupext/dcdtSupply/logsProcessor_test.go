package dcdtSupply

import (
	"math/big"
	"strings"
	"testing"

	"github.com/TerraDharitri/drt-go-chain/storage"
	"github.com/TerraDharitri/drt-go-chain/testscommon/marshallerMock"
	storageStubs "github.com/TerraDharitri/drt-go-chain/testscommon/storage"
	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessLogsSaveSupplyNothingInStorage(t *testing.T) {
	t.Parallel()

	token := []byte("nft-0001")
	logs := map[string]*data.LogData{
		"txLog": {
			LogHandler: &transaction.Log{
				Events: []*transaction.Event{
					{
						Identifier: []byte("something"),
					},
					{
						Identifier: []byte(core.BuiltInFunctionDCDTNFTCreate),
						Topics: [][]byte{
							token, big.NewInt(1).Bytes(), big.NewInt(10).Bytes(),
						},
					},
					{
						Identifier: []byte(core.BuiltInFunctionDCDTNFTAddQuantity),
						Topics: [][]byte{
							token, big.NewInt(1).Bytes(), big.NewInt(50).Bytes(),
						},
					},
					{
						Identifier: []byte(core.BuiltInFunctionDCDTNFTBurn),
						Topics: [][]byte{
							token, big.NewInt(1).Bytes(), big.NewInt(30).Bytes(),
						},
					},
					{
						Identifier: []byte(core.BuiltInFunctionDCDTNFTCreate),
						Topics: [][]byte{
							token, big.NewInt(2).Bytes(), big.NewInt(10).Bytes(),
						},
					},
					{
						Identifier: []byte(core.BuiltInFunctionDCDTNFTAddQuantity),
						Topics: [][]byte{
							token, big.NewInt(2).Bytes(), big.NewInt(50).Bytes(),
						},
					},
					{
						Identifier: []byte(core.BuiltInFunctionDCDTNFTBurn),
						Topics: [][]byte{
							token, big.NewInt(2).Bytes(), big.NewInt(30).Bytes(),
						},
					},
				},
			},
		},
		"log": nil,
	}

	putCalledNum := 0
	marshalizer := marshallerMock.MarshalizerMock{}
	storer := &storageStubs.StorerStub{
		GetCalled: func(key []byte) ([]byte, error) {
			return nil, storage.ErrKeyNotFound
		},
		PutCalled: func(key, data []byte) error {
			if string(key) == processedBlockKey {
				return nil
			}

			isCollectionSupply := strings.Count(string(key), "-") == 1

			supplyValue := int64(30)
			if isCollectionSupply {
				supplyValue *= 2
			}
			var supplyDCDT SupplyDCDT
			_ = marshalizer.Unmarshal(&supplyDCDT, data)
			require.Equal(t, big.NewInt(supplyValue), supplyDCDT.Supply)

			putCalledNum++
			return nil
		},
	}

	logsProc := newLogsProcessor(marshalizer, storer)

	err := logsProc.processLogs(1, logs, false)
	require.Nil(t, err)
	require.Equal(t, 3, putCalledNum)
}

func TestTestProcessLogsSaveSupplyExistsInStorage(t *testing.T) {
	t.Parallel()

	token := []byte("dcdt-miiu")

	logs := map[string]*data.LogData{
		"txLog": {
			LogHandler: &transaction.Log{
				Events: []*transaction.Event{
					{
						Identifier: []byte(core.BuiltInFunctionDCDTLocalBurn),
						Topics: [][]byte{
							token, big.NewInt(0).Bytes(), big.NewInt(20).Bytes(),
						},
					},
					{
						Identifier: []byte(core.BuiltInFunctionDCDTLocalMint),
						Topics: [][]byte{
							token, big.NewInt(0).Bytes(), big.NewInt(25).Bytes(),
						},
					},
					nil,
				},
			},
		},
	}

	marshalizer := marshallerMock.MarshalizerMock{}
	storer := &storageStubs.StorerStub{
		GetCalled: func(key []byte) ([]byte, error) {
			supplyDCDT := &SupplyDCDT{
				Supply: big.NewInt(1000),
			}
			return marshalizer.Marshal(supplyDCDT)
		},
		PutCalled: func(key, data []byte) error {
			if string(key) == processedBlockKey {
				return nil
			}

			supplyKey := string(token)
			require.Equal(t, supplyKey, string(key))

			var supplyDCDT SupplyDCDT
			_ = marshalizer.Unmarshal(&supplyDCDT, data)
			require.Equal(t, big.NewInt(1005), supplyDCDT.Supply)

			return nil
		},
	}

	logsProc := newLogsProcessor(marshalizer, storer)

	err := logsProc.processLogs(1, logs, false)
	require.Nil(t, err)
}

func TestMakePropertiesNotNil(t *testing.T) {
	t.Parallel()

	t.Run("supply is nil", func(t *testing.T) {
		t.Parallel()

		provided := SupplyDCDT{
			Supply: nil,
			Burned: big.NewInt(1),
			Minted: big.NewInt(2),
		}
		expected := SupplyDCDT{
			Supply: big.NewInt(0),
			Burned: big.NewInt(1),
			Minted: big.NewInt(2),
		}
		makePropertiesNotNil(&provided)
		assert.Equal(t, expected, provided)
	})
	t.Run("burned is nil", func(t *testing.T) {
		t.Parallel()

		provided := SupplyDCDT{
			Supply: big.NewInt(1),
			Burned: nil,
			Minted: big.NewInt(2),
		}
		expected := SupplyDCDT{
			Supply: big.NewInt(1),
			Burned: big.NewInt(0),
			Minted: big.NewInt(2),
		}
		makePropertiesNotNil(&provided)
		assert.Equal(t, expected, provided)
	})
	t.Run("minted is nil", func(t *testing.T) {
		t.Parallel()

		provided := SupplyDCDT{
			Supply: big.NewInt(1),
			Burned: big.NewInt(2),
			Minted: nil,
		}
		expected := SupplyDCDT{
			Supply: big.NewInt(1),
			Burned: big.NewInt(2),
			Minted: big.NewInt(0),
		}
		makePropertiesNotNil(&provided)
		assert.Equal(t, expected, provided)
	})
	t.Run("all are nil", func(t *testing.T) {
		t.Parallel()

		provided := SupplyDCDT{}
		expected := SupplyDCDT{
			Supply: big.NewInt(0),
			Burned: big.NewInt(0),
			Minted: big.NewInt(0),
		}
		makePropertiesNotNil(&provided)
		assert.Equal(t, expected, provided)
	})
	t.Run("none is nil", func(t *testing.T) {
		t.Parallel()

		provided := SupplyDCDT{
			Supply: big.NewInt(1),
			Burned: big.NewInt(2),
			Minted: big.NewInt(3),
		}
		expected := SupplyDCDT{
			Supply: big.NewInt(1),
			Burned: big.NewInt(2),
			Minted: big.NewInt(3),
		}
		makePropertiesNotNil(&provided)
		assert.Equal(t, expected, provided)
	})

}

package chainSimulator

import (
	"math/big"

	"github.com/TerraDharitri/drt-go-chain/node/chainSimulator/configs"

	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
)

const (
	minGasPrice     = 1000000000
	txVersion       = 1
	mockTxSignature = "sig"

	// OkReturnCode the const for the ok return code
	OkReturnCode = "ok"
)

var (
	// ZeroValue the variable for the zero big int
	ZeroValue = big.NewInt(0)
	// OneREWA the variable for one rewa value
	OneREWA = big.NewInt(1000000000000000000)
	// MinimumStakeValue the variable for the minimum stake value
	MinimumStakeValue = big.NewInt(0).Mul(OneREWA, big.NewInt(2500))
	// InitialAmount the variable for initial minting amount in account
	InitialAmount = big.NewInt(0).Mul(OneREWA, big.NewInt(100))
)

// GenerateTransaction will generate a transaction based on input data
func GenerateTransaction(sender []byte, nonce uint64, receiver []byte, value *big.Int, data string, gasLimit uint64) *transaction.Transaction {
	return &transaction.Transaction{
		Nonce:     nonce,
		Value:     value,
		SndAddr:   sender,
		RcvAddr:   receiver,
		Data:      []byte(data),
		GasLimit:  gasLimit,
		GasPrice:  minGasPrice,
		ChainID:   []byte(configs.ChainID),
		Version:   txVersion,
		Signature: []byte(mockTxSignature),
	}
}

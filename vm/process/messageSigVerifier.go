package process

import (
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	crypto "github.com/TerraDharitri/drt-go-chain-crypto"
	"github.com/TerraDharitri/drt-go-chain/vm"
)

type messageSigVerifier struct {
	kg           crypto.KeyGenerator
	singleSigner crypto.SingleSigner
}

// NewMessageSigVerifier creates a message signature verifier
func NewMessageSigVerifier(
	kg crypto.KeyGenerator,
	singleSigner crypto.SingleSigner,
) (*messageSigVerifier, error) {

	if check.IfNil(kg) {
		return nil, vm.ErrNilKeyGenerator
	}
	if check.IfNil(singleSigner) {
		return nil, vm.ErrNilSingleSigner
	}

	return &messageSigVerifier{
		kg:           kg,
		singleSigner: singleSigner,
	}, nil
}

// Verify checks if message was signed by public key given as byte array
func (m *messageSigVerifier) Verify(message []byte, signedMessage []byte, pubKey []byte) error {
	if len(pubKey) == 0 {
		return vm.ErrNilPublicKey
	}

	actPubKey, err := m.kg.PublicKeyFromByteArray(pubKey)
	if err != nil {
		return err
	}

	return m.singleSigner.Verify(actPubKey, message, signedMessage)
}

// IsInterfaceNil returns if underlying object is nil
func (m *messageSigVerifier) IsInterfaceNil() bool {
	return m == nil
}

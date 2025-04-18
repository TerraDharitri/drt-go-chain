package cryptoMocks

import crypto "github.com/TerraDharitri/drt-go-chain-crypto"

// MultiSignerContainerStub -
type MultiSignerContainerStub struct {
	GetMultiSignerCalled func(epoch uint32) (crypto.MultiSigner, error)
}

// GetMultiSigner -
func (stub *MultiSignerContainerStub) GetMultiSigner(epoch uint32) (crypto.MultiSigner, error) {
	if stub.GetMultiSignerCalled != nil {
		return stub.GetMultiSignerCalled(epoch)
	}

	return nil, nil
}

// IsInterfaceNil -
func (stub *MultiSignerContainerStub) IsInterfaceNil() bool {
	return stub == nil
}

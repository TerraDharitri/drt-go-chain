package crypto

import crypto "github.com/TerraDharitri/drt-go-chain-crypto"

// MultiSignerContainer defines the container for different versioned multiSigner instances
type MultiSignerContainer interface {
	GetMultiSigner(epoch uint32) (crypto.MultiSigner, error)
	IsInterfaceNil() bool
}

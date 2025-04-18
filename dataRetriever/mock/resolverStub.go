package mock

import (
	"github.com/TerraDharitri/drt-go-chain/dataRetriever"
	"github.com/TerraDharitri/drt-go-chain/p2p"
	"github.com/TerraDharitri/drt-go-chain-core/core"
)

// ResolverStub -
type ResolverStub struct {
	ProcessReceivedMessageCalled func(message p2p.MessageP2P) error
	SetDebugHandlerCalled        func(handler dataRetriever.DebugHandler) error
	CloseCalled                  func() error
}

// ProcessReceivedMessage -
func (rs *ResolverStub) ProcessReceivedMessage(message p2p.MessageP2P, _ core.PeerID, _ p2p.MessageHandler) error {
	return rs.ProcessReceivedMessageCalled(message)
}

// SetDebugHandler -
func (rs *ResolverStub) SetDebugHandler(handler dataRetriever.DebugHandler) error {
	if rs.SetDebugHandlerCalled != nil {
		return rs.SetDebugHandlerCalled(handler)
	}

	return nil
}

// Close -
func (rs *ResolverStub) Close() error {
	if rs.CloseCalled != nil {
		return rs.CloseCalled()
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (rs *ResolverStub) IsInterfaceNil() bool {
	return rs == nil
}

package topicsender

import (
	"fmt"
	"sync"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	logger "github.com/TerraDharitri/drt-go-chain-logger"
	"github.com/TerraDharitri/drt-go-chain/dataRetriever"
	"github.com/TerraDharitri/drt-go-chain/debug/handler"
	"github.com/TerraDharitri/drt-go-chain/p2p"
	"github.com/TerraDharitri/drt-go-chain/p2p/factory"
)

var log = logger.GetOrCreate("dataretriever/topicsender")

const (
	minPeersToQuery    = 2
	preferredPeerIndex = -1
)

// ArgBaseTopicSender is the base DTO used to create a new topic sender instance
type ArgBaseTopicSender struct {
	MainMessenger                   p2p.Messenger
	FullArchiveMessenger            p2p.Messenger
	TopicName                       string
	OutputAntiflooder               dataRetriever.P2PAntifloodHandler
	MainPreferredPeersHolder        dataRetriever.PreferredPeersHolderHandler
	FullArchivePreferredPeersHolder dataRetriever.PreferredPeersHolderHandler
	TargetShardId                   uint32
}

type baseTopicSender struct {
	mainMessenger                          p2p.Messenger
	fullArchiveMessenger                   p2p.Messenger
	topicName                              string
	outputAntiflooder                      dataRetriever.P2PAntifloodHandler
	mutDebugHandler                        sync.RWMutex
	debugHandler                           dataRetriever.DebugHandler
	mainPreferredPeersHolderHandler        dataRetriever.PreferredPeersHolderHandler
	fullArchivePreferredPeersHolderHandler dataRetriever.PreferredPeersHolderHandler
	targetShardId                          uint32
}

func createBaseTopicSender(args ArgBaseTopicSender) *baseTopicSender {
	return &baseTopicSender{
		mainMessenger:                          args.MainMessenger,
		fullArchiveMessenger:                   args.FullArchiveMessenger,
		topicName:                              args.TopicName,
		outputAntiflooder:                      args.OutputAntiflooder,
		debugHandler:                           handler.NewDisabledInterceptorDebugHandler(),
		mainPreferredPeersHolderHandler:        args.MainPreferredPeersHolder,
		fullArchivePreferredPeersHolderHandler: args.FullArchivePreferredPeersHolder,
		targetShardId:                          args.TargetShardId,
	}
}

func checkBaseTopicSenderArgs(args ArgBaseTopicSender) error {
	if check.IfNil(args.MainMessenger) {
		return fmt.Errorf("%w on main network", dataRetriever.ErrNilMessenger)
	}
	if check.IfNil(args.FullArchiveMessenger) {
		return fmt.Errorf("%w on full archive network", dataRetriever.ErrNilMessenger)
	}
	if check.IfNil(args.OutputAntiflooder) {
		return dataRetriever.ErrNilAntifloodHandler
	}
	if check.IfNil(args.MainPreferredPeersHolder) {
		return fmt.Errorf("%w on main network", dataRetriever.ErrNilPreferredPeersHolder)
	}
	if check.IfNil(args.FullArchivePreferredPeersHolder) {
		return fmt.Errorf("%w on full archive network", dataRetriever.ErrNilPreferredPeersHolder)
	}
	return nil
}

func (baseSender *baseTopicSender) sendToConnectedPeer(
	topic string,
	buff []byte,
	peer core.PeerID,
	messenger p2p.MessageHandler,
) error {
	msg := &factory.Message{
		DataField:  buff,
		PeerField:  peer,
		TopicField: topic,
	}

	isPreferredOnMain := baseSender.mainPreferredPeersHolderHandler.Contains(peer)
	isPreferredOnFullArchive := baseSender.fullArchivePreferredPeersHolderHandler.Contains(peer)
	shouldAvoidAntiFloodCheck := isPreferredOnMain || isPreferredOnFullArchive
	if shouldAvoidAntiFloodCheck {
		return messenger.SendToConnectedPeer(topic, buff, peer)
	}

	err := baseSender.outputAntiflooder.CanProcessMessage(msg, peer)
	if err != nil {
		return fmt.Errorf("%w while sending %d bytes to peer %s",
			err,
			len(buff),
			p2p.PeerIdToShortString(peer),
		)
	}

	return messenger.SendToConnectedPeer(topic, buff, peer)
}

// DebugHandler returns the debug handler used in resolvers
func (baseSender *baseTopicSender) DebugHandler() dataRetriever.DebugHandler {
	baseSender.mutDebugHandler.RLock()
	defer baseSender.mutDebugHandler.RUnlock()

	return baseSender.debugHandler
}

// SetDebugHandler sets the debug handler used in resolvers
func (baseSender *baseTopicSender) SetDebugHandler(handler dataRetriever.DebugHandler) error {
	if check.IfNil(handler) {
		return dataRetriever.ErrNilDebugHandler
	}

	baseSender.mutDebugHandler.Lock()
	baseSender.debugHandler = handler
	baseSender.mutDebugHandler.Unlock()

	return nil
}

// RequestTopic returns the topic with the request suffix used for sending requests
func (baseSender *baseTopicSender) RequestTopic() string {
	return baseSender.topicName + core.TopicRequestSuffix
}

// TargetShardID returns the target shard ID for this resolver should serve data
func (baseSender *baseTopicSender) TargetShardID() uint32 {
	return baseSender.targetShardId
}

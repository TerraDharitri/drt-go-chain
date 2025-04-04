package sender

import (
	"context"
	"fmt"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/marshal"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/heartbeat"
	"github.com/TerraDharitri/drt-go-chain/heartbeat/processor"
	"github.com/TerraDharitri/drt-go-chain/p2p"
	"github.com/TerraDharitri/drt-go-chain/p2p/factory"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/sharding"
)

const minDelayBetweenSends = time.Second

// ArgPeerShardSender represents the arguments for the peer shard sender
type ArgPeerShardSender struct {
	MainMessenger             p2p.Messenger
	FullArchiveMessenger      p2p.Messenger
	Marshaller                marshal.Marshalizer
	ShardCoordinator          sharding.Coordinator
	TimeBetweenSends          time.Duration
	TimeThresholdBetweenSends float64
	NodesCoordinator          processor.NodesCoordinator
}

type peerShardSender struct {
	mainMessenger             p2p.Messenger
	fullArchiveMessenger      p2p.Messenger
	marshaller                marshal.Marshalizer
	shardCoordinator          sharding.Coordinator
	timeBetweenSends          time.Duration
	timeThresholdBetweenSends float64
	nodesCoordinator          processor.NodesCoordinator
	cancel                    func()
}

// NewPeerShardSender creates a new instance of peerShardSender
func NewPeerShardSender(args ArgPeerShardSender) (*peerShardSender, error) {
	err := checkArgPeerShardSender(args)
	if err != nil {
		return nil, err
	}

	pss := &peerShardSender{
		mainMessenger:             args.MainMessenger,
		fullArchiveMessenger:      args.FullArchiveMessenger,
		marshaller:                args.Marshaller,
		shardCoordinator:          args.ShardCoordinator,
		timeBetweenSends:          args.TimeBetweenSends,
		timeThresholdBetweenSends: args.TimeThresholdBetweenSends,
		nodesCoordinator:          args.NodesCoordinator,
	}

	var ctx context.Context
	ctx, pss.cancel = context.WithCancel(context.Background())

	go pss.startSendingShard(ctx)

	return pss, nil
}

func checkArgPeerShardSender(args ArgPeerShardSender) error {
	if check.IfNil(args.MainMessenger) {
		return fmt.Errorf("%w for main", process.ErrNilMessenger)
	}
	if check.IfNil(args.FullArchiveMessenger) {
		return fmt.Errorf("%w for full archive", process.ErrNilMessenger)
	}
	if check.IfNil(args.Marshaller) {
		return process.ErrNilMarshalizer
	}
	if check.IfNil(args.ShardCoordinator) {
		return process.ErrNilShardCoordinator
	}
	if args.TimeBetweenSends < minDelayBetweenSends {
		return fmt.Errorf("%w for TimeBetweenSends, provided %d, min expected %d",
			heartbeat.ErrInvalidTimeDuration, args.TimeBetweenSends, minDelayBetweenSends)
	}
	if args.TimeThresholdBetweenSends < minThresholdBetweenSends || args.TimeThresholdBetweenSends > maxThresholdBetweenSends {
		return fmt.Errorf("%w for TimeThresholdBetweenSends, receieved %f, min allowed %f, max allowed %f",
			heartbeat.ErrInvalidThreshold, args.TimeThresholdBetweenSends, minThresholdBetweenSends, maxThresholdBetweenSends)
	}
	if check.IfNil(args.NodesCoordinator) {
		return heartbeat.ErrNilNodesCoordinator
	}

	return nil
}

func (pss *peerShardSender) startSendingShard(ctx context.Context) {
	duration := pss.computeRandomDuration(pss.timeBetweenSends)
	timer := time.NewTimer(duration)
	defer timer.Stop()

	log.Debug("starting peer shard sender's goroutine")
	for {
		duration = pss.computeRandomDuration(pss.timeBetweenSends)
		timer.Reset(duration)

		select {
		case <-timer.C:
			pss.broadcastShard()
		case <-ctx.Done():
			log.Debug("closing peerShardSender go routine")
			return
		}
	}
}

func (pss *peerShardSender) computeRandomDuration(baseDuration time.Duration) time.Duration {
	timeBetweenSendsInNano := baseDuration.Nanoseconds()
	maxThreshold := float64(timeBetweenSendsInNano) * pss.timeThresholdBetweenSends
	randThreshold := randomizer.Intn(int(maxThreshold))

	ret := time.Duration(timeBetweenSendsInNano + int64(randThreshold))
	return ret
}

func (pss *peerShardSender) broadcastShard() {
	if pss.isCurrentNodeValidator() {
		log.Debug("peerShardSender.sendShard current node is validator, will not send messages to connected peers")
		return
	}

	peerShard := &factory.PeerShard{
		ShardId: fmt.Sprintf("%d", pss.shardCoordinator.SelfId()),
	}

	peerShardBuff, err := pss.marshaller.Marshal(peerShard)
	if err != nil {
		return
	}

	log.Debug("broadcast peer shard", "shard", peerShard.ShardId)
	pss.mainMessenger.Broadcast(common.ConnectionTopic, peerShardBuff)
	pss.fullArchiveMessenger.Broadcast(common.ConnectionTopic, peerShardBuff)
}

func (pss *peerShardSender) isCurrentNodeValidator() bool {
	currentBLSKey := pss.nodesCoordinator.GetOwnPublicKey()
	_, _, err := pss.nodesCoordinator.GetValidatorWithPublicKey(currentBLSKey)

	return err == nil
}

// Close triggers the closing of the internal goroutine
func (pss *peerShardSender) Close() error {
	log.Debug("closing shardSender...")
	pss.cancel()

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (pss *peerShardSender) IsInterfaceNil() bool {
	return pss == nil
}

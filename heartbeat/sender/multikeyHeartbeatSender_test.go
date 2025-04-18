package sender

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	crypto "github.com/TerraDharitri/drt-go-chain-crypto"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/heartbeat"
	"github.com/TerraDharitri/drt-go-chain/heartbeat/mock"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	"github.com/TerraDharitri/drt-go-chain/testscommon/marshallerMock"
	"github.com/TerraDharitri/drt-go-chain/testscommon/p2pmocks"
	"github.com/stretchr/testify/assert"
)

func createMockMultikeyHeartbeatSenderArgs(argBase argBaseSender) argMultikeyHeartbeatSender {
	return argMultikeyHeartbeatSender{
		argBaseSender: argBase,
		peerTypeProvider: &mock.PeerTypeProviderStub{
			ComputeForPubKeyCalled: func(pubKey []byte) (common.PeerType, uint32, error) {
				return common.ObserverList, 0, nil
			},
		},
		versionNumber:     "version",
		baseVersionNumber: "base version",
		nodeDisplayName:   "default name",
		identity:          "default identity",
		peerSubType:       core.FullHistoryObserver,
		currentBlockProvider: &mock.CurrentBlockProviderStub{
			GetCurrentBlockHeaderCalled: func() data.HeaderHandler {
				return &testscommon.HeaderHandlerStub{}
			},
		},
		managedPeersHolder: &testscommon.ManagedPeersHolderStub{
			IsKeyManagedByCurrentNodeCalled: func(pkBytes []byte) bool {
				return true
			},
		},
		shardCoordinator:           createShardCoordinatorInShard(0),
		trieSyncStatisticsProvider: &testscommon.SizeSyncStatisticsHandlerStub{},
	}
}

func TestNewMultikeyHeartbeatSender(t *testing.T) {
	t.Parallel()

	t.Run("nil main messenger should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.mainMessenger = nil

		senderInstance, err := newMultikeyHeartbeatSender(args)
		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrNilMessenger))
	})
	t.Run("nil full archive messenger should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.fullArchiveMessenger = nil

		senderInstance, err := newMultikeyHeartbeatSender(args)
		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrNilMessenger))
	})
	t.Run("nil marshaller should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.marshaller = nil

		senderInstance, err := newMultikeyHeartbeatSender(args)
		assert.Nil(t, senderInstance)
		assert.Equal(t, heartbeat.ErrNilMarshaller, err)
	})
	t.Run("empty topic should error", func(t *testing.T) {
		t.Parallel()

		argsBase := createMockBaseArgs()
		argsBase.topic = ""

		args := createMockMultikeyHeartbeatSenderArgs(argsBase)
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.Equal(t, heartbeat.ErrEmptySendTopic, err)
	})
	t.Run("invalid time between sends should error", func(t *testing.T) {
		t.Parallel()

		argsBase := createMockBaseArgs()
		argsBase.timeBetweenSends = time.Second - time.Nanosecond

		args := createMockMultikeyHeartbeatSenderArgs(argsBase)
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrInvalidTimeDuration))
		assert.True(t, strings.Contains(err.Error(), "timeBetweenSends"))
		assert.False(t, strings.Contains(err.Error(), "timeBetweenSendsWhenError"))
	})
	t.Run("invalid time between sends should error", func(t *testing.T) {
		t.Parallel()

		argsBase := createMockBaseArgs()
		argsBase.timeBetweenSendsWhenError = time.Second - time.Nanosecond

		args := createMockMultikeyHeartbeatSenderArgs(argsBase)
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrInvalidTimeDuration))
		assert.True(t, strings.Contains(err.Error(), "timeBetweenSendsWhenError"))
	})
	t.Run("threshold too small should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.thresholdBetweenSends = 0.001
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrInvalidThreshold))
		assert.True(t, strings.Contains(err.Error(), "thresholdBetweenSends"))
	})
	t.Run("threshold too big should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.thresholdBetweenSends = 1.001
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrInvalidThreshold))
		assert.True(t, strings.Contains(err.Error(), "thresholdBetweenSends"))
	})
	t.Run("nil peerTypeProvider should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.peerTypeProvider = nil

		senderInstance, err := newMultikeyHeartbeatSender(args)
		assert.Nil(t, senderInstance)
		assert.Equal(t, heartbeat.ErrNilPeerTypeProvider, err)
	})
	t.Run("version number too long should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.versionNumber = strings.Repeat("a", maxSizeInBytes+1)
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrPropertyTooLong))
		assert.True(t, strings.Contains(err.Error(), "versionNumber"))
	})
	t.Run("base version number too long should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.baseVersionNumber = strings.Repeat("a", maxSizeInBytes+1)
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrPropertyTooLong))
		assert.True(t, strings.Contains(err.Error(), "baseVersionNumber"))
	})
	t.Run("node display name too long should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.nodeDisplayName = strings.Repeat("a", maxSizeInBytes+1)
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrPropertyTooLong))
		assert.True(t, strings.Contains(err.Error(), "nodeDisplayName"))
	})
	t.Run("identity too long should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.identity = strings.Repeat("a", maxSizeInBytes+1)
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.True(t, errors.Is(err, heartbeat.ErrPropertyTooLong))
		assert.True(t, strings.Contains(err.Error(), "identity"))
	})
	t.Run("nil currentBlockProvider should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.currentBlockProvider = nil

		senderInstance, err := newMultikeyHeartbeatSender(args)
		assert.Nil(t, senderInstance)
		assert.Equal(t, heartbeat.ErrNilCurrentBlockProvider, err)
	})
	t.Run("nil managed peers holder should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.managedPeersHolder = nil
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.Equal(t, heartbeat.ErrNilManagedPeersHolder, err)
	})
	t.Run("nil shard coordinator should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		args.shardCoordinator = nil
		senderInstance, err := newMultikeyHeartbeatSender(args)

		assert.Nil(t, senderInstance)
		assert.Equal(t, heartbeat.ErrNilShardCoordinator, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())

		senderInstance, err := newMultikeyHeartbeatSender(args)
		assert.NotNil(t, senderInstance)
		assert.Nil(t, err)
	})
}

func TestMultikeyHeartbeatSender_Execute(t *testing.T) {
	t.Parallel()

	t.Run("execute errors, should set the error time duration value", func(t *testing.T) {
		t.Parallel()

		wasCalled := false
		argsBase := createMockBaseArgs()
		argsBase.timeBetweenSendsWhenError = time.Second * 3
		argsBase.timeBetweenSends = time.Second * 2
		argsBase.marshaller = &marshallerMock.MarshalizerStub{
			MarshalCalled: func(obj interface{}) ([]byte, error) {
				return nil, expectedErr
			},
		}

		args := createMockMultikeyHeartbeatSenderArgs(argsBase)
		senderInstance, _ := newMultikeyHeartbeatSender(args)
		senderInstance.timerHandler = &mock.TimerHandlerStub{
			CreateNewTimerCalled: func(duration time.Duration) {
				assert.Equal(t, argsBase.timeBetweenSendsWhenError, duration)
				wasCalled = true
			},
		}

		senderInstance.Execute()
		assert.True(t, wasCalled)
	})
	t.Run("execute worked, should set the normal time duration value", func(t *testing.T) {
		t.Parallel()

		wasCalled := false
		argsBase := createMockBaseArgs()
		argsBase.timeBetweenSendsWhenError = time.Second * 3
		argsBase.timeBetweenSends = time.Second * 2

		args := createMockMultikeyHeartbeatSenderArgs(argsBase)
		senderInstance, _ := newMultikeyHeartbeatSender(args)
		senderInstance.timerHandler = &mock.TimerHandlerStub{
			CreateNewTimerCalled: func(duration time.Duration) {
				floatTBS := float64(argsBase.timeBetweenSends.Nanoseconds())
				maxDuration := floatTBS + floatTBS*argsBase.thresholdBetweenSends
				assert.True(t, time.Duration(maxDuration) > duration)
				assert.True(t, argsBase.timeBetweenSends <= duration)
				wasCalled = true
			},
		}

		senderInstance.Execute()
		assert.True(t, wasCalled)
	})
}

func TestMultikeyHeartbeatSender_execute(t *testing.T) {
	t.Parallel()

	t.Run("should send the current node heartbeat", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		mainBroadcastCalled := false
		fullArchiveBroadcastCalled := false
		recordedMessagesFromMain := make(map[core.PeerID][][]byte)
		recordedMessagesFromFullArchive := make(map[core.PeerID][][]byte)
		args.mainMessenger = &p2pmocks.MessengerStub{
			BroadcastCalled: func(topic string, buff []byte) {
				assert.Equal(t, args.topic, topic)
				recordedMessagesFromMain[args.mainMessenger.ID()] = append(recordedMessagesFromMain[args.mainMessenger.ID()], buff)
				mainBroadcastCalled = true
			},
		}
		args.fullArchiveMessenger = &p2pmocks.MessengerStub{
			BroadcastCalled: func(topic string, buff []byte) {
				assert.Equal(t, args.topic, topic)
				recordedMessagesFromFullArchive[args.mainMessenger.ID()] = append(recordedMessagesFromFullArchive[args.mainMessenger.ID()], buff)
				fullArchiveBroadcastCalled = true
			},
		}

		senderInstance, _ := newMultikeyHeartbeatSender(args)

		err := senderInstance.execute()
		assert.Nil(t, err)
		assert.True(t, mainBroadcastCalled)
		assert.True(t, fullArchiveBroadcastCalled)
		assert.Equal(t, 1, len(recordedMessagesFromMain))
		checkRecordedMessages(t, recordedMessagesFromMain, args, args.versionNumber, args.nodeDisplayName, args.mainMessenger.ID(), core.FullHistoryObserver)
		assert.Equal(t, 1, len(recordedMessagesFromFullArchive))
		checkRecordedMessages(t, recordedMessagesFromFullArchive, args, args.versionNumber, args.nodeDisplayName, args.mainMessenger.ID(), core.FullHistoryObserver)
		assert.Equal(t, uint64(1), args.currentBlockProvider.GetCurrentBlockHeader().GetNonce())
	})
	t.Run("should send the current node heartbeat and some multikey heartbeats", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		recordedMainMessages := make(map[core.PeerID][][]byte)
		args.mainMessenger = &p2pmocks.MessengerStub{
			BroadcastCalled: func(topic string, buff []byte) {
				assert.Equal(t, args.topic, topic)
				recordedMainMessages[args.mainMessenger.ID()] = append(recordedMainMessages[args.mainMessenger.ID()], buff)
			},
			BroadcastUsingPrivateKeyCalled: func(topic string, buff []byte, pid core.PeerID, skBytes []byte) {
				assert.Equal(t, args.topic, topic)
				recordedMainMessages[pid] = append(recordedMainMessages[pid], buff)
			},
		}
		args.managedPeersHolder = &testscommon.ManagedPeersHolderStub{
			IsKeyManagedByCurrentNodeCalled: func(pkBytes []byte) bool {
				return string(pkBytes) != "dd"
			},
			GetManagedKeysByCurrentNodeCalled: func() map[string]crypto.PrivateKey {
				return map[string]crypto.PrivateKey{
					"aa": &mock.PrivateKeyStub{}, // in shard 1, handled by current node
					"bb": &mock.PrivateKeyStub{}, // in shard 1, handled by current node
					"cc": &mock.PrivateKeyStub{}, // in shard 1, handled by current node
					"dd": &mock.PrivateKeyStub{}, // in shard 1, not handled by current node
					"ee": &mock.PrivateKeyStub{}, // in shard 2, handled by current node
				}
			},
			GetP2PIdentityCalled: func(pkBytes []byte) ([]byte, core.PeerID, error) {
				return []byte(string(pkBytes) + "_p2p"), core.PeerID(string(pkBytes) + "_pid"), nil
			},
			GetMachineIDCalled: func(pkBytes []byte) (string, error) {
				return string(pkBytes) + "_machineID", nil
			},
			GetNameAndIdentityCalled: func(pkBytes []byte) (string, string, error) {
				return string(pkBytes) + "_name", string(pkBytes) + "_identity", nil
			},
		}
		args.peerTypeProvider = &mock.PeerTypeProviderStub{
			ComputeForPubKeyCalled: func(pubKey []byte) (common.PeerType, uint32, error) {
				if string(pubKey) == "ee" {
					return "", 2, nil
				}
				return "", 1, nil
			},
		}
		args.shardCoordinator = createShardCoordinatorInShard(1)

		senderInstance, _ := newMultikeyHeartbeatSender(args)

		err := senderInstance.execute()
		assert.Nil(t, err)
		assert.Equal(t, 4, len(recordedMainMessages)) // current pid, aa, bb, cc

		checkRecordedMessages(t,
			recordedMainMessages,
			args,
			args.versionNumber,
			args.nodeDisplayName,
			args.mainMessenger.ID(),
			core.FullHistoryObserver)

		checkRecordedMessages(t,
			recordedMainMessages,
			args,
			args.baseVersionNumber+"/aa_machineID",
			"aa_name",
			"aa_pid",
			core.RegularPeer)

		checkRecordedMessages(t,
			recordedMainMessages,
			args,
			args.baseVersionNumber+"/bb_machineID",
			"bb_name",
			"bb_pid",
			core.RegularPeer)

		checkRecordedMessages(t,
			recordedMainMessages,
			args,
			args.baseVersionNumber+"/cc_machineID",
			"cc_name",
			"cc_pid",
			core.RegularPeer)

		assert.Equal(t, uint64(1), args.currentBlockProvider.GetCurrentBlockHeader().GetNonce())
	})
}

func TestMultikeyHeartbeatSender_generateMessageBytes(t *testing.T) {
	t.Parallel()

	t.Run("version too long should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		senderInstance, _ := newMultikeyHeartbeatSender(args)

		versionNumber := strings.Repeat("a", maxSizeInBytes+1)
		nodeDisplayName := "a"
		identity := "b"
		buff, err := senderInstance.generateMessageBytes(versionNumber, nodeDisplayName, identity, 0, []byte("public key"), 0)

		assert.True(t, errors.Is(err, heartbeat.ErrPropertyTooLong))
		assert.True(t, strings.Contains(err.Error(), "versionNumber"))
		assert.Nil(t, buff)
	})
	t.Run("node name too long should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		senderInstance, _ := newMultikeyHeartbeatSender(args)

		versionNumber := "a"
		nodeDisplayName := strings.Repeat("a", maxSizeInBytes+1)
		identity := "b"
		buff, err := senderInstance.generateMessageBytes(versionNumber, nodeDisplayName, identity, 0, []byte("public key"), 0)

		assert.True(t, errors.Is(err, heartbeat.ErrPropertyTooLong))
		assert.True(t, strings.Contains(err.Error(), "nodeDisplayName"))
		assert.Nil(t, buff)
	})
	t.Run("identity too long should error", func(t *testing.T) {
		t.Parallel()

		args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
		senderInstance, _ := newMultikeyHeartbeatSender(args)

		versionNumber := "a"
		nodeDisplayName := "b"
		identity := strings.Repeat("a", maxSizeInBytes+1)
		buff, err := senderInstance.generateMessageBytes(versionNumber, nodeDisplayName, identity, 0, []byte("public key"), 0)

		assert.True(t, errors.Is(err, heartbeat.ErrPropertyTooLong))
		assert.True(t, strings.Contains(err.Error(), "identity"))
		assert.Nil(t, buff)
	})
}

func TestMultikeyHeartbeatSender_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var senderInstance *multikeyHeartbeatSender
	assert.True(t, senderInstance.IsInterfaceNil())

	args := createMockMultikeyHeartbeatSenderArgs(createMockBaseArgs())
	senderInstance, _ = newMultikeyHeartbeatSender(args)
	assert.False(t, senderInstance.IsInterfaceNil())
}

func checkRecordedMessages(
	tb testing.TB,
	recordedMessages map[core.PeerID][][]byte,
	args argMultikeyHeartbeatSender,
	version string,
	nodeDisplayName string,
	pid core.PeerID,
	peerSubType core.P2PPeerSubType,
) {
	messages := recordedMessages[pid]
	assert.True(tb, len(messages) > 0)

	for _, message := range messages {
		checkRecordedMessage(tb, message, args, version, nodeDisplayName, peerSubType)
	}
}

func checkRecordedMessage(
	tb testing.TB,
	recordedMessage []byte,
	args argMultikeyHeartbeatSender,
	version string,
	nodeDisplayName string,
	peerSubType core.P2PPeerSubType,
) {
	msg := &heartbeat.HeartbeatV2{}
	err := args.marshaller.Unmarshal(msg, recordedMessage)
	assert.Nil(tb, err)

	assert.Equal(tb, version, msg.VersionNumber)
	assert.Equal(tb, nodeDisplayName, msg.NodeDisplayName)
	assert.Equal(tb, uint32(peerSubType), msg.PeerSubType)
}

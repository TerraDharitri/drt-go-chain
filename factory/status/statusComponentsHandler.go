package status

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/core/appStatusPolling"
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data"
	"github.com/TerraDharitri/drt-go-chain/common"
	"github.com/TerraDharitri/drt-go-chain/common/statistics"
	"github.com/TerraDharitri/drt-go-chain/common/statistics/machine"
	"github.com/TerraDharitri/drt-go-chain/debug/goroutine"
	"github.com/TerraDharitri/drt-go-chain/epochStart/notifier"
	"github.com/TerraDharitri/drt-go-chain/errors"
	"github.com/TerraDharitri/drt-go-chain/factory"
	"github.com/TerraDharitri/drt-go-chain/outport"
	"github.com/TerraDharitri/drt-go-chain/p2p"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/sharding"
)

var _ factory.ComponentHandler = (*managedStatusComponents)(nil)
var _ factory.StatusComponentsHolder = (*managedStatusComponents)(nil)
var _ factory.StatusComponentsHandler = (*managedStatusComponents)(nil)

type managedStatusComponents struct {
	*statusComponents
	statusComponentsFactory *statusComponentsFactory
	cancelFunc              func()
	mutStatusComponents     sync.RWMutex
}

// NewManagedStatusComponents returns a new instance of managedStatusComponents
func NewManagedStatusComponents(scf *statusComponentsFactory) (*managedStatusComponents, error) {
	if scf == nil {
		return nil, errors.ErrNilStatusComponentsFactory
	}

	return &managedStatusComponents{
		statusComponents:        nil,
		statusComponentsFactory: scf,
		cancelFunc:              nil,
		mutStatusComponents:     sync.RWMutex{},
	}, nil
}

// Create will create the status components
func (msc *managedStatusComponents) Create() error {
	components, err := msc.statusComponentsFactory.Create()
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrStatusComponentsFactoryCreate, err)
	}

	msc.mutStatusComponents.Lock()
	msc.statusComponents = components
	msc.mutStatusComponents.Unlock()

	return nil
}

// Close will close all the underlying components
func (msc *managedStatusComponents) Close() error {
	msc.mutStatusComponents.Lock()
	defer msc.mutStatusComponents.Unlock()

	if msc.statusComponents == nil {
		return nil
	}
	if msc.cancelFunc != nil {
		msc.cancelFunc()
	}

	err := msc.statusComponents.Close()
	if err != nil {
		return err
	}
	msc.statusComponents = nil

	return nil
}

// CheckSubcomponents verifies all subcomponents
func (msc *managedStatusComponents) CheckSubcomponents() error {
	msc.mutStatusComponents.RLock()
	defer msc.mutStatusComponents.RUnlock()

	if msc.statusComponents == nil {
		return errors.ErrNilStatusComponents
	}
	if check.IfNil(msc.outportHandler) {
		return errors.ErrNilOutportHandler
	}
	if check.IfNil(msc.softwareVersion) {
		return errors.ErrNilSoftwareVersion
	}
	if check.IfNil(msc.statusHandler) {
		return errors.ErrNilStatusHandler
	}
	if check.IfNil(msc.managedPeersMonitor) {
		return errors.ErrNilManagedPeersMonitor
	}

	return nil
}

// SetForkDetector sets the fork detector
func (msc *managedStatusComponents) SetForkDetector(forkDetector process.ForkDetector) error {
	if check.IfNil(forkDetector) {
		return errors.ErrNilForkDetector
	}

	msc.mutStatusComponents.Lock()
	msc.statusComponentsFactory.forkDetector = forkDetector
	msc.mutStatusComponents.Unlock()

	return nil
}

// StartPolling starts polling for the updated status
func (msc *managedStatusComponents) StartPolling() error {
	var ctx context.Context
	msc.mutStatusComponents.Lock()
	ctx, msc.cancelFunc = context.WithCancel(context.Background())
	msc.mutStatusComponents.Unlock()

	err := msc.startStatusPolling(ctx)
	if err != nil {
		return err
	}

	err = msc.startMachineStatisticsPolling(ctx)
	if err != nil {
		return err
	}

	msc.attachEpochGoRoutineAnalyser()

	return nil
}

// OutportHandler returns the outport handler
func (msc *managedStatusComponents) OutportHandler() outport.OutportHandler {
	msc.mutStatusComponents.RLock()
	defer msc.mutStatusComponents.RUnlock()

	if msc.statusComponents == nil {
		return nil
	}

	return msc.statusComponents.outportHandler
}

// SoftwareVersionChecker returns the software version checker handler
func (msc *managedStatusComponents) SoftwareVersionChecker() statistics.SoftwareVersionChecker {
	msc.mutStatusComponents.RLock()
	defer msc.mutStatusComponents.RUnlock()

	if msc.statusComponents == nil {
		return nil
	}

	return msc.statusComponents.softwareVersion
}

// ManagedPeersMonitor returns the managed peers monitor
func (msc *managedStatusComponents) ManagedPeersMonitor() common.ManagedPeersMonitor {
	msc.mutStatusComponents.RLock()
	defer msc.mutStatusComponents.RUnlock()

	if msc.statusComponents == nil {
		return nil
	}

	return msc.managedPeersMonitor
}

// IsInterfaceNil returns true if there is no value under the interface
func (msc *managedStatusComponents) IsInterfaceNil() bool {
	return msc == nil
}

func (msc *managedStatusComponents) startStatusPolling(ctx context.Context) error {
	// TODO: inject the context to the AppStatusPolling
	appStatusPollingHandler, err := appStatusPolling.NewAppStatusPolling(
		msc.statusComponentsFactory.statusCoreComponents.AppStatusHandler(),
		time.Duration(msc.statusComponentsFactory.config.GeneralSettings.StatusPollingIntervalSec)*time.Second,
		log,
	)
	if err != nil {
		return errors.ErrStatusPollingInit
	}

	err = registerPollConnectedPeers(appStatusPollingHandler, msc.statusComponentsFactory.networkComponents)
	if err != nil {
		return err
	}

	err = registerPollProbableHighestNonce(appStatusPollingHandler, msc.statusComponentsFactory.forkDetector)
	if err != nil {
		return err
	}

	err = registerShardsInformation(appStatusPollingHandler, msc.statusComponentsFactory.shardCoordinator)
	if err != nil {
		return err
	}

	appStatusPollingHandler.Poll(ctx)

	return nil
}

func registerPollConnectedPeers(
	appStatusPollingHandler *appStatusPolling.AppStatusPolling,
	networkComponents factory.NetworkComponentsHolder,
) error {

	p2pMetricsHandlerFunc := func(appStatusHandler core.AppStatusHandler) {
		if check.IfNil(networkComponents) {
			return
		}
		mainNetMessenger := networkComponents.NetworkMessenger()
		if check.IfNil(mainNetMessenger) {
			return
		}
		computeMetricsForMessenger(appStatusHandler, mainNetMessenger, "")

		fullArchiveNetMessenger := networkComponents.FullArchiveNetworkMessenger()
		if check.IfNil(fullArchiveNetMessenger) {
			return
		}
		computeMetricsForMessenger(appStatusHandler, fullArchiveNetMessenger, common.FullArchiveMetricSuffix)
	}

	err := appStatusPollingHandler.RegisterPollingFunc(p2pMetricsHandlerFunc)
	if err != nil {
		return errors.ErrPollingFunctionRegistration
	}

	return nil
}

func registerShardsInformation(
	appStatusPollingHandler *appStatusPolling.AppStatusPolling,
	coordinator sharding.Coordinator,
) error {

	computeShardsInfo := func(appStatusHandler core.AppStatusHandler) {
		shardId := uint64(coordinator.SelfId())
		numOfShards := uint64(coordinator.NumberOfShards())

		appStatusHandler.SetUInt64Value(common.MetricShardId, shardId)
		appStatusHandler.SetUInt64Value(common.MetricNumShardsWithoutMetachain, numOfShards)
	}

	err := appStatusPollingHandler.RegisterPollingFunc(computeShardsInfo)
	if err != nil {
		return fmt.Errorf("%w, cannot register handler func for shards information", err)
	}

	return nil
}

func computeMetricsForMessenger(
	appStatusHandler core.AppStatusHandler,
	netMessenger p2p.Messenger,
	suffix string,
) {
	computeNumConnectedPeers(appStatusHandler, netMessenger, suffix)
	computeConnectedPeers(appStatusHandler, netMessenger, suffix)
}

func computeNumConnectedPeers(
	appStatusHandler core.AppStatusHandler,
	netMessenger p2p.Messenger,
	suffix string,
) {
	numOfConnectedPeers := uint64(len(netMessenger.ConnectedAddresses()))
	appStatusHandler.SetUInt64Value(common.SuffixedMetric(common.MetricNumConnectedPeers, suffix), numOfConnectedPeers)
}

func computeConnectedPeers(
	appStatusHandler core.AppStatusHandler,
	netMessenger p2p.Messenger,
	suffix string,
) {
	peersInfo := netMessenger.GetConnectedPeersInfo()

	peerClassification := fmt.Sprintf("intraVal:%d,crossVal:%d,intraObs:%d,crossObs:%d,unknown:%d,",
		len(peersInfo.IntraShardValidators),
		len(peersInfo.CrossShardValidators),
		len(peersInfo.IntraShardObservers),
		len(peersInfo.CrossShardObservers),
		len(peersInfo.UnknownPeers),
	)
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricNumConnectedPeersClassification, suffix), peerClassification)
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricP2PNumConnectedPeersClassification, suffix), peerClassification)

	setP2pConnectedPeersMetrics(appStatusHandler, peersInfo, suffix)
	setCurrentP2pNodeAddresses(appStatusHandler, netMessenger, suffix)
}

func setP2pConnectedPeersMetrics(appStatusHandler core.AppStatusHandler, info *p2p.ConnectedPeersInfo, suffix string) {
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricP2PUnknownPeers, suffix), sliceToString(info.UnknownPeers))
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricP2PIntraShardValidators, suffix), mapToString(info.IntraShardValidators))
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricP2PIntraShardObservers, suffix), mapToString(info.IntraShardObservers))
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricP2PCrossShardValidators, suffix), mapToString(info.CrossShardValidators))
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricP2PCrossShardObservers, suffix), mapToString(info.CrossShardObservers))
}

func sliceToString(input []string) string {
	return strings.Join(input, ",")
}

func mapToString(input map[uint32][]string) string {
	strs := make([]string, 0, len(input))
	keys := make([]uint32, 0, len(input))
	for shard := range input {
		keys = append(keys, shard)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, key := range keys {
		strs = append(strs, sliceToString(input[key]))
	}

	return strings.Join(strs, ",")
}

func setCurrentP2pNodeAddresses(
	appStatusHandler core.AppStatusHandler,
	netMessenger p2p.Messenger,
	suffix string,
) {
	appStatusHandler.SetStringValue(common.SuffixedMetric(common.MetricP2PPeerInfo, suffix), sliceToString(netMessenger.Addresses()))
}

func registerPollProbableHighestNonce(
	appStatusPollingHandler *appStatusPolling.AppStatusPolling,
	forkDetector process.ForkDetector,
) error {

	probableHighestNonceHandlerFunc := func(appStatusHandler core.AppStatusHandler) {
		probableHigherNonce := forkDetector.ProbableHighestNonce()
		appStatusHandler.SetUInt64Value(common.MetricProbableHighestNonce, probableHigherNonce)
	}

	err := appStatusPollingHandler.RegisterPollingFunc(probableHighestNonceHandlerFunc)
	if err != nil {
		return fmt.Errorf("%w, cannot register handler func for forkdetector's probable higher nonce", err)
	}

	return nil
}

func (msc *managedStatusComponents) startMachineStatisticsPolling(ctx context.Context) error {
	appStatusPollingHandler, err := appStatusPolling.NewAppStatusPolling(msc.statusComponentsFactory.statusCoreComponents.AppStatusHandler(), time.Second, log)
	if err != nil {
		return fmt.Errorf("%w, cannot init AppStatusPolling", err)
	}

	err = registerCpuStatistics(ctx, appStatusPollingHandler)
	if err != nil {
		return err
	}

	err = registerMemStatistics(ctx, appStatusPollingHandler)
	if err != nil {
		return err
	}

	netStats := msc.statusComponentsFactory.statusCoreComponents.NetworkStatistics()
	epochNotifier := msc.statusComponentsFactory.coreComponents.EpochNotifier()
	err = registerNetStatistics(appStatusPollingHandler, netStats, epochNotifier)
	if err != nil {
		return err
	}

	appStatusPollingHandler.Poll(ctx)

	return nil
}

func registerMemStatistics(_ context.Context, appStatusPollingHandler *appStatusPolling.AppStatusPolling) error {
	return appStatusPollingHandler.RegisterPollingFunc(func(appStatusHandler core.AppStatusHandler) {
		mem := machine.AcquireMemStatistics()

		appStatusHandler.SetUInt64Value(common.MetricMemLoadPercent, mem.PercentUsed)
		appStatusHandler.SetUInt64Value(common.MetricMemTotal, mem.Total)
		appStatusHandler.SetUInt64Value(common.MetricMemUsedGolang, mem.UsedByGolang)
		appStatusHandler.SetUInt64Value(common.MetricMemUsedSystem, mem.UsedBySystem)
		appStatusHandler.SetUInt64Value(common.MetricMemHeapInUse, mem.HeapInUse)
		appStatusHandler.SetUInt64Value(common.MetricMemStackInUse, mem.StackInUse)
	})
}

func registerNetStatistics(appStatusPollingHandler *appStatusPolling.AppStatusPolling, netStats factory.NetworkStatisticsProvider, epochNotifier process.EpochNotifier) error {
	epochNotifier.RegisterNotifyHandler(netStats)

	return appStatusPollingHandler.RegisterPollingFunc(func(appStatusHandler core.AppStatusHandler) {
		appStatusHandler.SetUInt64Value(common.MetricNetworkRecvBps, netStats.BpsRecv())
		appStatusHandler.SetUInt64Value(common.MetricNetworkRecvBpsPeak, netStats.BpsRecvPeak())
		appStatusHandler.SetUInt64Value(common.MetricNetworkRecvPercent, netStats.PercentRecv())

		appStatusHandler.SetUInt64Value(common.MetricNetworkSentBps, netStats.BpsSent())
		appStatusHandler.SetUInt64Value(common.MetricNetworkSentBpsPeak, netStats.BpsSentPeak())
		appStatusHandler.SetUInt64Value(common.MetricNetworkSentPercent, netStats.PercentSent())

		appStatusHandler.SetUInt64Value(common.MetricNetworkRecvBytesInCurrentEpochPerHost, netStats.TotalBytesReceivedInCurrentEpoch())
		appStatusHandler.SetUInt64Value(common.MetricNetworkSendBytesInCurrentEpochPerHost, netStats.TotalBytesSentInCurrentEpoch())
	})
}

func registerCpuStatistics(ctx context.Context, appStatusPollingHandler *appStatusPolling.AppStatusPolling) error {
	cpuStats, err := machine.NewCpuStatistics()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Debug("registerCpuStatistics go routine is stopping...")
				return
			default:
			}
			cpuStats.ComputeStatistics()
		}
	}()

	return appStatusPollingHandler.RegisterPollingFunc(func(appStatusHandler core.AppStatusHandler) {
		appStatusHandler.SetUInt64Value(common.MetricCpuLoadPercent, cpuStats.CpuPercentUsage())
	})
}

// String returns the name of the component
func (msc *managedStatusComponents) String() string {
	return factory.StatusComponentsName
}

func (msc *managedStatusComponents) attachEpochGoRoutineAnalyser() {
	currentConfig := msc.statusComponentsFactory.config
	enableEpochDebug := currentConfig.Debug.EpochStart.GoRoutineAnalyserEnabled
	log.Debug("attachEpochGoRoutineAnalyser", "GoRoutineAnalyserEnabled", enableEpochDebug)
	if enableEpochDebug {
		analyser, err := goroutine.NewGoRoutinesAnalyser(goroutine.NewGoRoutinesProcessor())

		if err != nil {
			log.Debug("attachEpochGoRoutineAnalyser", "error", err)
			return
		}

		subscribeHandler := notifier.NewHandlerForEpochStart(func(hdr data.HeaderHandler) {
			analyser.DumpGoRoutinesToLogWithTypes()
		}, func(_ data.HeaderHandler) {}, common.NetStatisticsOrder)

		msc.statusComponentsFactory.epochStartNotifier.RegisterHandler(subscribeHandler)
	}
}

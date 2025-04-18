package statistics_test

import (
	errorsGo "errors"
	"fmt"
	"testing"
	"time"

	logger "github.com/TerraDharitri/drt-go-chain-logger"
	stats "github.com/TerraDharitri/drt-go-chain/common/statistics"
	"github.com/TerraDharitri/drt-go-chain/common/statistics/disabled"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/stretchr/testify/assert"
)

var log = logger.GetOrCreate("common/statistics.test")

func generateMockConfig() config.Config {
	return config.Config{
		ResourceStats: config.ResourceStatsConfig{
			RefreshIntervalInSec: 1,
		},
	}
}

func TestNewResourceMonitor_NilNetStatisticsProviderShouldErr(t *testing.T) {
	t.Parallel()

	resourceMonitor, err := stats.NewResourceMonitor(
		generateMockConfig(),
		nil)

	assert.Equal(t, stats.ErrNilNetworkStatisticsProvider, err)
	assert.Nil(t, resourceMonitor)
}

func TestNewResourceMonitor_InvalidRefreshValueShouldErr(t *testing.T) {
	t.Parallel()

	resourceMonitor, err := stats.NewResourceMonitor(
		config.Config{
			ResourceStats: config.ResourceStatsConfig{
				RefreshIntervalInSec: 0,
			},
		},
		disabled.NewDisabledNetStatistics())

	assert.True(t, errorsGo.Is(err, stats.ErrInvalidRefreshIntervalValue))
	assert.Nil(t, resourceMonitor)
}

func TestResourceMonitor_NewResourceMonitorShouldWork(t *testing.T) {
	t.Parallel()

	resourceMonitor, err := stats.NewResourceMonitor(generateMockConfig(), disabled.NewDisabledNetStatistics())

	assert.Nil(t, err)
	assert.NotNil(t, resourceMonitor)
}

func TestResourceMonitor_GenerateStatisticsShouldPass(t *testing.T) {
	t.Parallel()

	resourceMonitor, err := stats.NewResourceMonitor(generateMockConfig(), disabled.NewDisabledNetStatistics())

	assert.Nil(t, err)
	statistics := resourceMonitor.GenerateStatistics()

	assert.NotNil(t, statistics)
	log.Info("sample statistics", statistics...)
}

func TestResourceMonitor_SaveStatisticsShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			assert.Fail(t, fmt.Sprintf("test should not have paniced: %v", r))
		}
	}()

	resourceMonitor, err := stats.NewResourceMonitor(generateMockConfig(), disabled.NewDisabledNetStatistics())

	assert.Nil(t, err)
	resourceMonitor.LogStatistics()
}

func TestResourceMonitor_StartMonitoringShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			assert.Fail(t, fmt.Sprintf("test should not have paniced: %v", r))
		}
	}()

	cfg := generateMockConfig()
	resourceMonitor, err := stats.NewResourceMonitor(cfg, disabled.NewDisabledNetStatistics())

	assert.Nil(t, err)
	resourceMonitor.StartMonitoring()
	threshold := time.Millisecond * 300
	time.Sleep(time.Second*time.Duration(cfg.ResourceStats.RefreshIntervalInSec) + threshold) // allow one loop

	assert.Nil(t, resourceMonitor.Close())
}

func TestResourceMonitor_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	resourceMonitor, _ := stats.NewResourceMonitor(generateMockConfig(), nil)
	assert.True(t, resourceMonitor.IsInterfaceNil())

	resourceMonitor, _ = stats.NewResourceMonitor(generateMockConfig(), disabled.NewDisabledNetStatistics())
	assert.False(t, resourceMonitor.IsInterfaceNil())
}

package coreComponents

import (
	"fmt"
	"testing"
	"time"

	"github.com/TerraDharitri/drt-go-chain/integrationTests/factory"
	"github.com/TerraDharitri/drt-go-chain/node"
	"github.com/TerraDharitri/drt-go-chain/testscommon/goroutines"
	"github.com/TerraDharitri/drt-go-chain-core/data/endProcess"
	"github.com/stretchr/testify/require"
)

// ------------ Test CoreComponents --------------------
func TestCoreComponents_Create_Close_ShouldWork(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	time.Sleep(time.Second * 4)

	gc := goroutines.NewGoCounter(goroutines.TestsRelevantGoRoutines)
	idxInitial, _ := gc.Snapshot()
	factory.PrintStack()

	configs := factory.CreateDefaultConfig(t)
	chanStopNodeProcess := make(chan endProcess.ArgEndProcess)
	nr, err := node.NewNodeRunner(configs)
	require.Nil(t, err)
	coreComponents, err := nr.CreateManagedCoreComponents(chanStopNodeProcess)
	require.Nil(t, err)
	require.NotNil(t, coreComponents)

	time.Sleep(2 * time.Second)

	err = coreComponents.Close()
	require.Nil(t, err)

	time.Sleep(2 * time.Second)

	idx, _ := gc.Snapshot()
	diff := gc.DiffGoRoutines(idxInitial, idx)
	require.Equal(t, 0, len(diff), fmt.Sprintf("%v", diff))
}

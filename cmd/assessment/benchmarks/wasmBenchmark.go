package benchmarks

import (
	"fmt"
	"time"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	wasmConfig "github.com/TerraDharitri/drt-go-chain-vm/config"
	wasmvm "github.com/TerraDharitri/drt-go-chain/integrationTests/vm/wasm/wasmvm"
	"github.com/TerraDharitri/drt-go-chain/vm/systemSmartContracts/defaults"
)

// ArgWasmBenchmark is the wasm type benchmark argument used in constructor
type ArgWasmBenchmark struct {
	Name         string
	ScFilename   string
	TestingValue uint64
	Function     string
	Arguments    [][]byte
	NumRuns      int
}

type wasmBenchmark struct {
	name         string
	scFilename   string
	testingValue uint64
	function     string
	arguments    [][]byte
	numRuns      int
}

// NewWasmBenchmark creates a new benchmark based on SC execution through wasm VM
func NewWasmBenchmark(arg ArgWasmBenchmark) *wasmBenchmark {
	return &wasmBenchmark{
		name:         arg.Name,
		scFilename:   arg.ScFilename,
		testingValue: arg.TestingValue,
		function:     arg.Function,
		arguments:    arg.Arguments,
		numRuns:      arg.NumRuns,
	}
}

// Run returns the time needed for the benchmark to be run
func (wb *wasmBenchmark) Run() (time.Duration, error) {
	if !core.FileExists(wb.scFilename) {
		return 0, fmt.Errorf("%w, file %s", ErrFileDoesNotExist, wb.scFilename)
	}

	result, err := wasmvm.RunTest(wb.scFilename, wb.testingValue, wb.function, wb.arguments, wb.numRuns, createTestGasMap(), 0)
	if err != nil {
		return 0, err
	}

	return result.ExecutionTimeSpan, err
}

func createTestGasMap() map[string]map[string]uint64 {
	gasSchedule := wasmConfig.MakeGasMapForTests()
	gasSchedule = defaults.FillGasMapInternal(gasSchedule, 1)

	return gasSchedule
}

// Name returns the benchmark's name
func (wb *wasmBenchmark) Name() string {
	return fmt.Sprintf("%s, function %s, numRuns %d", wb.name, wb.function, wb.numRuns)
}

// IsInterfaceNil returns true if there is no value under the interface
func (wb *wasmBenchmark) IsInterfaceNil() bool {
	return wb == nil
}

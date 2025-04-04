package testscommon

import "github.com/TerraDharitri/drt-go-chain-core/display"

// TableDisplayerMock -
type TableDisplayerMock struct {
	DisplayTableCalled func(tableHeader []string, lines []*display.LineData, message string)
}

// DisplayTable -
func (mock *TableDisplayerMock) DisplayTable(tableHeader []string, lines []*display.LineData, message string) {
	if mock.DisplayTableCalled != nil {
		mock.DisplayTableCalled(tableHeader, lines, message)
	}
}

func (mock *TableDisplayerMock) IsInterfaceNil() bool {
	return mock == nil
}

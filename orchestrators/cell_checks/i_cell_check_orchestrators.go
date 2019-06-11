package cell_checks

type ICellCheckOrchestrators interface {
	RunCellChecks() ([][]interface{}, []interface{})
}

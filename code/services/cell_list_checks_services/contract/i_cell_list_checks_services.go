package contract

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
)

type ICellListChecksServices interface {
	Set_syntactic_checks_results()
	Get_cell_list_checks_result() service_results.CellListChecksResults
	Set_cell_checks_result(in_scope_cell cells.Cells, cell_checks_result service_results.CellChecksResults)
	Get_cell_list_checks_parameter() service_parameters.CellListChecksParameters
}

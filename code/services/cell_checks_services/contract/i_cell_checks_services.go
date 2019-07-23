package contract

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
)

type ICellChecksServices interface {
	Set_cell_checks_result()
	Get_cell_checks_result() service_results.CellChecksResults
	Get_cell_checks_parameter() service_parameters.CellChecksParameters
	Set_issues_result([]issues.Issues)
	Set_fixes_result(cell_checks_fix fixes.Fixes)
}

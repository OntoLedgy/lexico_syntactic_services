package contract

import (
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/service_parameters"
)

type ICellCheckServices interface {
	Set_cell_check_result()
	Set_cell_check_result_value(check_result *check_results.CheckResults)
	Get_check_parameter() *service_parameters.CellCheckParameters
	Get_check_result() *check_results.CheckResults
}

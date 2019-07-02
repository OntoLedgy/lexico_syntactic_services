package cell_check_services

import (
	"syntactic_checker/code/object_model/check_results"
)

type ICellCheckService interface {
	Set_cell_check_result()
	Get_check_result() *check_results.CheckResults
}

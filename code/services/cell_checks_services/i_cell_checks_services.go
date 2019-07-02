package cell_checks_services

import (
	"syntactic_checker/code/object_model/service_results"
)

type ICellChecksServices interface {
	Set_cell_checks_result()
	Get_cell_checks_result() service_results.CellChecksResults
}

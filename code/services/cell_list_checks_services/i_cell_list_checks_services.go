package cell_list_checks_services

import (
	"syntactic_checker/code/object_model/service_results"
)

type ICellListChecksService interface {
	Set_syntactic_checks_results()
	Get_cell_list_checks_result() service_results.CellListChecksResults
}

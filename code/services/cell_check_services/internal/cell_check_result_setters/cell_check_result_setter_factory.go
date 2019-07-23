package cell_check_result_setters

import (
	"syntactic_checker/code/services/cell_check_services/contract"
)

func Create(cell_check_service contract.ICellCheckServices) *CellCheckResultSetters {

	cell_check_results_setter := new(CellCheckResultSetters)

	cell_check_results_setter.ICellCheckServices = cell_check_service

	return cell_check_results_setter
}

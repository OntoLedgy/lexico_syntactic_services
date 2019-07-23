package internal

import (
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_check_services/internal/cell_check_result_setters"
)

type CellCheckServices struct {
	Cell_check_parameter service_parameters.CellCheckParameters
	Check_result         *check_results.CheckResults
}

func (
	cell_check_service *CellCheckServices) Set_cell_check_result() {

	cell_check_result_setter :=
		cell_check_result_setters.
			Create(cell_check_service)

	cell_check_result_setter.
		Set_cell_check_result()

}

func (
	cell_check_service *CellCheckServices) Set_cell_check_result_value(
	check_result *check_results.CheckResults) {

	cell_check_service.
		Check_result =
		check_result
}

func (
	cell_check_service *CellCheckServices) Get_check_result() *check_results.CheckResults {

	return cell_check_service.Check_result
}

func (
	cell_check_service *CellCheckServices) Get_check_parameter() *service_parameters.CellCheckParameters {

	return &cell_check_service.Cell_check_parameter
}

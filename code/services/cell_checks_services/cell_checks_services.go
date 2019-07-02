package cell_checks_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
)

type cellChecksService struct {
	Cell_checks_parameter service_parameters.CellChecksParameters
	Cell_checks_result    service_results.CellChecksResults
}

func (
	cell_checks_service *cellChecksService) Set_cell_checks_result() {

	cell_check_issues :=
		cell_checks_service.
			check_cell_for_issues()

	cell_checks_service.
		set_cell_issues_and_fix(
			cell_check_issues)
}

func (cell_checks_service *cellChecksService) Get_cell_checks_result() service_results.CellChecksResults {

	return cell_checks_service.Cell_checks_result
}

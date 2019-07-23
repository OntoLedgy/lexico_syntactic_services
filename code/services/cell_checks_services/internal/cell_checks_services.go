package internal

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/cell_checks_services/internal/cell_checks_result_setters"
)

type CellChecksService struct {
	Cell_checks_parameter service_parameters.CellChecksParameters
	Cell_checks_result    service_results.CellChecksResults
}

func (
	cell_checks_service *CellChecksService) Set_cell_checks_result() {

	cell_checks_result_setters.
		Set_cell_issues_and_fix(
			cell_checks_service)

}

func (cell_checks_service *CellChecksService) Get_cell_checks_result() service_results.CellChecksResults {

	return cell_checks_service.Cell_checks_result
}

func (cell_checks_service *CellChecksService) Get_cell_checks_parameter() service_parameters.CellChecksParameters {

	return cell_checks_service.Cell_checks_parameter
}

func (cell_checks_service *CellChecksService) Set_issues_result(cell_checks_issues []issues.Issues) {

	cell_checks_service.Cell_checks_result.Cell_checks_issues = cell_checks_issues

}

func (cell_checks_service *CellChecksService) Set_fixes_result(cell_checks_fix fixes.Fixes) {

	cell_checks_service.Cell_checks_result.Cell_checks_fix = cell_checks_fix

}

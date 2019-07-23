package internal

import (
	"fmt"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/cell_list_checks_services/internal/cell_list_checks_processors"
)

type CellListChecksServices struct {
	Cell_list_checks_parameter service_parameters.CellListChecksParameters
	Cell_list_checks_result    service_results.CellListChecksResults
}

func (
	cell_list_checks_service *CellListChecksServices) Set_syntactic_checks_results() {

	fmt.Println(
		"\nRunning cell list checks service...")

	cell_list_checks_processor :=
		cell_list_checks_processors.
			Create(
				cell_list_checks_service)

	cell_list_checks_processor.
		Process_cell_list_for_cell_checks()

}

func (
	cell_list_checks_service *CellListChecksServices) Get_cell_list_checks_result() service_results.CellListChecksResults {

	return cell_list_checks_service.Cell_list_checks_result
}

func (
	cell_list_checks_service *CellListChecksServices) Set_cell_checks_result(
	in_scope_cell cells.Cells,
	cell_checks_result service_results.CellChecksResults) {

	there_are_issues :=
		cell_checks_result.
			Cell_checks_issues !=
			nil

	if there_are_issues {

		cell_checks_result.
			In_scope_cell =
			in_scope_cell

		cell_list_checks_service.
			Cell_list_checks_result.
			Cell_list_checks_results =
			append(
				cell_list_checks_service.
					Cell_list_checks_result.
					Cell_list_checks_results,
				cell_checks_result)

	}

}

func (
	cell_list_checks_service *CellListChecksServices) Get_cell_list_checks_parameter() service_parameters.CellListChecksParameters {

	return cell_list_checks_service.Cell_list_checks_parameter
}

package cell_list_checks_services

import (
	"fmt"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/cell_checks_services"
)

func (
	cell_list_checks_service *cellListChecksService) process_cell_list_for_cell_checks() {

	fmt.Printf(
		"Processing checks for cell list: %s\n",
		cell_list_checks_service.Cell_list_checks_parameter.List_of_in_scope_issue_types)

	in_scope_cells :=
		cell_list_checks_service.
			Cell_list_checks_parameter.
			List_of_in_scope_cells.
			Cells

	cell_list_checks_service.
		iterate_cell_list_for_cell_checks(
			in_scope_cells)

}

func (
	cell_list_checks_service *cellListChecksService) iterate_cell_list_for_cell_checks(
	in_scope_cells []cells.Cells) {

	for _, in_scope_cell := range in_scope_cells {

		cell_list_checks_service.
			process_cell_checks(
				in_scope_cell)

	}
}

func (
	cell_list_checks_service *cellListChecksService) process_cell_checks(
	in_scope_cell cells.Cells) {

	cell_checks_service_parameter :=
		cell_list_checks_service.
			generate_cell_checks_service_parameters(
				in_scope_cell)

	cell_checks_result :=
		cell_list_checks_service.
			generate_cell_checks_result(
				cell_checks_service_parameter)

	cell_list_checks_service.
		process_cell_checks_result(
			in_scope_cell,
			cell_checks_result)

}

//---------------------------split 1 : run cell checks or get cell checks result
func (
	cell_list_checks_service *cellListChecksService) generate_cell_checks_service_parameters(
	in_scope_cell cells.Cells) *service_parameters.CellChecksParameters {

	in_scope_issue_types :=
		cell_list_checks_service.
			Cell_list_checks_parameter.
			List_of_in_scope_issue_types

	cell_checks_parameter :=
		new(
			service_parameters.
				CellChecksParameters)

	cell_checks_parameter.
		In_scope_cell =
		in_scope_cell

	cell_checks_parameter.
		List_of_in_scope_issue_types =
		in_scope_issue_types

	return cell_checks_parameter
}

func (
	cell_list_checks_service *cellListChecksService) generate_cell_checks_result(
	cell_checks_parameter *service_parameters.CellChecksParameters) service_results.CellChecksResults {

	cell_checks_service_factory :=
		new(
			cell_checks_services.
				CellChecksServiceFactory)

	cell_checks_service :=
		cell_checks_service_factory.Create(
			*cell_checks_parameter)

	cell_checks_service.
		Set_cell_checks_result()

	cell_checks_result :=
		cell_checks_service.
			Get_cell_checks_result()

	return cell_checks_result
}

//------------------------------split 2: process cell checks result

func (
	cell_list_checks_service *cellListChecksService) process_cell_checks_result(
	in_scope_cell cells.Cells,
	cell_checks_result service_results.CellChecksResults) {

	there_are_issues :=
		cell_checks_result.
			Cell_checks_issues != nil

	if there_are_issues {

		cell_checks_result.
			In_scope_cell =
			in_scope_cell

		cell_list_checks_service.
			append_cell_checks_result(
				cell_checks_result)

	}

}

func (
	cell_list_checks_service *cellListChecksService) append_cell_checks_result(
	cell_checks_result service_results.CellChecksResults) {

	cell_list_checks_service.
		Cell_list_checks_result.
		Cell_list_checks_results =
		append(
			cell_list_checks_service.Cell_list_checks_result.Cell_list_checks_results,
			cell_checks_result)

}

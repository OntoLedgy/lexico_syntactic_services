package cell_list_checks_processors

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/cell_checks_services"
)

func Strip_cell_identifier_and_run_string_checks(
	in_scope_cell cells.Cells,
	issue_types []issues.IssueTypes) service_results.CellChecksResults {

	cell_checks_service_parameter :=
		generate_cell_checks_service_parameters(
			in_scope_cell, // TODO - strip identifiers here and send string
			issue_types)

	cell_checks_result :=
		generate_cell_checks_result(
			cell_checks_service_parameter)

	return cell_checks_result

}

func generate_cell_checks_service_parameters(
	in_scope_cell cells.Cells,
	in_scope_issue_types []issues.IssueTypes) *service_parameters.CellChecksParameters {

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

func generate_cell_checks_result(
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

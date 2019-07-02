package cell_check_services

import (
	"fmt"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_check_services/internal/regex_checkers"
)

type cellCheckService struct {
	Cell_check_parameter service_parameters.CellCheckParameters
	Check_results        *check_results.CheckResults
}

func (
	cell_check_service *cellCheckService) Set_cell_check_result() {

	//TODO - Stage 3 - add switch to include non-regex in_scope_check types.

	cell_value :=
		cell_check_service.
			Cell_check_parameter.
			In_scope_cell.
			Cell_value

	check_regex :=
		cell_check_service.
			Cell_check_parameter.
			In_scope_issue_type.
			Issue_check_regex

	cell_value_is_not_empty :=
		cell_value != ""

	if cell_value_is_not_empty {

		cell_check_service.
			process_regex_check(
				check_regex,
				cell_value)

	} else {

		fmt.Printf(
			"\nWARNING: In_scope_cell value for row_id:[%s] is null\n",
			cell_check_service.
				Cell_check_parameter.
				In_scope_cell.Cell_identifier)

	}

}

func (
	cell_check_service *cellCheckService) Get_check_result() *check_results.CheckResults {

	return cell_check_service.Check_results
}

func (
	cell_check_service *cellCheckService) process_regex_check(
	check_regex string,
	cell_value string) {

	regex_checker_factory :=
		new(
			regex_checkers.
				RegexCheckerFactories)

	regex_checker :=
		regex_checker_factory.
			Create()

	regex_checker.
		Process_regex_check(
			check_regex,
			cell_value)

	string_edit_ranges :=
		regex_checker.
			String_edit_ranges

	check_results :=
		new(
			check_results.
				CheckResults)

	check_results.
		Check_result_string_edit_ranges =
		string_edit_ranges

	cell_check_service.
		Check_results =
		check_results

}

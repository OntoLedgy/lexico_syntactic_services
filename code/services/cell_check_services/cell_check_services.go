package cell_check_services

import (
	"fmt"
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_check_services/internal/deprecate_cell_editors"
	"syntactic_checker/code/services/cell_check_services/internal/regex_checkers"
)

type cellCheckService struct {
	Cell_check_parameter service_parameters.CellCheckParameters
	Cell_edit_history    *string_editor_object_model.StringEditHistory
	Regex_check_result   *regex_checkers.RegexCheckResults
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
			"\nWARNING: In_scope_cell value for row_id:%s is null\n",
			cell_value)

	}

}

func (
	cell_check_service *cellCheckService) Get_cell_regex_check_result() *regex_checkers.RegexCheckResults {

	return cell_check_service.Regex_check_result
}

func (
	cell_check_service *cellCheckService) Get_cell_edit_history() *string_editor_object_model.StringEditHistory {

	return cell_check_service.Cell_edit_history
}

func (
	cell_check_service *cellCheckService) Get_check_result() *check_results.CheckResults {

	return cell_check_service.Check_results
}

func (
	cell_check_service *cellCheckService) process_regex_check(
	check_regex string,
	cell_value string) {

	var regex_check_result *regex_checkers.RegexCheckResults

	regex_checker_factory :=
		new(
			regex_checkers.
				RegexCheckerFactories)

	regex_checker :=
		regex_checker_factory.
			Create()

	regex_check_result =
		regex_checker.
			Process_regex_check(
				check_regex,
				cell_value)

	cell_check_service.
		Regex_check_result =
		regex_check_result

	string_edit_ranges :=
		regex_checker.
			String_edit_ranges

	check_results := new(check_results.CheckResults)

	check_results.Check_result_edit_string_ranges = string_edit_ranges

	cell_check_service.Check_results =
		check_results

	cell_check_service.
		process_regex_result()

}

func (
	cell_check_service *cellCheckService) process_regex_result() {

	there_is_a_regex_result :=
		cell_check_service.Regex_check_result != nil

	if there_is_a_regex_result {

		//deprecate

		cell_editor := deprecate_cell_editors.
			Create(
				cell_check_service.Cell_check_parameter.In_scope_cell,
				cell_check_service.Cell_check_parameter.In_scope_issue_type,
				*cell_check_service.Regex_check_result)

		cell_value_edit_history :=
			cell_editor.
				Edit_cell()

		cell_check_service.
			Cell_edit_history =
			cell_value_edit_history

	}
}

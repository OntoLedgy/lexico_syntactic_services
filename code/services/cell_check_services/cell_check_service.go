package cell_check_services

import (
	"fmt"
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_services/cell_editors"
	"syntactic_checker/code/services/cell_check_services/regex_checkers"
)

type CellCheckService struct {
	Issue_type              issues.IssueTypes
	In_scope_cell           object_model.Cells
	Regex_check_result      *regex_checkers.RegexCheckResults
	Cell_value_edit_history string_editor_object_model.StringEditHistory
}

func (
	cell_check_service *CellCheckService) Set_cell_check_results() {

	var regex_check_result *regex_checkers.RegexCheckResults

	//TODO - Stage 3 - add switch to include non-regex in_scope_check types.

	if cell_check_service.In_scope_cell.Cell_value != "" {

		regex_check_result =
			regex_checkers.
				Process_regex_check(
					cell_check_service.Issue_type.Issue_check_regex,
					cell_check_service.In_scope_cell.Cell_value)

		cell_check_service.
			Regex_check_result =
			regex_check_result

		cell_check_service.
			process_regex_result()

	} else {

		fmt.Printf(
			"\nWARNING: In_scope_cell value for row_id:%s is null\n",
			cell_check_service.In_scope_cell.Cell_value)

	}

}

func (
	cell_check_service *CellCheckService) process_regex_result() {

	if cell_check_service.Regex_check_result != nil {

		cell_editor := cell_editors.
			Create_cell_editor(
				cell_check_service.In_scope_cell,
				cell_check_service.Issue_type,
				*cell_check_service.Regex_check_result)

		cell_value_edit_history :=
			cell_editor.
				Edit_cell()

		cell_check_service.
			Cell_value_edit_history =
			*cell_value_edit_history
	}

}

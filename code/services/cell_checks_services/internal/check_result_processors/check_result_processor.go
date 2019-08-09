package check_result_processors

import (
	"string_editor/object_model"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_checks_services/internal/cell_editors"
)

type CheckResultProcessors struct {
	Check_results       *check_results.CheckResults
	In_scope_cell       cells.Cells
	In_scope_issue_type issues.IssueTypes
	Cell_edit_history   *object_model.StringEditHistory
}

func (
	check_result_processor *CheckResultProcessors) Process_regex_result() {

	there_is_a_regex_result :=
		check_result_processor.
			Check_results != nil

	if there_is_a_regex_result {

		check_result_processor.
			set_string_edit_history()

	}
}

func (
	check_result_processor *CheckResultProcessors) set_string_edit_history() {

	replacement_string :=
		check_result_processor.
			In_scope_issue_type.
			Get_replacement_string()

	cell_editor :=
		cell_editors.
			Create(
				check_result_processor.In_scope_cell,
				check_result_processor.Check_results,
				replacement_string)

	cell_value_edit_history :=
		cell_editor.
			Edit_cell()

	check_result_processor.
		Cell_edit_history =
		cell_value_edit_history
}

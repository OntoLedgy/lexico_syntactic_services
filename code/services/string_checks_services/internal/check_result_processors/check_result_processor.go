package check_result_processors

import (
	"string_editor/object_model"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_checks_services/internal/string_editors"
)

type CheckResultProcessors struct {
	Check_results       *check_results.CheckResults
	Identified_string   identified_strings.IdentifiedStrings
	String_value        string
	In_scope_issue_type issues.IssueTypes
	String_edit_history *object_model.StringEditHistory
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

	string_editor :=
		string_editors.
			Create(
				check_result_processor.Identified_string,
				check_result_processor.String_value,
				check_result_processor.Check_results,
				replacement_string)

	string_edit_history :=
		string_editor.
			Edit_string()

	check_result_processor.
		String_edit_history =
		string_edit_history
}

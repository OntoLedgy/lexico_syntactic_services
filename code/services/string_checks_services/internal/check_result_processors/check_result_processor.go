package check_result_processors

import (
	"string_editor/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/string_checks_services/internal/check_result_processors/string_editors"
)

type checkResultProcessors struct {
	check_results *service_results.StringCheckResults
	//string_value        string
	in_scope_issue_type issues.IssueTypes
	String_edit_history *object_model.StringEditHistories
}

func (
	check_result_processor *checkResultProcessors) Process_regex_result() {

	there_is_a_regex_result :=
		check_result_processor.
			check_results != nil

	if there_is_a_regex_result {

		check_result_processor.
			set_string_edit_history()

	}
}

func (
	check_result_processor *checkResultProcessors) set_string_edit_history() {

	replacement_string :=
		check_result_processor.
			in_scope_issue_type.
			Get_replacement_string()

	string_editor :=
		string_editors.
			Create(
				check_result_processor.String_edit_history.GetCurrentString(),
				check_result_processor.check_results,
				replacement_string)

	check_result_processor.String_edit_history =
		string_editor.
			Edit_string(check_result_processor.String_edit_history)

}

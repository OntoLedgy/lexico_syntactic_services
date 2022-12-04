package check_result_processors

import (
	"github.com/OntoLedgy/string_editing_services/object_model"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/internal/check_result_processors/string_editors"
)

type checkResultProcessors struct {
	check_results       *service_results.StringCheckResults
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

	if check_result_processor.in_scope_issue_type.Issue_replacement_parameters.Is_replaceable == "TRUE" {

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
	} else {
		check_result_processor.
			String_edit_history.
			SetStringChanges("Not Applicable", "Not Applicable")
	}
}

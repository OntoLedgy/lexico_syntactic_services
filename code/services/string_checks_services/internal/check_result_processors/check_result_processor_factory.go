package check_result_processors

import (
	"string_editor/object_model"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/object_model/issues"
)

func Create(
	string_edit_history *object_model.StringEditHistories,
	check_result *service_results.StringCheckResults,
	issue_type issues.IssueTypes) *checkResultProcessors {

	check_result_processor :=
		new(
			checkResultProcessors)

	check_result_processor.
		check_results =
		check_result

	check_result_processor.
		in_scope_issue_type =
		issue_type

	check_result_processor.
		String_edit_history =
		string_edit_history

	return check_result_processor
}

package check_result_processors

import (
	"string_editor/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_check_services/contract"
)

func Create(
	string_edit_history *object_model.StringEditHistories,
	string_check_service contract.IStringCheckServices,
	issue_type issues.IssueTypes) *checkResultProcessors {

	check_result_processor :=
		new(
			checkResultProcessors)

	check_result_processor.
		check_results =
		string_check_service.
			Get_string_check_result()

	check_result_processor.
		in_scope_issue_type =
		issue_type

	check_result_processor.
		String_edit_history =
		string_edit_history

	return check_result_processor
}

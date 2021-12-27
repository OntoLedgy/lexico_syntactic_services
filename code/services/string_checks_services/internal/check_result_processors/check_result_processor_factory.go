package check_result_processors

import (
	"github.com/OntoLedgy/string_editing_services/object_model"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
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

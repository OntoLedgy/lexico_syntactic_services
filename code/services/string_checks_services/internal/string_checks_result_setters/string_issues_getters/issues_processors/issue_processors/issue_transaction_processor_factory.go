package issue_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
)

func Create_issue_result_processor(
	string_value *identified_strings.Strings,
	issue_type *issues.IssueTypes) *IssueResultProcessors {

	issue_result_processor :=
		new(IssueResultProcessors)

	issue_result_processor.
		string_value =
		string_value

	issue_result_processor.
		issue_type =
		issue_type
	return issue_result_processor
}

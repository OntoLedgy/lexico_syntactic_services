package issue_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
)

func Create(
	string_value *identified_strings.Strings,
	issue_type issues.IssueTypes) *IssueCheckResultProcessors {

	issue_processor := new(IssueCheckResultProcessors)

	issue_processor.string_value = string_value
	issue_processor.issue_type = issue_type

	return issue_processor

}

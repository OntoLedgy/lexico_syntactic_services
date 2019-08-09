package issues_processor

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

func Create(
	string_value string,
	identified_string identified_strings.IdentifiedStrings,
	in_scope_issue_types []issues.IssueTypes) *issuesProcessors {

	issues_processor :=
		new(
			issuesProcessors)

	issues_processor.
		identified_string =
		identified_string

	issues_processor.
		string_value =
		string_value

	issues_processor.
		in_scope_issue_types =
		in_scope_issue_types

	return issues_processor

}

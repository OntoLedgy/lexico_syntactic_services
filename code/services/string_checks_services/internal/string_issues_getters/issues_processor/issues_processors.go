package issues_processor

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type issuesProcessors struct {
	in_scope_issue_types []issues.IssueTypes
	identified_string    identified_strings.IdentifiedStrings
	string_value         string
	string_checks_issues []issues.Issues
}

func (
	issues_processor *issuesProcessors) Set_string_check_issues() {

	in_scope_issue_types :=
		issues_processor.
			in_scope_issue_types

	for _, in_scope_issue_type := range in_scope_issue_types {

		string_check_issue :=
			issues_processor.
				get_string_check_issue(
					in_scope_issue_type)

		issues_processor.
			process_string_check_issue(
				string_check_issue,
				in_scope_issue_type)
	}

}

func (
	issues_processor *issuesProcessors) Get_string_checks_issues() []issues.Issues {

	return issues_processor.string_checks_issues

}

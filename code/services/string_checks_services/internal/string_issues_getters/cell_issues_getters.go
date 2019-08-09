package string_issues_getters

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_checks_services/internal/string_issues_getters/issues_processor"
)

type StringIssuesGetters struct {
}

//TODO - make this a method for the type

func Get_string_issues(
	identified_string identified_strings.IdentifiedStrings,
	string_value string,
	list_of_in_scope_issue_types []issues.IssueTypes) []issues.Issues {

	issues_processor :=
		issues_processor.
			Create(
				string_value,
				identified_string,
				list_of_in_scope_issue_types)

	issues_processor.
		Set_string_check_issues()

	string_checks_issues :=
		issues_processor.
			Get_string_checks_issues()

	return string_checks_issues
}

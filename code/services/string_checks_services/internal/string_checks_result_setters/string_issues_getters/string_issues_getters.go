package string_issues_getters

import (
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_issues_getters/issues_processors"
)

type StringIssuesGetters struct {
	string_checks_parameter service_parameters.StringChecksParameters
	String_checks_issues    []issues.Issues
}

func (
	issue_getter *StringIssuesGetters) Get_string_checks_issues() []issues.Issues {

	issues_processor :=
		issues_processors.
			Create(
				issue_getter.string_checks_parameter)

	issues_processor.
		Set_string_check_issues()

	issue_getter.String_checks_issues =
		issues_processor.
			Get_string_checks_issues()

	return issue_getter.String_checks_issues
}

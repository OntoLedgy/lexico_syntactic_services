package string_issues_getters

import (
	"syntactic_checker/code/object_model/service_parameters"
)

type StringIssuesGetterFactory struct{}

func (string_issues_getter_factory *StringIssuesGetterFactory) Create(
	string_checks_parameter service_parameters.StringChecksParameters) *StringIssuesGetters {

	sting_issue_getter := new(StringIssuesGetters)

	sting_issue_getter.
		string_checks_parameter =
		string_checks_parameter

	return sting_issue_getter
}

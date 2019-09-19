package string_issues_getters

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
)

type StringIssuesGetterFactory struct{}

func (string_issues_getter_factory *StringIssuesGetterFactory) Create(
	string_checks_parameter service_inputs.StringChecksInputs) *StringIssuesGetters {

	sting_issue_getter := new(StringIssuesGetters)

	sting_issue_getter.
		string_checks_input =
		string_checks_parameter

	return sting_issue_getter
}

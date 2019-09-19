package string_issues_getters

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_issues_getters/issues_processors"
)

type StringIssuesGetters struct {
	string_checks_input      service_inputs.StringChecksInputs
	issue_check_results_list service_results.IssueChecksResultLists
	String_checks_issues     []issues.Issues
}

func (
	string_issue_getter *StringIssuesGetters) Get_string_checks_issues_results_list() *service_results.IssueChecksResultLists {

	issues_processor :=
		issues_processors.
			Create(
				string_issue_getter.
					string_checks_input)

	issues_processor.
		Set_string_check_issues()

	string_check_issues_results_list :=
		issues_processor.
			Get_string_checks_issues()

	return string_check_issues_results_list
}

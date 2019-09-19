package issues_processors

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_issues_getters/issues_processors/issue_processors"
)

type issuesProcessors struct {
	string_checks_parameter         service_inputs.StringChecksInputs
	string_issue_checks_result_list *service_results.IssueChecksResultLists
}

func (
	issues_processor *issuesProcessors) Set_string_check_issues() {

	in_scope_issue_types :=
		issues_processor.
			string_checks_parameter.
			Issue_types

	issues_processor.
		iterate_issue_types_and_run_checks(
			in_scope_issue_types)

}

func (
	issues_processor *issuesProcessors) Get_string_checks_issues() *service_results.IssueChecksResultLists {

	return issues_processor.string_issue_checks_result_list

}

func (
	issues_processor *issuesProcessors) iterate_issue_types_and_run_checks(
	in_scope_issue_types []issues.IssueTypes) {

	for _, in_scope_issue_type := range in_scope_issue_types {

		issues_processor.
			get_and_set_string_issue_check_result(
				in_scope_issue_type)
	}

}

func (
	issues_processor *issuesProcessors) get_and_set_string_issue_check_result(
	in_scope_issue_type issues.IssueTypes) {

	issue_processor :=
		issue_processors.
			Create(
				issues_processor.string_checks_parameter.String_to_check,
				in_scope_issue_type)

	string_check_issue_result :=
		issue_processor.
			Get_string_issue_check_result(
				in_scope_issue_type)

	issues_processor.
		append_string_issue_check_result(
			string_check_issue_result,
			in_scope_issue_type)
}

func (
	issues_processor *issuesProcessors) append_string_issue_check_result(
	string_check_issue_result *service_results.IssueCheckResults,
	in_scope_issue_type issues.IssueTypes) {

	there_is_an_issue :=
		string_check_issue_result != nil

	if there_is_an_issue {

		string_check_issue_result.
			String_checks_issue.
			Issue_type =
			in_scope_issue_type

		issues_processor.
			string_issue_checks_result_list =
			new(
				service_results.IssueChecksResultLists)

		issues_processor.
			string_issue_checks_result_list.
			String_checks_issue_results = append(
			issues_processor.string_issue_checks_result_list.String_checks_issue_results,
			string_check_issue_result)
	}

}

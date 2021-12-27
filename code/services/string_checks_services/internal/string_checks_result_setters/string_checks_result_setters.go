package string_checks_result_setters

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/contract"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_fix_getters"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_issues_getters"
)

type StringChecksResultSetters struct {
	contract.IStringChecksServices
}

func (
	string_checks_result_setter StringChecksResultSetters) Set_string_issues_and_fix() {

	string_checks_issues_result :=
		string_checks_result_setter.
			get_string_checks_issue_results()

	there_are_issues := string_checks_issues_result != nil

	if there_are_issues {
		string_checks_result_setter.
			Set_string_issues_result(
				string_checks_issues_result)
	}

	there_are_issues_for_automatic_fixing :=
		string_checks_result_setter.
			are_there_issues_for_automatic_fixing(string_checks_issues_result)

	if there_are_issues_for_automatic_fixing {

		string_check_fix_result :=
			string_checks_result_setter.
				get_string_checks_fix()

		string_checks_result_setter.
			Set_string_fixes_result(
				string_check_fix_result)

	}
}

func (
	string_checks_result_setter StringChecksResultSetters) are_there_issues_for_automatic_fixing(
	string_checks_issues_result *service_results.IssueChecksResultLists) bool {

	there_are_issues_for_automatic_fixing :=
		false

	if string_checks_issues_result != nil {
		for _, result := range string_checks_issues_result.String_checks_issue_results {
			if result.String_checks_issue.Issue_type.Issue_severity_level == "Automatic Fixing" {
				there_are_issues_for_automatic_fixing = true
			}
		}
	}

	return there_are_issues_for_automatic_fixing
}

func (
	string_checks_result_setter StringChecksResultSetters) get_string_checks_issue_results() *service_results.IssueChecksResultLists {

	string_checks_input :=
		string_checks_result_setter.
			Get_string_checks_input()

	string_issue_getter_factory :=
		new(
			string_issues_getters.
				StringIssuesGetterFactory)

	string_issue_getter :=
		string_issue_getter_factory.Create(
			*string_checks_input)

	string_checks_issues :=
		string_issue_getter.
			Get_string_checks_issues_results_list()

	return string_checks_issues
}

func (
	string_checks_result_setter StringChecksResultSetters) get_string_checks_fix() service_results.FixChecksResults {

	string_fix_getter_factory :=
		new(
			string_fix_getters.StringFixGetterFactory)

	string_fix_getter :=
		string_fix_getter_factory.
			Create(
				*string_checks_result_setter.
					Get_string_checks_input())

	string_check_fix :=
		string_fix_getter.
			Get_string_check_fix()

	return string_check_fix

}

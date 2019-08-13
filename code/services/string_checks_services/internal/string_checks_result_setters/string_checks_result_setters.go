package string_checks_result_setters

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_checks_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_fix_getters"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_issues_getters"
)

type StringChecksResultSetters struct {
	contract.IStringChecksServices
}

func (
	string_checks_result_setter StringChecksResultSetters) Set_string_issues_and_fix() {

	string_checks_issues :=
		string_checks_result_setter.
			get_string_checks_issues()

	there_are_issues :=
		string_checks_issues != nil

	if there_are_issues {

		string_check_fix :=
			string_checks_result_setter.
				get_string_checks_fix()

		//TODO - Merge these into a single setter?

		string_checks_result_setter.
			Set_issues_result(
				string_checks_issues)

		string_checks_result_setter.
			Set_string_fixes_result(
				string_check_fix)

	}
}

func (
	string_checks_result_setter StringChecksResultSetters) get_string_checks_issues() []issues.Issues {

	string_checks_parameter :=
		string_checks_result_setter.
			Get_string_checks_parameter()

	string_issue_getter_factory :=
		new(
			string_issues_getters.
				StringIssuesGetterFactory)

	string_issue_getter :=
		string_issue_getter_factory.Create(
			string_checks_parameter)

	string_checks_issues :=
		string_issue_getter.
			Get_string_checks_issues()

	return string_checks_issues
}

func (
	string_checks_result_setter StringChecksResultSetters) get_string_checks_fix() fixes.Fixes {

	string_fix_getter_factory :=
		new(
			string_fix_getters.StringFixGetterFactory)

	string_fix_getter :=
		string_fix_getter_factory.
			Create(
				string_checks_result_setter.
					Get_string_checks_parameter())

	string_check_fix :=
		string_fix_getter.
			Get_string_check_fix()

	return string_check_fix

}

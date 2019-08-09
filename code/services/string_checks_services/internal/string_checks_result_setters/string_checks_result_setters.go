package string_checks_result_setters

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_checks_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal/string_fix_getters"
	"syntactic_checker/code/services/string_checks_services/internal/string_issues_getters"
)

type StringChecksResultSetters struct {
	identified_string            identified_strings.IdentifiedStrings //Deprecate
	string_value                 string
	list_of_in_scope_issue_types []issues.IssueTypes
}

//TODO - convert this to method

func Set_string_issues_and_fix(
	string_checks_service contract.IStringChecksServices) {

	string_checks_parameter :=
		string_checks_service.
			Get_string_checks_parameter()

	identified_string :=
		string_checks_parameter.
			Identified_string

	string_value :=
		string_checks_parameter.
			String_value

	list_of_in_scope_issue_types :=
		string_checks_parameter.
			In_scope_issue_types

	string_checks_issues :=
		string_issues_getters.
			Get_string_issues(
				identified_string,
				string_value,
				list_of_in_scope_issue_types)

	there_are_issues :=
		string_checks_issues != nil

	if there_are_issues {

		string_checks_service.
			Set_issues_result(
				string_checks_issues)

		string_check_fix :=
			get_string_check_fix(
				string_checks_parameter)

		string_checks_service.
			Set_string_fixes_result(
				string_check_fix)

	}
}

func get_string_check_fix(
	string_checks_parameter service_parameters.StringChecksParameters) fixes.Fixes {

	fix_processor_factory :=
		new(
			string_fix_getters.FixProcessorsFactory)

	fix_processor :=
		fix_processor_factory.
			Create(string_checks_parameter)

	string_check_fix :=
		fix_processor.
			Get_string_check_fix(
				string_checks_parameter.In_scope_issue_types,
				string_checks_parameter.Identified_string)

	return string_check_fix

}

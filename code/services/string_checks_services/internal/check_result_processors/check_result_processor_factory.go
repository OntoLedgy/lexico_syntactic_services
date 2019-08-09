package check_result_processors

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_check_services/contract"
)

func Create(
	string_check_service contract.IStringCheckServices,
	issue_type issues.IssueTypes,
	identified_string identified_strings.IdentifiedStrings,
	string_value string) *CheckResultProcessors {

	check_result_processor :=
		new(
			CheckResultProcessors)

	check_result_processor.
		Check_results =
		string_check_service.
			Get_check_result()

	check_result_processor.
		In_scope_issue_type =
		issue_type

	check_result_processor.
		Identified_string =
		identified_string

	check_result_processor.
		String_value =
		string_value

	return check_result_processor
}

package identified_string_list_checks_processors

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/string_checks_services"
)

func strip_identified_string_identifier_and_run_string_checks(
	identified_string identified_strings.IdentifiedStrings,
	issue_types []issues.IssueTypes) service_results.StringChecksResults {

	identified_string_checks_service_parameter :=
		generate_string_checks_service_parameters(
			identified_string, // TODO - strip identifiers here and send string
			issue_types)

	identified_string_checks_result :=
		generate_identified_string_checks_result(
			identified_string_checks_service_parameter)

	return identified_string_checks_result

}

func generate_string_checks_service_parameters(
	identified_string identified_strings.IdentifiedStrings,
	in_scope_issue_types []issues.IssueTypes) *service_parameters.StringChecksParameters {

	string_checks_parameter :=
		new(
			service_parameters.
				StringChecksParameters)

	string_checks_parameter.
		Identified_string =
		identified_string

	string_checks_parameter.String_value =
		identified_string.String_value

	string_checks_parameter.
		In_scope_issue_types =
		in_scope_issue_types

	return string_checks_parameter
}

func generate_identified_string_checks_result(
	identified_string_checks_parameter *service_parameters.StringChecksParameters) service_results.StringChecksResults {

	string_checks_service_factory :=
		new(
			string_checks_services.
				StringChecksServiceFactory)

	string_checks_service :=
		string_checks_service_factory.Create(
			*identified_string_checks_parameter)

	string_checks_service.
		Set_string_checks_result()

	identified_string_checks_result :=
		string_checks_service.
			Get_string_checks_result()

	return identified_string_checks_result
}

package identified_string_checks_processors

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/string_checks_services"
)

type identifiedStringChecksProcessor struct {
	identified_string    identified_strings.IdentifiedStrings
	issue_types          []issues.IssueTypes
	string_checks_result service_results.StringChecksResults
}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) Get_string_checks_result() service_results.StringChecksResults {

	identified_string_checks_processor.
		strip_identified_string_identifier_and_run_string_checks()

	return identified_string_checks_processor.string_checks_result

}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) strip_identified_string_identifier_and_run_string_checks() {

	string_checks_parameter :=
		identified_string_checks_processor.
			generate_string_checks_service_parameters()

	identified_string_checks_processor.
		set_identified_string_checks_result(
			string_checks_parameter)

}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) generate_string_checks_service_parameters() *service_parameters.StringChecksParameters {

	string_checks_parameter :=
		new(
			service_parameters.
				StringChecksParameters)

	string_checks_parameter.String_value =
		identified_string_checks_processor.
			identified_string.
			String_value

	string_checks_parameter.
		In_scope_issue_types =
		identified_string_checks_processor.
			issue_types

	return string_checks_parameter
}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) set_identified_string_checks_result(
	identified_string_checks_parameter *service_parameters.StringChecksParameters) {

	string_checks_service_factory :=
		new(
			string_checks_services.
				StringChecksServiceFactory)

	string_checks_service :=
		string_checks_service_factory.Create(
			*identified_string_checks_parameter)

	string_checks_service.
		Set_string_checks_result()

	identified_string_checks_processor.
		string_checks_result =
		string_checks_service.
			Get_string_checks_result()

}

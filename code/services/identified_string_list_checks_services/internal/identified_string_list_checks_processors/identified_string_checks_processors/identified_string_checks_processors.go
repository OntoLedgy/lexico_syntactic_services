package identified_string_checks_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services"
)

type identifiedStringChecksProcessor struct {
	identified_string    *identified_strings.IdentifiedStrings
	issue_types          []issues.IssueTypes
	string_checks_result service_results.StringChecksResults
}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) Get_string_checks_result() service_results.StringChecksResults {

	identified_string_checks_processor.
		strip_string_identifier_and_run_string_checks()

	return identified_string_checks_processor.string_checks_result

}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) strip_string_identifier_and_run_string_checks() {

	string_checks_input :=
		identified_string_checks_processor.
			generate_string_checks_service_input()

	identified_string_checks_processor.
		set_identified_string_checks_result(
			string_checks_input)

}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) generate_string_checks_service_input() *service_inputs.StringChecksInputs {

	string_checks_input :=
		new(
			service_inputs.
				StringChecksInputs)

	string_checks_input.
		StringToCheck = new(identified_strings.Strings)

	string_checks_input.
		StringToCheck.StringValue =
		identified_string_checks_processor.
			identified_string.
			String_identified.
			StringValue

	string_checks_input.
		IssueTypes =
		identified_string_checks_processor.
			issue_types

	return string_checks_input
}

func (
	identified_string_checks_processor *identifiedStringChecksProcessor) set_identified_string_checks_result(
	identified_string_checks_input *service_inputs.StringChecksInputs) {

	string_checks_service_factory :=
		new(
			string_checks_services.
				StringChecksServiceFactory)

	string_checks_service :=
		string_checks_service_factory.Create(
			identified_string_checks_input)

	string_checks_service.
		Set_string_checks_result()

	identified_string_checks_processor.
		string_checks_result =
		*string_checks_service.
			Get_string_checks_result()

}

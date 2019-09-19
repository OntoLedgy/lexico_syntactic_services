package internal

import (
	"logger/standard_global_logger"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/interservice_i_o_objects"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/services/identified_string_list_checks_services/internal/identified_string_list_checks_processors"
)

type IdentifiedStringListChecksServices struct {
	Identified_string_list_checks_i_o_object *interservice_i_o_objects.IdentifiedStringListChecksIOObjects
}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Set_syntactic_checks_results() {

	standard_global_logger.Global_logger.Println(
		"\nRunning identified string list checks service...")

	identified_string_list_checks_processor :=
		identified_string_list_checks_processors.
			Create(
				identified_string_list_checks_service)

	identified_string_list_checks_processor.
		Process_identified_string_list_for_checks()

}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Get_identified_string_list_checks_result() *service_results.IdentifiedStringListChecksResults {

	return identified_string_list_checks_service.Identified_string_list_checks_i_o_object.Identified_string_list_checks_results
}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Set_identified_string_checks_result(
	identified_string *identified_strings.IdentifiedStrings,
	string_checks_result *service_results.StringChecksResults) {

	there_are_issues :=
		string_checks_result.
			String_checks_issues_list !=
			nil

	if there_are_issues {

		identified_string_list_checks_service.
			append_string_checks_result_to_list(
				identified_string,
				string_checks_result)

	}

}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Get_identified_string_list_checks_input() *service_inputs.IdentifiedStringListChecksInputs {

	return identified_string_list_checks_service.Identified_string_list_checks_i_o_object.Identified_string_list_checks_input
}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) append_string_checks_result_to_list(
	identified_string *identified_strings.IdentifiedStrings,
	string_checks_result *service_results.StringChecksResults) {

	identified_string_checks_result :=
		new(service_results.
			IdentifiedStringChecksResults)

	identified_string_checks_result.
		Identified_string =
		identified_string

	identified_string_checks_result.
		String_checks_result =
		string_checks_result

	identified_string_list_checks_service.
		Identified_string_list_checks_i_o_object.
		Identified_string_list_checks_results.
		Identified_string_checks_results =
		append(
			identified_string_list_checks_service.
				Identified_string_list_checks_i_o_object.
				Identified_string_list_checks_results.
				Identified_string_checks_results,
			identified_string_checks_result)
}

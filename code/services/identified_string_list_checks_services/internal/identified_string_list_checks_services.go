package internal

import (
	"fmt"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/identified_string_list_checks_services/internal/identified_string_list_checks_processors"
)

type IdentifiedStringListChecksServices struct {
	Identified_string_list_checks_parameter service_parameters.IdentifiedStringListChecksParameters
	Identified_string_list_checks_result    service_results.IdentifiedStringListChecksResults
}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Set_syntactic_checks_results() {

	fmt.Println(
		"\nRunning identified string list checks service...")

	identified_string_list_checks_processor :=
		identified_string_list_checks_processors.
			Create(
				identified_string_list_checks_service)

	identified_string_list_checks_processor.
		Process_identified_string_list_for_checks()

}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Get_identified_string_list_checks_result() service_results.IdentifiedStringListChecksResults {

	return identified_string_list_checks_service.Identified_string_list_checks_result
}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Set_identified_string_checks_result(
	identified_string identified_strings.IdentifiedStrings,
	identified_string_checks_result service_results.StringChecksResults) {

	there_are_issues :=
		identified_string_checks_result.
			Identified_string_checks_issues !=
			nil

	if there_are_issues {

		identified_string_checks_result.
			Identified_string =
			identified_string

		identified_string_list_checks_service.
			Identified_string_list_checks_result.
			Identified_string_list_checks_results =
			append(
				identified_string_list_checks_service.
					Identified_string_list_checks_result.
					Identified_string_list_checks_results,
				identified_string_checks_result)

	}

}

func (
	identified_string_list_checks_service *IdentifiedStringListChecksServices) Get_identified_string_list_checks_parameter() service_parameters.IdentifiedStringListChecksParameters {

	return identified_string_list_checks_service.Identified_string_list_checks_parameter
}

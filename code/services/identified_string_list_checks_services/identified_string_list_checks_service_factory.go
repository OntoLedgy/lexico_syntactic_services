package identified_string_list_checks_services

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/services/identified_string_list_checks_services/contract"
	"syntactic_checker/code/services/identified_string_list_checks_services/internal"
)

type IdentifiedStringListChecksServiceFactory struct{}

func (IdentifiedStringListChecksServiceFactory) Create(
	identified_string_list_checks_input *service_inputs.IdentifiedStringListChecksInputs) contract.IIdentifiedStringListChecksServices {

	identified_string_list_checks_service :=
		new(
			internal.IdentifiedStringListChecksServices)

	identified_string_list_checks_i_o_object :=
		new(
			interservice_i_o_objects.IdentifiedStringListChecksIOObjects)

	identified_string_list_checks_i_o_object.
		Identified_string_list_checks_results =
		new(
			service_results.IdentifiedStringListChecksResults)

	identified_string_list_checks_i_o_object.
		Identified_string_list_checks_input =
		identified_string_list_checks_input

	identified_string_list_checks_service.
		Identified_string_list_checks_i_o_object =
		identified_string_list_checks_i_o_object

	return identified_string_list_checks_service
}

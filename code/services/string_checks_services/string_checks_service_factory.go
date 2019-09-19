package string_checks_services

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/services/string_checks_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal"
)

type StringChecksServiceFactory struct{}

func (
	StringChecksServiceFactory) Create(
	string_checks_input *service_inputs.StringChecksInputs) contract.IStringChecksServices {

	string_checks_service :=
		new(
			internal.StringChecksService)

	string_checks_service.
		String_checks_i_o_object =
		new(
			interservice_i_o_objects.StringChecksIOObjects)

	string_checks_service.
		String_checks_i_o_object.
		String_checks_result =
		new(
			service_results.StringChecksResults)

	string_checks_service.
		String_checks_i_o_object.
		String_checks_input =
		string_checks_input

	return string_checks_service
}

package string_check_services

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_check_services/contract"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_check_services/internal"
)

type StringCheckServiceFactory struct{}

func (
	StringCheckServiceFactory) Create(
	string_check_input *service_inputs.StringCheckInputs) contract.IStringCheckServices {

	string_check_service :=
		new(
			internal.StringCheckServices)

	string_check_service.
		String_check_i_o_object =
		new(interservice_i_o_objects.StringCheckIOObjects)

	string_check_service.
		String_check_i_o_object.
		String_check_result =
		new(service_results.StringCheckResults)

	string_check_service.
		String_check_i_o_object.
		String_check_input =
		string_check_input

	return string_check_service

}

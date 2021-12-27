package internal

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_check_services/internal/string_check_result_setters"
)

type StringCheckServices struct {
	String_check_i_o_object *interservice_i_o_objects.StringCheckIOObjects
}

func (
	string_check_service *StringCheckServices) Set_string_check_result() {

	string_check_result_setter :=
		string_check_result_setters.
			Create(string_check_service)

	string_check_result_setter.
		Generate_and_set_string_check_result()

}

func (
	string_check_service *StringCheckServices) Set_string_check_result_value(
	string_check_result *service_results.StringCheckResults) {

	string_check_service.
		String_check_i_o_object.
		String_check_result =
		string_check_result
}

func (
	string_check_service *StringCheckServices) Get_string_check_result() *service_results.StringCheckResults {

	return string_check_service.String_check_i_o_object.String_check_result
}

func (
	string_check_service *StringCheckServices) Get_check_parameter() *service_inputs.StringCheckInputs {

	return string_check_service.String_check_i_o_object.String_check_input
}

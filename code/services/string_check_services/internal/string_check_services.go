package internal

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/string_check_services/internal/string_check_result_setters"
)

type StringCheckServices struct {
	String_check_parameter service_parameters.StringCheckParameters
	String_Check_result    *service_results.StringCheckResults
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
		String_Check_result =
		string_check_result
}

func (
	string_check_service *StringCheckServices) Get_string_check_result() *service_results.StringCheckResults {

	return string_check_service.String_Check_result
}

func (
	string_check_service *StringCheckServices) Get_check_parameter() *service_parameters.StringCheckParameters {

	return &string_check_service.String_check_parameter
}

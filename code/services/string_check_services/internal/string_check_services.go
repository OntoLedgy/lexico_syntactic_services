package internal

import (
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_check_services/internal/string_check_result_setters"
)

type StringCheckServices struct {
	String_check_parameter service_parameters.StringCheckParameters
	Check_result           *check_results.CheckResults
}

func (
	string_check_service *StringCheckServices) Set_string_check_result() {

	string_check_result_setter :=
		string_check_result_setters.
			Create(string_check_service)

	string_check_result_setter.
		Set_string_check_result()

}

func (
	string_check_service *StringCheckServices) Set_string_check_result_value(
	check_result *check_results.CheckResults) {

	string_check_service.
		Check_result =
		check_result
}

func (
	string_check_service *StringCheckServices) Get_check_result() *check_results.CheckResults {

	return string_check_service.Check_result
}

func (
	string_check_service *StringCheckServices) Get_check_parameter() *service_parameters.StringCheckParameters {

	return &string_check_service.String_check_parameter
}

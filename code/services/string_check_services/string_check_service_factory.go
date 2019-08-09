package string_check_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_check_services/contract"
	"syntactic_checker/code/services/string_check_services/internal"
)

type StringCheckServiceFactory struct{}

func (
	StringCheckServiceFactory) Create(
	string_check_parameter service_parameters.StringCheckParameters) contract.IStringCheckServices {

	string_check_service :=
		new(
			internal.StringCheckServices)

	string_check_service.
		String_check_parameter =
		string_check_parameter

	return string_check_service

}

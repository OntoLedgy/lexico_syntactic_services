package string_checks_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_checks_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal"
)

type StringChecksServiceFactory struct{}

func (
	StringChecksServiceFactory) Create(
	string_checks_parameter service_parameters.StringChecksParameters) contract.IStringChecksServices {

	string_checks_service :=
		new(
			internal.StringChecksService)

	string_checks_service.
		String_checks_parameter =
		string_checks_parameter

	return string_checks_service
}

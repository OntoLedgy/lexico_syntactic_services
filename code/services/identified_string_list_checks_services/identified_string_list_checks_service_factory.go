package identified_string_list_checks_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/identified_string_list_checks_services/contract"
	"syntactic_checker/code/services/identified_string_list_checks_services/internal"
)

type IdentifiedStringListChecksServiceFactory struct{}

func (IdentifiedStringListChecksServiceFactory) Create(
	identified_string_list_checks_parameter service_parameters.IdentifiedStringListChecksParameters) contract.IIdentifiedStringListChecksServices {

	identified_string_list_checks_service :=
		new(
			internal.IdentifiedStringListChecksServices)

	identified_string_list_checks_service.
		Identified_string_list_checks_parameter =
		identified_string_list_checks_parameter

	return identified_string_list_checks_service
}

package identified_string_list_checks_processors

import "github.com/OntoLedgy/syntactic_checker/code/services/identified_string_list_checks_services/contract"

func Create(
	identified_string_list_checks_service contract.IIdentifiedStringListChecksServices) *IdentifiedStringListChecksProcessors {

	identified_string_list_checks_processor :=
		new(
			IdentifiedStringListChecksProcessors)

	identified_string_list_checks_processor.
		IIdentifiedStringListChecksServices =
		identified_string_list_checks_service

	return identified_string_list_checks_processor

}

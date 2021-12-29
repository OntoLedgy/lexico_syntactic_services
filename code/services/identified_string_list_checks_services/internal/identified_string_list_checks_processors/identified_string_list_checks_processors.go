package identified_string_list_checks_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/infrastructure/logging"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/services/identified_string_list_checks_services/contract"
	"github.com/OntoLedgy/syntactic_checker/code/services/identified_string_list_checks_services/internal/identified_string_list_checks_processors/identified_string_checks_processors"
)

type IdentifiedStringListChecksProcessors struct {
	contract.IIdentifiedStringListChecksServices
}

func (
	identified_string_list_checks_processor *IdentifiedStringListChecksProcessors) Process_identified_string_list_for_checks() {

	identified_string_list_checks_processor.
		iterate_identified_string_list_for_checks()

}

func (
	identified_string_list_checks_processor *IdentifiedStringListChecksProcessors) iterate_identified_string_list_for_checks() {

	logger := logging.GlobalLogger

	identified_string_list_checks_parameter :=
		identified_string_list_checks_processor.
			Get_identified_string_list_checks_input()

	logger.Printf("Processing checks for identified string list: \n",
		identified_string_list_checks_parameter.
			List_of_in_scope_issue_types)

	identified_string_list :=
		identified_string_list_checks_parameter.
			Identified_string_list.
			Identified_string_list

	for _, identified_string := range identified_string_list {

		identified_string_list_checks_processor.
			process_and_set_identified_string_checks_result(
				identified_string)

	}
}

func (
	identified_string_list_checks_processor *IdentifiedStringListChecksProcessors) process_and_set_identified_string_checks_result(
	identified_string *identified_strings.IdentifiedStrings) {

	identified_string_list_checks_input :=
		identified_string_list_checks_processor.
			Get_identified_string_list_checks_input()

	issue_types :=
		identified_string_list_checks_input.
			List_of_in_scope_issue_types

	identified_string_check_processor :=
		identified_string_checks_processors.Create(
			identified_string,
			issue_types)

	string_checks_result :=
		identified_string_check_processor.
			Get_string_checks_result()

	identified_string_list_checks_processor.
		Set_identified_string_checks_result(
			identified_string,
			&string_checks_result)
}

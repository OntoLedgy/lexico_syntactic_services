package identified_string_list_checks_processors

import (
	"fmt"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/services/identified_string_list_checks_services/contract"
)

type IdentifiedStringListChecksProcessors struct {
	contract.IIdentifiedStringListChecksServices
}

func (identified_string_list_checks_processor *IdentifiedStringListChecksProcessors) Process_identified_string_list_for_checks() {

	identified_string_list_checks_processor.
		iterate_identified_string_list_for_checks()

}

func (identified_string_list_checks_processor *IdentifiedStringListChecksProcessors) iterate_identified_string_list_for_checks() {

	identified_string_list_checks_parameter :=
		identified_string_list_checks_processor.
			Get_identified_string_list_checks_parameter()

	fmt.Printf("Processing checks for identified string list: \n",
		identified_string_list_checks_parameter.
			List_of_in_scope_issue_types)

	identified_strings :=
		identified_string_list_checks_parameter.
			Identified_string_list.
			Identified_strings

	for _, identified_string := range identified_strings {

		identified_string_list_checks_processor.
			process_and_set_identified_string_checks_result(
				identified_string)

	}
}

func (
	identified_string_list_checks_processor *IdentifiedStringListChecksProcessors) process_and_set_identified_string_checks_result(
	identified_string identified_strings.IdentifiedStrings) {

	identified_string_list_checks_parameter :=
		identified_string_list_checks_processor.
			Get_identified_string_list_checks_parameter()

	issue_types :=
		identified_string_list_checks_parameter.
			List_of_in_scope_issue_types

	identified_string_checks_result :=
		strip_identified_string_identifier_and_run_string_checks(
			identified_string,
			issue_types)

	identified_string_list_checks_processor.
		Set_identified_string_checks_result(
			identified_string,
			identified_string_checks_result)
}

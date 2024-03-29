package string_fix_getters

import (
	"github.com/OntoLedgy/syntactic_checker/code/infrastructure/logging"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_fix_getters/fix_processors"
)

type StringFixGetters struct {
	string_checks_input service_inputs.StringChecksInputs
	String_checks_fix   service_results.FixChecksResults
}

func (
	string_fix_getter *StringFixGetters) Get_string_check_fix() service_results.FixChecksResults {
	logger := logging.GlobalLogger

	logger.Printf(
		"\nProcessing fixes for %s...\n",
		string_fix_getter.string_checks_input.StringToCheck)

	fix_processor :=
		fix_processors.Create(
			string_fix_getter.
				string_checks_input)

	fix_processor.
		SetStringChecksFix()

	logger.Printf(
		"\nFix transaction: %v \n",
		fix_processor.FixChecksResults)

	return fix_processor.FixChecksResults

}

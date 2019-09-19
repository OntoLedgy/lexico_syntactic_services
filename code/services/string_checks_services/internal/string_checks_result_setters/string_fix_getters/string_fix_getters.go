package string_fix_getters

import (
	//"fmt"
	"logger/standard_global_logger"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_fix_getters/fix_processors"
)

type StringFixGetters struct {
	string_checks_input service_inputs.StringChecksInputs
	String_checks_fix   service_results.FixChecksResults
}

func (
	string_fix_getter *StringFixGetters) Get_string_check_fix() service_results.FixChecksResults {
	logger := standard_global_logger.Global_logger

	logger.Printf(
		"\nProcessing fixes for %s...\n",
		string_fix_getter.string_checks_input.String_to_check)

	fix_processor :=
		fix_processors.Create(
			string_fix_getter.
				string_checks_input)

	fix_processor.
		Set_string_checks_fix()

	logger.Printf(
		"\nFix transaction: %v \n",
		fix_processor.Fix_check_result)

	return fix_processor.Fix_check_result

}

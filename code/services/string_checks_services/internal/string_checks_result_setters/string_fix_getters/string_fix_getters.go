package string_fix_getters

import (
	"fmt"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_fix_getters/fix_processors"
)

type StringFixGetters struct {
	string_checks_parameter service_parameters.StringChecksParameters
	String_checks_fix       fixes.Fixes
}

func (
	string_fix_getter *StringFixGetters) Get_string_check_fix() fixes.Fixes {

	fmt.Printf(
		"\nProcessing fixes for %s...\n",
		string_fix_getter.string_checks_parameter.String_value)

	fix_processor := fix_processors.Create(
		string_fix_getter.
			string_checks_parameter)

	fix_processor.
		Set_string_checks_fix()

	fmt.Printf(
		"\nFix transaction: %v \n",
		fix_processor.String_checks_fix)

	return fix_processor.String_checks_fix

}

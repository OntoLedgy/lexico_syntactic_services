package string_check_result_setters

import (
	"syntactic_checker/code/services/string_check_services/contract"
)

type StringCheckResultSetters struct {
	contract.IStringCheckServices
}

func (
	string_check_results_setter *StringCheckResultSetters) Set_string_check_result() {

	string_check_parameter :=
		string_check_results_setter.
			Get_check_parameter()

	string_value :=
		string_check_parameter.
			Identified_string.
			String_value

	//string_value = string_check_parameter.String_value

	check_regex :=
		string_check_parameter.
			In_scope_issue_type.
			Issue_check_regex

	//TODO - Stage 3 - Add other check types (non - regex)

	string_check_results_setter.
		Set_string_check_result_value(
			Get_regex_check_result(
				check_regex,
				string_value))

}

package string_check_result_setters

import (
	"fmt"
	"syntactic_checker/code/services/string_check_services/contract"
	"syntactic_checker/code/services/string_check_services/internal/string_check_result_setters/regex_check_result_getter"
)

type StringCheckResultSetters struct {
	contract.IStringCheckServices
}

func (
	string_check_results_setter *StringCheckResultSetters) Generate_and_set_string_check_result() {

	string_check_parameter :=
		string_check_results_setter.
			Get_check_parameter()

	string_value :=
		string_check_parameter.
			String_value

	check_regex :=
		string_check_parameter.
			In_scope_issue_type.
			Issue_check_regex

	if string_check_parameter.In_scope_issue_type.Issue_check_type == "regex" {

		regex_checker :=
			regex_check_result_getter.
				Create(
					check_regex,
					string_value)

		string_check_result :=
			regex_checker.
				Get_regex_check_result()

		string_check_results_setter.
			Set_string_check_result_value(
				string_check_result)

	} else {
		fmt.Println("running non-regex checks")
	}

}

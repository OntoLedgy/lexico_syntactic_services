package string_check_result_setters

import (
	//"fmt"
	"logger/standard_global_logger"
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
			String_to_check

	check_regex :=
		string_check_parameter.
			In_scope_issue_type.
			Issue_check_regex

	if string_check_parameter.In_scope_issue_type.Issue_check_type == "regex-re2" {

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
		logger := standard_global_logger.Global_logger

		logger.Printf("running non-regex-re2 checks: %s \n",
			string_check_parameter.
				In_scope_issue_type.
				Issue_check_type)
	}

}

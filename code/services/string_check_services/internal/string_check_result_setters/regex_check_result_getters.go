package string_check_result_setters

import (
	"fmt"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/services/string_check_services/internal/regex_checkers"
)

type RegexCheckResultGetters struct {
}

func Get_regex_check_result(
	check_regex string,
	string_value string) *check_results.CheckResults {

	string_value_is_not_empty :=
		string_value != ""

	if string_value_is_not_empty {

		check_results := get_result(
			check_regex,
			string_value)

		return check_results

	} else {
		fmt.Printf(
			"\nWARNING: Identified_string value :[%s] is null\n",
			string_value)

		return nil

	}
}

func get_result( //TODO - rename
	check_regex string,
	string_value string) *check_results.CheckResults {

	regex_checker_factory :=
		new(
			regex_checkers.RegexCheckerFactories)

	regex_checker :=
		regex_checker_factory.
			Create()

	regex_checker.
		Process_regex_check(
			check_regex,
			string_value)

	string_edit_ranges :=
		regex_checker.
			String_edit_ranges

	check_results :=
		new(
			check_results.
				CheckResults)

	check_results.
		Check_result_string_edit_ranges =
		string_edit_ranges

	return check_results
}

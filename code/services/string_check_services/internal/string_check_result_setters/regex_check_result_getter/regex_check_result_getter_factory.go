package regex_check_result_getter

import "syntactic_checker/code/object_model/identified_strings"

func Create(
	check_regex string,
	string_value *identified_strings.Strings) *RegexCheckResultGetter {

	regex_checker := new(RegexCheckResultGetter)

	regex_checker.
		string_to_check =
		string_value

	regex_checker.
		check_regex_pattern =
		check_regex

	return regex_checker

}

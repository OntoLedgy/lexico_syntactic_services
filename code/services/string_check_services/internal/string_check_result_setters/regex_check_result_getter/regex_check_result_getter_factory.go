package regex_check_result_getter

func Create(
	check_regex string,
	string_value string) *RegexCheckResultGetter {

	regex_checker := new(RegexCheckResultGetter)

	regex_checker.
		string_value =
		string_value

	regex_checker.
		check_regex_pattern =
		check_regex

	return regex_checker

}

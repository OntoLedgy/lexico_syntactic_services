package regex_checkers

type RegexCheckerFactory struct {
}

func (
	*RegexCheckerFactory) Create(
	string_value string,
	check_regex_pattern string) *regexCheckers {

	regex_checker :=
		new(
			regexCheckers)

	regex_checker.
		check_regex_pattern =
		check_regex_pattern

	regex_checker.
		string_value =
		string_value

	return regex_checker
}

package regex_checkers

type RegexCheckerFactories struct {
}

func (
	regex_checker_factory *RegexCheckerFactories) Create() *regexCheckers {

	regex_checker :=
		new(
			regexCheckers)

	return regex_checker
}

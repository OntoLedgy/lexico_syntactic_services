package regex_check_result_getter

import (
	"github.com/OntoLedgy/syntactic_checker/code/infrastructure/logging"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_check_services/internal/regex_checkers"
)

type RegexCheckResultGetter struct {
	check_regex_pattern string                      //TODO - wrap regex pattern as an object
	string_to_check     *identified_strings.Strings //TODO - wrap string as an object
}

func (
	regex_checks_result_getter *RegexCheckResultGetter) Get_regex_check_result() *service_results.StringCheckResults {

	string_value_is_not_empty :=
		regex_checks_result_getter.string_to_check.String_value != ""

	if string_value_is_not_empty {

		string_check_result :=
			regex_checks_result_getter.
				generate_string_check_result()

		return string_check_result

	} else {
		logging.GlobalLogger.Printf(
			"\nWARNING: Identified_string value :[%s] is null\n",
			regex_checks_result_getter.string_to_check)

		return nil

	}
}

func (
	regex_checks_result_getter *RegexCheckResultGetter) generate_string_check_result() *service_results.StringCheckResults {

	regex_checker_factory :=
		new(
			regex_checkers.RegexCheckerFactory)

	regex_checker :=
		regex_checker_factory.
			Create(
				regex_checks_result_getter.string_to_check.String_value,
				regex_checks_result_getter.check_regex_pattern)

	regex_checker.
		Process_regex_check()

	string_edit_ranges :=
		regex_checker.
			String_edit_ranges

	check_results :=
		new(
			service_results.
				StringCheckResults)

	check_results.
		Check_result_string_edit_ranges =
		string_edit_ranges

	return check_results
}

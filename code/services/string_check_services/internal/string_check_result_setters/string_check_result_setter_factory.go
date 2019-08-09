package string_check_result_setters

import (
	"syntactic_checker/code/services/string_check_services/contract"
)

func Create(string_check_service contract.IStringCheckServices) *StringCheckResultSetters {

	string_check_results_setter := new(StringCheckResultSetters)

	string_check_results_setter.IStringCheckServices = string_check_service

	return string_check_results_setter
}

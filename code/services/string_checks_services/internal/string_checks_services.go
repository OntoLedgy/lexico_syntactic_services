package internal

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters"
)

type StringChecksService struct {
	String_checks_parameter service_parameters.StringChecksParameters
	String_checks_result    service_results.StringChecksResults
}

func (
	string_checks_service *StringChecksService) Set_string_checks_result() {

	string_checks_result_setter :=
		string_checks_result_setters.
			Create(
				string_checks_service)

	string_checks_result_setter.
		Set_string_issues_and_fix()

}

func (
	string_checks_service *StringChecksService) Get_string_checks_result() service_results.StringChecksResults {

	return string_checks_service.String_checks_result
}

func (
	string_checks_service *StringChecksService) Get_string_checks_parameter() service_parameters.StringChecksParameters {

	return string_checks_service.String_checks_parameter
}

func (
	string_checks_service *StringChecksService) Set_issues_result(
	string_checks_issues []issues.Issues) {

	string_checks_service.String_checks_result.String_checks_issues = string_checks_issues

}

func (string_checks_service *StringChecksService) Set_string_fixes_result(
	string_checks_fix fixes.Fixes) {

	string_checks_service.String_checks_result.String_checks_fix = string_checks_fix

}

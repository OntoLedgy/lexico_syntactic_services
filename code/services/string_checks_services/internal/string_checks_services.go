package internal

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters"
)

type StringChecksService struct {
	String_checks_i_o_object *interservice_i_o_objects.StringChecksIOObjects
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
	string_checks_service *StringChecksService) Set_string_issues_result(
	string_checks_issue_result_list *service_results.IssueChecksResultLists) {

	string_checks_service.String_checks_i_o_object.String_checks_result.String_checks_issues_list = string_checks_issue_result_list

}

func (
	string_checks_service *StringChecksService) Get_string_checks_result() *service_results.StringChecksResults {

	return string_checks_service.String_checks_i_o_object.String_checks_result
}

func (
	string_checks_service *StringChecksService) Get_string_checks_input() *service_inputs.StringChecksInputs {

	return string_checks_service.String_checks_i_o_object.String_checks_input
}

func (
	string_checks_service *StringChecksService) Set_string_fixes_result(
	string_checks_fix service_results.FixChecksResults) {

	string_checks_service.String_checks_i_o_object.String_checks_result.String_checks_fix = string_checks_fix

}

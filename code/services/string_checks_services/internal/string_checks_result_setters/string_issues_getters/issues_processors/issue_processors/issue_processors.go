package issue_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_check_services"
)

type IssueCheckResultProcessors struct {
	issue_type   issues.IssueTypes
	string_value *identified_strings.Strings
}

func (
	issue_processor *IssueCheckResultProcessors) Get_string_issue_check_result(
	issue_type issues.IssueTypes) *service_results.IssueCheckResults {

	if issue_type.Issue_is_molecular == "FALSE" {

		string_check_issue_result :=
			issue_processor.
				get_string_atomic_issue_check_result()

		return string_check_issue_result
	}
	return nil
}

func (
	issue_processor *IssueCheckResultProcessors) get_string_atomic_issue_check_result() *service_results.IssueCheckResults {

	string_check_input :=
		new(
			service_inputs.
				StringCheckInputs)

	string_check_input.
		In_scope_issue_type =
		issue_processor.
			issue_type

	string_check_input.
		String_to_check =
		issue_processor.
			string_value

	String_check_service_factory :=
		new(
			string_check_services.
				StringCheckServiceFactory)

	string_check_service :=
		String_check_service_factory.
			Create(
				string_check_input)

	string_check_service.
		Set_string_check_result()

	issue_check_result_processor :=
		Create_issue_result_processor(
			issue_processor.
				string_value,
			&issue_processor.
				issue_type)

	string_check_issue_result :=
		issue_check_result_processor.
			process_issue_result(
				string_check_service)

	return string_check_issue_result
}

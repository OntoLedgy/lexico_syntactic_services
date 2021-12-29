package issue_processors

import (
	"github.com/OntoLedgy/string_editing_services/object_model"
	"github.com/OntoLedgy/syntactic_checker/code/infrastructure/logging"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_check_services/contract"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/internal/check_result_processors"
)

//TODO - split to separate type (issue_transaction_processor)

type IssueResultProcessors struct {
	string_value *identified_strings.Strings
	issue        *issues.Issues
	issue_type   *issues.IssueTypes
}

func (
	issue_transaction_processor *IssueResultProcessors) process_issue_result(
	string_check_service contract.IStringCheckServices) *service_results.IssueCheckResults {

	issues_found :=
		string_check_service.Get_string_check_result() != nil &&
			string_check_service.Get_string_check_result().Check_result_string_edit_ranges != nil

	if issues_found {

		logging.GlobalLogger.Printf(
			"\nprocessing issues results...\n")

		string_check_issue_result :=
			issue_transaction_processor.
				generate_issue_check_result(
					string_check_service)

		return string_check_issue_result

	}
	return nil
}

func (
	issue_transaction_processor *IssueResultProcessors) generate_issue_check_result(
	string_check_service contract.IStringCheckServices) *service_results.IssueCheckResults {

	string_check_issue :=
		new(
			issues.
				Issues)

	string_check_issue.
		Object_uuid =
		string_check_issue.
			Objects.
			Set_object_uuid()

	string_check_issue.Issue_type = *issue_transaction_processor.issue_type

	//TODO - Stage 3 - replace with factory

	string_edit_history :=
		new(
			object_model.StringEditHistories)

	string_edit_history.
		Create(
			issue_transaction_processor.
				string_value.String_value)

	check_result_processor :=
		check_result_processors.
			Create(
				string_edit_history,
				string_check_service.Get_string_check_result(),
				string_check_issue.Issue_type)

	check_result_processor.
		Process_regex_result()

	issue_check_result :=
		new(
			service_results.IssueCheckResults)

	issue_check_result.
		String_checks_issue =
		string_check_issue

	issue_check_result.
		String_edit_history =
		check_result_processor.
			String_edit_history

	return issue_check_result
}

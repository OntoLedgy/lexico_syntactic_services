package issues_processors

import (
	"fmt"
	"string_editor/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_check_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal/check_result_processors"
)

//TODO - split to separate type (issue_processor)

func (
	issues_processor *issuesProcessors) process_issue_transactions(
	string_check_service contract.IStringCheckServices) *issues.Issues {

	issues_found :=
		string_check_service.Get_string_check_result() != nil &&
			string_check_service.Get_string_check_result().Check_result_string_edit_ranges != nil

	if issues_found {

		fmt.Printf(
			"\nprocessing issues...\n")

		string_check_issue :=
			issues_processor.
				generate_issue_transaction(
					string_check_service)

		return string_check_issue

	}
	return nil
}

func (
	issues_processor *issuesProcessors) generate_issue_transaction(
	string_check_service contract.IStringCheckServices) *issues.Issues {

	string_check_issue :=
		new(
			issues.
				Issues)

	string_check_issue.
		Object_uuid =
		string_check_issue.
			Objects.
			Set_object_uuid()

	//TODO - Stage 3 - replace with factory

	string_edit_history :=
		new(
			object_model.StringEditHistories)

	string_edit_history.
		Create(
			issues_processor.
				string_checks_parameter.
				String_value)

	check_result_processor :=
		check_result_processors.Create(
			string_edit_history,
			string_check_service,
			string_check_issue.Issue_type)

	check_result_processor.
		Process_regex_result()

	string_check_issue.
		String_edit_history =
		*check_result_processor.
			String_edit_history

	return string_check_issue
}

package issues_processor

import (
	"fmt"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/string_check_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal/check_result_processors"
)

func (
	issues_processor *issuesProcessors) process_issue_transactions(
	identified_string_check_service contract.IStringCheckServices) *issues.Issues {

	issues_found :=
		identified_string_check_service.Get_check_result() != nil &&
			identified_string_check_service.Get_check_result().Check_result_string_edit_ranges != nil

	if issues_found {

		fmt.Printf(
			"\nprocessing issues...\n")

		identified_string_check_issue :=
			issues_processor.
				generate_issue_transaction(
					identified_string_check_service)

		return &identified_string_check_issue

	}
	return nil
}

func (
	issues_processor *issuesProcessors) generate_issue_transaction(
	string_check_service contract.IStringCheckServices) issues.Issues {

	string_check_issue :=
		new(
			issues.
				Issues)

	string_check_issue.
		Object_uuid =
		string_check_issue.
			Objects.
			Set_object_uuid()

	//TODO - Stage 3 - repalce with factory
	check_result_processor :=
		check_result_processors.Create(
			string_check_service,
			string_check_issue.Issue_type,
			issues_processor.identified_string,
			issues_processor.string_value)

	check_result_processor.
		Process_regex_result()

	string_check_issue.
		String_edit_history =
		*check_result_processor.
			String_edit_history

	return *string_check_issue
}

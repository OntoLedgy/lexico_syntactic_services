package fix_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
	"github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/internal/string_checks_result_setters/string_issues_getters/issues_processors/issue_processors"
)

type StringCheckProcessors struct {
	issue_type issues.IssueTypes
}

func (
	fixProcessor *FixProcessors) getStringCheckFix(
	issue_type issues.IssueTypes) {

	//TODO - Stage 3 - improve fix generation process (too wet).

	current_string := new(identified_strings.Strings)

	current_string.StringValue =
		fixProcessor.
			FixChecksResults.
			StringValueEditHistory.
			GetCurrentString()

	issue_check_result_processor := issue_processors.Create(
		current_string,
		issue_type)

	issue_check_result :=
		issue_check_result_processor.
			Get_string_issue_check_result(issue_type)

	//TODO - use factory?
	string_check_input :=
		new(
			service_inputs.
				StringCheckInputs)

	string_check_input.
		String_to_check = new(
		identified_strings.
			Strings)

	string_check_input.
		String_to_check.
		StringValue =
		fixProcessor.
			FixChecksResults.
			StringValueEditHistory.
			GetCurrentString()

	string_check_input.
		In_scope_issue_type =
		issue_type

	fixProcessor.
		update_string_edit_history(
			issue_check_result,
			issue_type)

	/*TODO - this looks like rouge code - Deprecate?
		string_check_service_factory :=
			new(
				string_check_services.
				StringCheckServiceFactory)

		string_check_service :=
			string_check_service_factory.
				Create(
					string_check_input)

		string_check_service.
			Set_string_check_result()
	//TODO - END of Rougue Code*/

}

func (
	fixProcessor *FixProcessors) update_string_edit_history(
	issue_check_result *service_results.IssueCheckResults,
	issue_type issues.IssueTypes) {

	string_check_result :=
		issue_check_result

	there_is_a_check_result :=
		string_check_result != nil

	if there_is_a_check_result {

		fixProcessor.
			FixChecksResults.
			StringValueEditHistory =
			string_check_result.
				String_edit_history

		if fixProcessor.FixChecksResults.Issue_check_result_list == nil {
			fixProcessor.FixChecksResults.Issue_check_result_list = new(service_results.IssueChecksResultLists)

		}

		fixProcessor.
			FixChecksResults.
			Issue_check_result_list.
			String_checks_issue_results =
			append(
				fixProcessor.
					FixChecksResults.
					Issue_check_result_list.
					String_checks_issue_results,
				string_check_result)

		fixProcessor.
			FixChecksResults.
			StringValueEditHistory.
			SetCurrentString(
				fixProcessor.
					FixChecksResults.StringValueEditHistory.
					GetModifiedString())

	}

}

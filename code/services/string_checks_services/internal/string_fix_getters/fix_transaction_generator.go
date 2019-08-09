package string_fix_getters

import (
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_check_services"
	"syntactic_checker/code/services/string_check_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal/check_result_processors"
)

func Generate_fix_transaction(
	string_value string,
	identified_string identified_strings.IdentifiedStrings,
	issue_types []issues.IssueTypes) fixes.Fixes {

	var string_check_fix fixes.Fixes

	string_edit_history :=
		new(
			string_editor_object_model.
				StringEditHistory)

	interim_identified_string_modified :=
		identified_string

	interim_string_modified :=
		string_value

	string_check_fix.
		Object_uuid =
		string_check_fix.
			Objects.
			Set_object_uuid()

	for _, issue_type := range issue_types {
		//TODO - Stage 3 - improve fix generation process (too wet).
		// should just process the fix object rather than the modified and marked strings.

		string_check_parameter :=
			new(
				service_parameters.StringCheckParameters)

		//TODO - string rename - deprecate

		string_check_parameter.
			Identified_string =
			interim_identified_string_modified
		//----

		string_check_parameter.
			String_value =
			interim_string_modified

		string_check_parameter.
			In_scope_issue_type =
			issue_type

		string_check_service_factory :=
			new(
				string_check_services.
					StringCheckServiceFactory)

		string_check_service :=
			string_check_service_factory.
				Create(
					*string_check_parameter)

		string_check_service.
			Set_string_check_result()

		interim_identified_string_modified, interim_string_modified, string_edit_history =
			update_modified_string_and_string_edit_history(
				issue_type,
				string_check_service,
				string_edit_history,
				interim_identified_string_modified,
				interim_string_modified)

	}

	string_check_fix.
		String_value_edit_history =
		*string_edit_history

	return string_check_fix
}

func update_modified_string_and_string_edit_history(
	issue_type issues.IssueTypes,
	string_check_service contract.IStringCheckServices,
	string_edit_history *string_editor_object_model.StringEditHistory,
	interim_identified_string_modified identified_strings.IdentifiedStrings,
	interim_string_modified string) (
	identified_strings.IdentifiedStrings,
	string,
	*string_editor_object_model.StringEditHistory) {

	string_check_result :=
		string_check_service.
			Get_check_result()

	there_is_a_check_result :=
		string_check_result.
			Check_result_string_edit_ranges != nil

	if there_is_a_check_result {

		check_result_processor :=
			check_result_processors.
				Create(
					string_check_service,
					issue_type,
					interim_identified_string_modified,
					interim_string_modified)

		check_result_processor.
			Process_regex_result()

		string_edit_history =
			check_result_processor.
				String_edit_history

		interim_identified_string_modified.
			String_value =
			string_edit_history.
				Get_modified_string()

	}

	return interim_identified_string_modified, interim_string_modified, string_edit_history
}

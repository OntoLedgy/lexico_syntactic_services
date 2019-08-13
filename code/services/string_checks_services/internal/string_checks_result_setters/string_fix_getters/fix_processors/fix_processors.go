package fix_processors

import (
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_check_services"
	"syntactic_checker/code/services/string_check_services/contract"
	"syntactic_checker/code/services/string_checks_services/internal/check_result_processors"
)

type FixProcessors struct {
	string_checks_parameter service_parameters.StringChecksParameters
	String_checks_fix       fixes.Fixes
}

func (
	fix_processor *FixProcessors) Set_string_checks_fix() {

	fix_processor.
		String_checks_fix =
		fix_processor.get_string_checks_fix()

}

func (
	fix_processor *FixProcessors) get_string_checks_fix() fixes.Fixes {

	var string_check_fix fixes.Fixes

	string_edit_history :=
		new(
			string_editor_object_model.
				StringEditHistories)

	string_edit_history.
		Create(fix_processor.
			string_checks_parameter.
			String_value)

	string_check_fix.
		Object_uuid =
		string_check_fix.
			Objects.
			Set_object_uuid()

	for _, issue_type := range fix_processor.string_checks_parameter.In_scope_issue_types {

		string_edit_history =
			fix_processor.
				get_string_check_fix(
					issue_type,
					string_edit_history)

	}

	string_check_fix.
		String_value_edit_history =
		*string_edit_history

	return string_check_fix
}

//TODO - split into a separate type (issue fix processor)?

func (
	fix_processor *FixProcessors) get_string_check_fix(
	issue_type issues.IssueTypes,
	string_edit_history *string_editor_object_model.StringEditHistories) *string_editor_object_model.StringEditHistories {

	//TODO - Stage 3 - improve fix generation process (too wet).
	// should just process the fix object rather than the modified and marked strings.

	string_check_parameter :=
		new(
			service_parameters.
				StringCheckParameters)

	string_check_parameter.
		String_value =
		string_edit_history.
			GetCurrentString()

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

	string_edit_history =
		fix_processor.
			update_string_edit_history(
				issue_type,
				string_check_service,
				string_edit_history)

	return string_edit_history
}

func (
	*FixProcessors) update_string_edit_history(
	issue_type issues.IssueTypes,
	string_check_service contract.IStringCheckServices,
	string_edit_history *string_editor_object_model.StringEditHistories) *string_editor_object_model.StringEditHistories {

	string_check_result :=
		string_check_service.
			Get_string_check_result()

	there_is_a_check_result :=
		string_check_result.
			Check_result_string_edit_ranges != nil

	if there_is_a_check_result {

		check_result_processor :=
			check_result_processors.
				Create(
					string_edit_history,
					string_check_service,
					issue_type)

		check_result_processor.
			Process_regex_result()

		string_edit_history =
			check_result_processor.
				String_edit_history

		string_edit_history.
			SetCurrentString(
				string_edit_history.
					Get_modified_string())

	}

	return string_edit_history
}

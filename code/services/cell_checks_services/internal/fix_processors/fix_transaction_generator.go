package fix_processors

import (
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_check_services"
	"syntactic_checker/code/services/cell_checks_services/internal/check_result_processors"
)

func Generate_fix_transaction(
	in_scope_cell cells.Cells,
	issue_types []issues.IssueTypes) fixes.Fixes {

	var cell_check_fix fixes.Fixes

	string_edit_history :=
		new(
			string_editor_object_model.
				StringEditHistory)

	interim_cell_modified :=
		in_scope_cell

	cell_check_fix.
		Object_uuid =
		cell_check_fix.
			Objects.
			Set_object_uuid()

	for _, issue_type := range issue_types {
		//TODO - Stage 3 - improve fix generation process (too wet).  should just process the fix object rather than the modified and marked cells.

		cell_check_parameter := new(service_parameters.CellCheckParameters)

		cell_check_parameter.
			In_scope_cell =
			interim_cell_modified

		cell_check_parameter.
			In_scope_issue_type =
			issue_type

		cell_check_service_factory :=
			new(
				cell_check_services.
					CellCheckServiceFactory)

		cell_check_service :=
			cell_check_service_factory.
				Create(
					*cell_check_parameter)

		cell_check_service.
			Set_cell_check_result()

		interim_cell_modified, string_edit_history =
			update_modified_cell_and_string_edit_history(
				issue_type,
				cell_check_service,
				string_edit_history,
				interim_cell_modified)

	}

	cell_check_fix.
		Cell_value_edit_history =
		*string_edit_history

	return cell_check_fix
}

func update_modified_cell_and_string_edit_history(
	issue_type issues.IssueTypes,
	cell_check_service cell_check_services.ICellCheckService,
	string_edit_history *string_editor_object_model.StringEditHistory,
	interim_cell_modified cells.Cells) (
	cells.Cells,
	*string_editor_object_model.StringEditHistory) {

	cell_check_result :=
		cell_check_service.
			Get_check_result()

	there_is_a_cell_check_result :=
		cell_check_result.Check_result_string_edit_ranges != nil

	if there_is_a_cell_check_result {

		check_result_processor :=
			new(
				check_result_processors.
					CheckResultProcessors)

		check_result_processor.
			Check_results =
			cell_check_service.
				Get_check_result()

		check_result_processor.
			In_scope_issue_type =
			issue_type

		check_result_processor.
			In_scope_cell =
			interim_cell_modified

		check_result_processor.
			Process_regex_result()

		string_edit_history =
			check_result_processor.
				Cell_edit_history

		interim_cell_modified.
			Cell_value =
			string_edit_history.
				Get_modified_string()

	}

	return interim_cell_modified, string_edit_history
}

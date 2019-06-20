package fixes_processor

import (
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_services"
)

func Generate_fix_transaction(
	in_scope_cell object_model.Cells,
	issue_types []issues.IssueTypes) fixes.Fixes {

	var cell_check_fix fixes.Fixes
	string_edit_history := new(string_editor_object_model.StringEditHistory)

	interim_cell_modified := in_scope_cell

	cell_check_fix.Object_uuid = cell_check_fix.Objects.Set_object_uuid()
	cell_check_fix.Cell = in_scope_cell

	for _, issue_type := range issue_types {
		//TODO - Stage 3 - improve fix generation process (too wet).  should just process the fix object rather than the modified and marked cells.
		cell_check_service :=
			cell_check_services.Create_cell_check_service(
				interim_cell_modified,
				issue_type).(*cell_check_services.CellCheckService)

		cell_check_service.
			Set_cell_check_results()

		interim_cell_modified, string_edit_history =
			update_modified_cell_and_string_edit_history(
				cell_check_service,
				string_edit_history,
				interim_cell_modified)

	}

	cell_check_fix.
		String_edit_history =
		string_edit_history

	return cell_check_fix
}

func update_modified_cell_and_string_edit_history(
	cell_check_service *cell_check_services.CellCheckService,
	string_edit_history *string_editor_object_model.StringEditHistory,
	interim_cell_modified object_model.Cells) (object_model.Cells, *string_editor_object_model.StringEditHistory) {

	if cell_check_service.Regex_check_result != nil {

		string_edit_history =
			&cell_check_service.
				Cell_value_edit_history

		interim_cell_modified.Cell_value =
			cell_check_service.
				Cell_value_edit_history.
				Get_modified_string()

	}

	return interim_cell_modified, string_edit_history
}

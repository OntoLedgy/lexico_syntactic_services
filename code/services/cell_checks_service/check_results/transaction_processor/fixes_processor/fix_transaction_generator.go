package fixes_processor

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_service"
	"syntactic_checker/code/services/cell_check_service/regex_checkers"
	"syntactic_checker/code/services/cell_checks_service/cell_editor"
)

func Generate_fix_transaction(
	in_scope_cell object_model.InScopeCells,
	issue_types []issues.IssueTypes) (object_model.InScopeCells, object_model.InScopeCells, fixes.Fixes) {

	var interim_cell_modified, interim_cell_marked object_model.InScopeCells

	interim_cell_modified = in_scope_cell
	interim_cell_marked = in_scope_cell

	for _, issue_type := range issue_types {

		_, cell_regex_check_result :=
			cell_check_service.
				Generate_cell_check_result(
					interim_cell_modified,
					issue_type)

		interim_cell_modified, interim_cell_marked =
			Generate_fixed_and_marked_cells(
				cell_regex_check_result,
				interim_cell_modified,
				issue_type,
				interim_cell_marked)

	}

	var cell_check_fix fixes.Fixes

	cell_check_fix.Cell = in_scope_cell
	cell_check_fix.Marked_string = interim_cell_marked.Cell_value
	cell_check_fix.Modified_string = interim_cell_modified.Cell_value

	return interim_cell_modified, interim_cell_marked, cell_check_fix
}

func Generate_fixed_and_marked_cells(
	cell_syntactic_check_issue_result *regex_checkers.RegexCheckResults,
	interim_cell_modified object_model.InScopeCells,
	issue_type issues.IssueTypes,
	interim_cell_marked object_model.InScopeCells) (object_model.InScopeCells, object_model.InScopeCells) {

	if cell_syntactic_check_issue_result != nil {

		interim_cell_modified.Cell_value, interim_cell_marked.Cell_value =
			cell_editor.Edit_cell(
				interim_cell_modified.Cell_value,
				issue_type,
				cell_syntactic_check_issue_result.Regex_match_indices)

	}

	return interim_cell_modified, interim_cell_marked
}

package fixes_processor

import (
	"syntactic_checker/cell_checkers"
	"syntactic_checker/cell_checkers/regex_checkers"
	"syntactic_checker/helpers/cell_editor"
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
)

func Generate_fix_transaction(
	in_scope_cell object_model.InScopeCell,
	issue_types []issues.IssueTypes) (object_model.InScopeCell, object_model.InScopeCell) {

	var interim_cell_modified, interim_cell_marked object_model.InScopeCell

	interim_cell_modified = in_scope_cell
	interim_cell_marked = in_scope_cell
	for _, issue_type := range issue_types {

		_, cell_regex_check_result :=
			cell_checkers.
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
	return interim_cell_modified, interim_cell_marked
}

func Generate_fixed_and_marked_cells(
	cell_syntactic_check_issue_result *regex_checkers.RegexCheckResults,
	interim_cell_modified object_model.InScopeCell,
	issue_type issues.IssueTypes,
	interim_cell_marked object_model.InScopeCell) (object_model.InScopeCell, object_model.InScopeCell) {

	if cell_syntactic_check_issue_result != nil {

		interim_cell_modified.Cell_value, interim_cell_marked.Cell_value =
			cell_editor.Edit_cell(
				interim_cell_modified.Cell_value,
				issue_type,
				cell_syntactic_check_issue_result.Regex_match_indices)

	}

	return interim_cell_modified, interim_cell_marked
}

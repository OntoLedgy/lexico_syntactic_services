package cell_editors

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/issues"
)

func Create(
	cell_to_edit cells.Cells,
	issue_type issues.IssueTypes,
	check_results *check_results.CheckResults) *cellEditor {

	cell_editor := new(cellEditor)

	cell_editor.cell_to_edit = cell_to_edit

	cell_editor.issue_type = issue_type

	cell_editor.check_results = check_results

	return cell_editor

}

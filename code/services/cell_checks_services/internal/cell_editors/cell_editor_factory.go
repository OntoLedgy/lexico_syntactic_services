package cell_editors

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/check_results"
)

func Create(
	cell_to_edit cells.Cells,
	check_results *check_results.CheckResults,
	replacement_string string) *cellEditor {

	cell_editor :=
		new(
			cellEditor)

	cell_editor.
		cell_to_edit =
		cell_to_edit

	cell_editor.
		replacement_string =
		replacement_string

	cell_editor.
		check_results =
		check_results

	return cell_editor

}

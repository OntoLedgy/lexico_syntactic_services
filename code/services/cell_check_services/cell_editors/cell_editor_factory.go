package cell_editors

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_services/regex_checkers"
)

func Create_cell_editor(
	cell_to_edit object_model.Cells,
	issue_type issues.IssueTypes,
	regex_result regex_checkers.RegexCheckResults) *cellEditor {

	cell_editor := new(cellEditor)

	cell_editor.cell_to_edit = cell_to_edit

	cell_editor.issue_type = issue_type

	cell_editor.regex_result = regex_result

	return cell_editor

}

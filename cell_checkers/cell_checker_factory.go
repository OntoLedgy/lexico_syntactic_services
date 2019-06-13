package cell_checkers

import (
	"syntactic_checker/cell_checkers/regex_checkers"
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
)

func Create_cell_checker(
	cell object_model.InScopeCell,
	in_scope_syntactic_check_type issues.IssueTypes) ICellCheckers {

	var cell_checker CellCheckers

	cell_checker.Cell = cell
	cell_checker.Issue_type = in_scope_syntactic_check_type

	return cell_checker

}

func Generate_cell_check_result(
	in_scope_cell object_model.InScopeCell,
	in_scope_syntactic_check_type issues.IssueTypes) (ICellCheckers, *regex_checkers.RegexCheckResults) {

	cell_checker := Create_cell_checker(
		in_scope_cell,
		in_scope_syntactic_check_type)

	cell_syntactic_check_result :=
		cell_checker.
			CheckCell()

	return cell_checker, cell_syntactic_check_result
}

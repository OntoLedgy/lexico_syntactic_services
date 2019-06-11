package cell_checkers

import "syntactic_checker/object_model"

func Create_cell_checker(
	cell object_model.InScopeCell,
	in_scope_syntactic_check_type object_model.IssueTypes) ICellCheckers {

	var cell_checker CellCheckers

	cell_checker.Cell = cell
	cell_checker.Issue_type = in_scope_syntactic_check_type

	return cell_checker

}

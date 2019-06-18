package cell_check_service

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_service/regex_checkers"
)

func Create_cell_checker(
	cell object_model.InScopeCells,
	in_scope_syntactic_check_type issues.IssueTypes) ICellCheckers {

	var cell_checker CellCheckers

	//TODO should be a constructor (internal)
	cell_checker.Cell = cell
	cell_checker.Issue_type = in_scope_syntactic_check_type

	return cell_checker

}

func Generate_cell_check_result(
	in_scope_cell object_model.InScopeCells,
	in_scope_syntactic_check_type issues.IssueTypes) (ICellCheckers, *regex_checkers.RegexCheckResults) {

	cell_checker := Create_cell_checker(
		in_scope_cell,
		in_scope_syntactic_check_type)

	cell_syntactic_check_result :=
		cell_checker.
			CheckCell()

	return cell_checker, cell_syntactic_check_result
}

package cell_check_services

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
)

func Create_cell_check_service(
	cell object_model.Cells,
	in_scope_syntactic_check_type issues.IssueTypes) ICellCheckService {

	cell_checker := new(
		CellCheckService)

	//TODO should be a constructor (internal)
	cell_checker.In_scope_cell = cell
	cell_checker.Issue_type = in_scope_syntactic_check_type

	return cell_checker

}

package cell_list_checks_services

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
)

func Create_cell_list_checks_service(
	in_scope_cells object_model.ListOfCells,
	issue_types []issues.IssueTypes) iCellListChecksService {

	cell_list_checks_service :=
		new(
			cellListChecksService)

	cell_list_checks_service.
		In_scope_cells =
		in_scope_cells

	cell_list_checks_service.
		Issue_types =
		issue_types

	return cell_list_checks_service
}

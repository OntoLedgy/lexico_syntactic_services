package cell_checks_services

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/transactions"
)

type CellChecksService struct {
	In_scope_cell object_model.Cells
	Issue_types   []issues.IssueTypes
	cell_fix      fixes.Fixes
	cell_issues   transactions.IssuesTransactions
}

func (cell_checks_service *CellChecksService) Set_cell_checks_results() {

	cell_checks_service.
		set_cell_checks_issues_and_fixes()

}

func (cell_checks_service *CellChecksService) Get_cell_checks_issues() transactions.IssuesTransactions {

	return cell_checks_service.cell_issues
}

func (cell_checks_service *CellChecksService) Get_cell_checks_fix() fixes.Fixes {

	return cell_checks_service.cell_fix
}

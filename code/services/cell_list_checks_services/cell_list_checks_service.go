package cell_list_checks_services

import (
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/transactions"
)

type cellListChecksService struct {
	Issue_types                   []issues.IssueTypes
	In_scope_cells                object_model.ListOfCells
	Cell_list_issues_transactions transactions.IssuesTransactions
	Cell_list_fix_transactions    transactions.FixTransactions
}

func (
	cell_list_checks_service *cellListChecksService) Set_syntactic_checks_results() {

	fmt.Println("\nPreparing report..")

	cell_list_checks_service.
		iterate_cell_list_for_cell_checks()

}

func (
	cell_list_checks_service *cellListChecksService) Get_fix_transactions() transactions.FixTransactions {

	return cell_list_checks_service.Cell_list_fix_transactions
}

func (
	cell_list_checks_service *cellListChecksService) Get_issue_transactions() transactions.IssuesTransactions {

	return cell_list_checks_service.Cell_list_issues_transactions
}

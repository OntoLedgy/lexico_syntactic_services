package cell_checks_service

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/cell_checks_service/check_results/transaction_processor/fixes_processor"
	"syntactic_checker/code/services/cell_checks_service/check_results/transaction_processor/issues_processor"
)

type CellChecksService struct {
	In_scope_cell object_model.InScopeCells
	Issue_types   []issues.IssueTypes
	cell_fix      fixes.Fixes
	cell_issues   transactions.IssueTransactions
}

func (
	cell_checks_service *CellChecksService) Get_cell_checks_results() (
	[]interface{}, fixes.Fixes, transactions.IssueTransactions) {

	cell_check_fixes_transaction :=
		cell_checks_service.
			get_cell_check_issues_and_fixes()

	return cell_check_fixes_transaction, cell_checks_service.cell_fix, cell_checks_service.cell_issues

}

func (
	cell_checks_service *CellChecksService) get_cell_check_issues_and_fixes() []interface{} {

	cell_check_issues_transactions, cell_check_issues :=
		issues_processor.
			Get_cell_check_issues(
				cell_checks_service.Issue_types,
				cell_checks_service.In_scope_cell)

	cell_check_fixes_transaction, cell_checks_fix :=
		fixes_processor.
			Get_cell_check_fix(
				cell_checks_service.Issue_types,
				cell_checks_service.In_scope_cell,
				cell_check_issues_transactions)

	cell_checks_service.cell_issues = cell_check_issues
	cell_checks_service.cell_fix = cell_checks_fix

	return cell_check_fixes_transaction
}

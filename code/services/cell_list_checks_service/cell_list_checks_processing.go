package cell_list_checks_service

import (
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/cell_checks_service"
)

func (cell_list_checks_service *cellListChecksService) iterate_on_cell_list_for_cell_checks() ([][]interface{}, []fixes.Fixes, transactions.IssueTransactions) {

	//TODO - Deprecate slice approach and use objects for fixes and issues
	var cell_syntactic_check_fix_transaction []interface{}
	var cells_syntactic_check_fix_transactions [][]interface{}
	//TODO - Deprecate slice approach and use objects for fixes and issues

	var cell_check_fix fixes.Fixes
	var column_check_fixes_transaction []fixes.Fixes
	var cell_check_issues transactions.IssueTransactions

	in_scope_cells := cell_list_checks_service.In_scope_cells
	issue_types := cell_list_checks_service.Issue_types

	fmt.Printf(
		"processing checks: %s\n",
		issue_types)

	for _, in_scope_cell := range in_scope_cells.Cells {

		column_check_fixes_transaction, cell_list_checks_service.cell_list_check_issues, cells_syntactic_check_fix_transactions =
			cell_list_checks_service.process_cell_checks(
				in_scope_cell,
				cell_syntactic_check_fix_transaction,
				cell_check_fix,
				cell_check_issues,
				column_check_fixes_transaction,
				cell_list_checks_service.cell_list_check_issues,
				cells_syntactic_check_fix_transactions)

	}

	return cells_syntactic_check_fix_transactions, column_check_fixes_transaction, cell_list_checks_service.cell_list_check_issues
}

func (cell_list_checks_service *cellListChecksService) process_cell_checks(
	in_scope_cell object_model.InScopeCells,
	cell_syntactic_check_fix_transaction []interface{},
	cell_check_fix fixes.Fixes,
	cell_check_issues transactions.IssueTransactions,
	cell_list_fixes []fixes.Fixes,
	cell_list_check_issues transactions.IssueTransactions,
	cells_syntactic_check_fix_transactions [][]interface{}) ([]fixes.Fixes, transactions.IssueTransactions, [][]interface{}) {

	cell_checks_service :=
		cell_checks_service.Create_cell_checks_service(
			in_scope_cell,
			cell_list_checks_service.Issue_types)

	cell_syntactic_check_fix_transaction, cell_check_fix, cell_check_issues =
		cell_checks_service.
			Get_cell_checks_results()

	if cell_check_issues.Issues != nil {

		cell_list_fixes =
			append(
				cell_list_fixes,
				cell_check_fix)

		cell_list_check_issues.Issues =
			append(
				cell_list_check_issues.Issues,
				cell_check_issues.Issues...)
	}
	cells_syntactic_check_fix_transactions =
		process_syntactic_check_fix_transactions(
			cell_syntactic_check_fix_transaction,
			cells_syntactic_check_fix_transactions)

	return cell_list_fixes, cell_list_check_issues, cells_syntactic_check_fix_transactions
}

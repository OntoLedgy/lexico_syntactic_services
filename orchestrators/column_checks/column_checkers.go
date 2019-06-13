package column_checks

import (
	"fmt"
	"syntactic_checker/check_results_processor/finaliser"
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
	"syntactic_checker/orchestrators/cell_checks"
)

type ColumnCheckers struct {
	Issue_types    []issues.IssueTypes
	In_scope_cells object_model.InScopeCells
}

func (
	cells_checks_orchestrator *ColumnCheckers) RunColumnChecks() map[string][][]string {

	var syntactic_check_result_report map[string][][]string

	cells_syntactic_check_issues_transactions, cells_syntactic_check_fix_transactions :=
		process_syntactic_checks_for_cells(
			cells_checks_orchestrator.In_scope_cells,
			cells_checks_orchestrator.Issue_types)

	fmt.Println("\nPreparing report..")

	syntactic_check_result_report =
		finaliser.Prepare_syntactic_checks_results_transactions(
			cells_syntactic_check_issues_transactions,
			cells_syntactic_check_fix_transactions)

	return syntactic_check_result_report

}

func process_syntactic_checks_for_cells(
	in_scope_cells object_model.InScopeCells,
	issue_types []issues.IssueTypes) ([][]interface{}, [][]interface{}) {

	var cells_syntactic_check_issues_transactions [][]interface{}
	var cell_syntactic_check_issues_transactions [][]interface{}
	var cell_syntactic_check_fix_transaction []interface{}
	var cells_syntactic_check_fix_transactions [][]interface{}

	fmt.Printf(
		"processing checks: %s\n",
		issue_types)

	for _, in_scope_cell := range in_scope_cells.Cells {

		cell_checks_orchestrator :=
			cell_checks.Create_cell_checks_orchestrator(
				in_scope_cell,
				issue_types)

		cell_syntactic_check_issues_transactions, cell_syntactic_check_fix_transaction =
			cell_checks_orchestrator.
				RunCellChecks()

		cells_syntactic_check_issues_transactions, cells_syntactic_check_fix_transactions =
			process_syntactic_check_results_transactions(
				cells_syntactic_check_issues_transactions,
				cell_syntactic_check_issues_transactions,
				cell_syntactic_check_fix_transaction,
				cells_syntactic_check_fix_transactions)

	}

	/*
		if cells_syntactic_check_issues_transactions != nil { //#TODO add to logger
			fmt.Printf(
				"\n++++++++++++++++++++++++IssueTypes completed: %s issues found.\nExample transaction row: %s ",
				len(cells_syntactic_check_issues_transactions),
				cells_syntactic_check_issues_transactions[0])
		}*/
	return cells_syntactic_check_issues_transactions, cells_syntactic_check_fix_transactions
}

func process_syntactic_check_results_transactions(
	cells_syntactic_check_issues_transactions [][]interface{},
	cell_syntactic_check_issues_transactions [][]interface{},
	cell_syntactic_check_fix_transaction []interface{},
	cells_syntactic_check_fix_transactions [][]interface{}) ([][]interface{}, [][]interface{}) {

	//TODO - Stage 2 - move to separate function (Record Issue Transactions)
	cells_syntactic_check_issues_transactions =
		append(
			cells_syntactic_check_issues_transactions,
			cell_syntactic_check_issues_transactions...)
	//TODO - Stage 2 - move to separate function (Record Fix Transactions)
	if cell_syntactic_check_fix_transaction != nil {

		cells_syntactic_check_fix_transactions =
			append(
				cells_syntactic_check_fix_transactions,
				cell_syntactic_check_fix_transaction)

		cell_syntactic_check_fix_transaction = nil

	}
	return cells_syntactic_check_issues_transactions, cells_syntactic_check_fix_transactions
}

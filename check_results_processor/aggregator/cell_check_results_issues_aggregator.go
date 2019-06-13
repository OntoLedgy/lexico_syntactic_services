package aggregator

import (
	"fmt"
	"syntactic_checker/cell_checkers"
)

func Aggregate_syntactic_check_issues_transactions(
	cell_checker cell_checkers.CellCheckers,
	cell_syntactic_check_issue_transaction []interface{},
	cell_syntactic_check_issue_transactions [][]interface{}) [][]interface{} {

	if cell_syntactic_check_issue_transaction != nil {

		cell_syntactic_check_issue_transaction =
			append(
				cell_syntactic_check_issue_transaction,
				cell_checker.Issue_type.Issue_type_uuid,
				cell_checker.Cell.Cell_identifier)

		cell_syntactic_check_issue_transactions =
			append(cell_syntactic_check_issue_transactions,
				cell_syntactic_check_issue_transaction)

		fmt.Printf(
			"\n\ncell_syntactic_check_issue_transactions:%v\n\n\n",
			cell_syntactic_check_issue_transactions)
	}

	return cell_syntactic_check_issue_transactions
}

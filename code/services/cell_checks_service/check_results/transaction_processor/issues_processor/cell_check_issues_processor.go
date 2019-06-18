package issues_processor

import (
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/cell_check_service"
	"syntactic_checker/code/services/cell_check_service/regex_checkers"
	"syntactic_checker/code/services/cell_checks_service/check_results/aggregator"
)

func Get_cell_check_issues(
	in_scope_syntactic_check_types []issues.IssueTypes,
	in_scope_cell object_model.InScopeCells) ([][]interface{}, transactions.IssueTransactions) {

	var cell_syntactic_check_issue_transactions [][]interface{}
	var cell_check_issue issues.Issues
	var cell_checks_issues_transaction transactions.IssueTransactions

	for _, in_scope_syntactic_check_type := range in_scope_syntactic_check_types {

		cell_checker, cell_regex_check_result :=
			cell_check_service.
				Generate_cell_check_result(
					in_scope_cell,
					in_scope_syntactic_check_type)

		cell_syntactic_check_issue_transactions, cell_check_issue =
			Process_cell_check_issue_transactions(
				cell_checker.(cell_check_service.CellCheckers),
				cell_regex_check_result,
				cell_syntactic_check_issue_transactions)

		if cell_check_issue.Modified_cell_value != "" {
			cell_checks_issues_transaction.Issues =
				append(
					cell_checks_issues_transaction.Issues,
					cell_check_issue)
		}
	}

	return cell_syntactic_check_issue_transactions, cell_checks_issues_transaction

}

func Process_cell_check_issue_transactions(
	cell_checker cell_check_service.CellCheckers,
	regex_check_result *regex_checkers.RegexCheckResults,
	cell_syntactic_check_issue_transactions [][]interface{}) ([][]interface{}, issues.Issues) {

	var cell_syntactic_check_issue_transaction []interface{}

	var cell_check_issue issues.Issues

	if regex_check_result != nil {

		fmt.Printf(
			"\nprocessing issues...\n")

		cell_syntactic_check_issue_transaction, cell_check_issue =
			Generate_issue_transaction(
				*regex_check_result,
				cell_checker.Issue_type)

		cell_check_issue.Issue_type = cell_checker.Issue_type
		cell_check_issue.Cell = cell_checker.Cell

		cell_syntactic_check_issue_transactions =
			aggregator.Aggregate_syntactic_check_issues_transactions(
				cell_checker,
				cell_syntactic_check_issue_transaction,
				cell_syntactic_check_issue_transactions)

	}

	return cell_syntactic_check_issue_transactions, cell_check_issue
}

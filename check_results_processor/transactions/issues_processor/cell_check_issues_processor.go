package issues_processor

import (
	"fmt"
	"syntactic_checker/cell_checkers"
	"syntactic_checker/cell_checkers/regex_checkers"
	"syntactic_checker/check_results_processor/aggregator"
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
)

func Get_cell_check_issue_transactions(
	in_scope_syntactic_check_types []issues.IssueTypes,
	in_scope_cell object_model.InScopeCell) [][]interface{} {

	var cell_syntactic_check_issue_transactions [][]interface{}

	for _, in_scope_syntactic_check_type := range in_scope_syntactic_check_types {

		cell_checker, cell_regex_check_result :=
			cell_checkers.
				Generate_cell_check_result(
					in_scope_cell,
					in_scope_syntactic_check_type)

		cell_syntactic_check_issue_transactions =
			Process_cell_check_issue_transactions(
				cell_checker.(cell_checkers.CellCheckers),
				cell_regex_check_result,
				cell_syntactic_check_issue_transactions)
	}

	return cell_syntactic_check_issue_transactions

}

func Process_cell_check_issue_transactions(
	cell_checker cell_checkers.CellCheckers,
	regex_check_result *regex_checkers.RegexCheckResults,
	cell_syntactic_check_issue_transactions [][]interface{}) [][]interface{} {

	var cell_syntactic_check_issue_transaction []interface{}

	if regex_check_result != nil {

		fmt.Printf(
			"\nprocessing issues...\n")

		cell_syntactic_check_issue_transaction =
			Generate_issue_transaction(
				*regex_check_result,
				cell_checker.Issue_type)

		cell_syntactic_check_issue_transactions =
			aggregator.Aggregate_syntactic_check_issues_transactions(
				cell_checker,
				cell_syntactic_check_issue_transaction,
				cell_syntactic_check_issue_transactions)

	}

	return cell_syntactic_check_issue_transactions
}

package cell_checks

import (
	"syntactic_checker/cell_checkers"
	"syntactic_checker/check_results/aggregator"
	"syntactic_checker/object_model"
)

func Process_syntactic_check_issues_for_cell(
	in_scope_syntactic_check_types []object_model.IssueTypes,
	cell_syntactic_check_issue_transactions [][]interface{},
	in_scope_cell object_model.InScopeCell) [][]interface{} {

	for _, in_scope_syntactic_check_type := range in_scope_syntactic_check_types {

		cell_checker := cell_checkers.
			Create_cell_checker(
				in_scope_cell,
				in_scope_syntactic_check_type)

		cell_syntactic_check_issue_result :=
			cell_checker.
				CheckCell()

		cell_syntactic_check_issue_transactions =
			aggregator.Aggregate_syntactic_check_issues(
				cell_checker.(cell_checkers.CellCheckers),
				cell_syntactic_check_issue_result,
				cell_syntactic_check_issue_transactions)
	}

	return cell_syntactic_check_issue_transactions

}

package cell_checks

import (
	"syntactic_checker/object_model"
)

type CellChecksOrchestrators struct {
	In_scope_cell object_model.InScopeCell
	Issue_types   []object_model.IssueTypes
}

func (
	cell_checks_orchestrator CellChecksOrchestrators) RunCellChecks() ([][]interface{}, []interface{}) {

	cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction :=
		process_syntactic_checks_for_cell(
			cell_checks_orchestrator.In_scope_cell,
			cell_checks_orchestrator.Issue_types)

	return cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction

}

func process_syntactic_checks_for_cell(
	in_scope_cell object_model.InScopeCell,
	in_scope_syntactic_check_types []object_model.IssueTypes) ([][]interface{}, []interface{}) {

	var cell_syntactic_check_issue_transactions [][]interface{}
	var cell_syntactic_check_aggregated_fixes_transaction []interface{}
	//var cell_syntactic_check_issue_results []regex_checkers.RegexCheckResults

	cell_syntactic_check_issue_transactions =
		Process_syntactic_check_issues_for_cell(
			in_scope_syntactic_check_types,
			cell_syntactic_check_issue_transactions,
			in_scope_cell)

	cell_syntactic_check_aggregated_fixes_transaction =
		Process_syntactic_check_fixes_for_cell(
			in_scope_syntactic_check_types,
			in_scope_cell,
			cell_syntactic_check_issue_transactions)

	return cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction
}

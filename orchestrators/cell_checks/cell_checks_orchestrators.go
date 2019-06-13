package cell_checks

import (
	"syntactic_checker/check_results_processor/transactions/fixes_processor"
	"syntactic_checker/check_results_processor/transactions/issues_processor"
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
)

type CellChecksOrchestrators struct {
	In_scope_cell object_model.InScopeCell
	Issue_types   []issues.IssueTypes
}

func (
	cell_checks_orchestrator CellChecksOrchestrators) RunCellChecks() ([][]interface{}, []interface{}) {

	cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction :=
		get_cell_check_issues_and_fixes(
			cell_checks_orchestrator.In_scope_cell,
			cell_checks_orchestrator.Issue_types)

	return cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction

}

func get_cell_check_issues_and_fixes(
	in_scope_cell object_model.InScopeCell,
	in_scope_syntactic_check_types []issues.IssueTypes) ([][]interface{}, []interface{}) {

	cell_syntactic_check_issue_transactions :=
		issues_processor.
			Get_cell_check_issue_transactions(
				in_scope_syntactic_check_types,
				in_scope_cell)

	cell_syntactic_check_aggregated_fixes_transaction :=
		fixes_processor.
			Process_cell_check_fixes(
				in_scope_syntactic_check_types,
				in_scope_cell,
				cell_syntactic_check_issue_transactions)

	return cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction
}

package cell_list_checks_service

import (
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/syntactic_check_service/results_processor"
)

type cellListChecksService struct {
	Issue_types            []issues.IssueTypes
	In_scope_cells         object_model.ListOfInScopeCells
	cell_list_check_issues transactions.IssueTransactions
}

func (
	cell_list_checks_service *cellListChecksService) Get_syntactic_check_results() map[string][][]string {

	var syntactic_check_result_report map[string][][]string

	fmt.Println("\nPreparing report..")

	cells_syntactic_check_fix_transactions, column_check_fixes, column_check_issues :=
		cell_list_checks_service.
			iterate_on_cell_list_for_cell_checks()

	syntactic_check_result_report =
		results_processor.Prepare_syntactic_checks_results_transactions(
			cells_syntactic_check_fix_transactions,
			column_check_issues,
			column_check_fixes)

	return syntactic_check_result_report

}

package cell_list_checks_services

import (
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/services/cell_checks_services"
)

func (cell_list_checks_service *cellListChecksService) iterate_cell_list_for_cell_checks() {

	fmt.Printf(
		"processing checks: %s\n",
		cell_list_checks_service.Issue_types)

	in_scope_cells :=
		cell_list_checks_service.
			In_scope_cells.
			Cells

	for _, in_scope_cell := range in_scope_cells {

		cell_list_checks_service.
			process_cell_checks(
				in_scope_cell)
	}

}

func (
	cell_list_checks_service *cellListChecksService) process_cell_checks(
	in_scope_cell object_model.Cells) {

	in_scope_issue_types :=
		cell_list_checks_service.
			Issue_types

	cell_checks_service :=
		cell_checks_services.Create_cell_checks_service(
			in_scope_cell,
			in_scope_issue_types)

	cell_checks_service.
		Set_cell_checks_results()

	cell_list_checks_service.
		append_check_results(
			cell_checks_service)

}

func (
	cell_list_checks_service *cellListChecksService) append_check_results(
	cell_checks_service cell_checks_services.ICellCheckOrchestrators) {

	cell_check_issues := cell_checks_service.Get_cell_checks_issues()

	if cell_check_issues.Issues != nil {

		cell_check_fixes := cell_checks_service.Get_cell_checks_fix()

		cell_list_checks_service.Cell_list_fix_transactions.Fixes =
			append(
				cell_list_checks_service.Cell_list_fix_transactions.Fixes,
				cell_check_fixes)

		cell_list_checks_service.Cell_list_issues_transactions.Issues =
			append(
				cell_list_checks_service.Cell_list_issues_transactions.Issues,
				cell_check_issues.Issues...)
	}
}

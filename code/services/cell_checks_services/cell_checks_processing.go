package cell_checks_services

import (
	"syntactic_checker/code/services/cell_checks_services/transaction_processor/fixes_processor"
	"syntactic_checker/code/services/cell_checks_services/transaction_processor/issues_processor"
)

func (
	cell_checks_service *CellChecksService) set_cell_checks_issues_and_fixes() {

	issues_processor :=
		issues_processor.
			Create_issues_processor(
				cell_checks_service.In_scope_cell,
				cell_checks_service.Issue_types)

	issues_processor.
		Set_cell_check_issues()

	cell_check_issues :=
		issues_processor.
			Get_cell_check_issue_transactions()

	cell_checks_fix :=
		fixes_processor.
			Get_cell_check_fix(
				cell_checks_service.Issue_types,
				cell_checks_service.In_scope_cell,
				cell_check_issues.Issues)

	if cell_checks_fix != nil {

		cell_checks_service.
			cell_issues =
			cell_check_issues

		cell_checks_service.
			cell_fix =
			*cell_checks_fix

	}

}

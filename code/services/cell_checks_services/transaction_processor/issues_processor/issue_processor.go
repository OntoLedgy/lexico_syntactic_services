package issues_processor

import (
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_services"
)

func (
	issues_processor *issuesProcessor) get_cell_check_issue(
	in_scope_issue_type issues.IssueTypes) *issues.Issues {

	in_scope_cell :=
		issues_processor.
			in_scope_cell

	cell_check_service :=
		cell_check_services.
			Create_cell_check_service(
				in_scope_cell,
				in_scope_issue_type).(*cell_check_services.CellCheckService)

	cell_check_service.
		Set_cell_check_results()

	cell_check_issue :=
		process_issue_transactions(
			cell_check_service)

	return cell_check_issue
}

func (issues_processor *issuesProcessor) append_cell_check_issue(
	cell_check_issue *issues.Issues) {

	if cell_check_issue != nil {

		issues_processor.cell_check_issues_transaction.Issues =
			append(
				issues_processor.cell_check_issues_transaction.Issues,
				*cell_check_issue)
	}

}

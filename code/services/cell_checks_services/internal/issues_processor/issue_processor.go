package issues_processor

import (
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_check_services"
)

func (
	issues_processor *issuesProcessors) get_cell_check_issue(
	in_scope_issue_type issues.IssueTypes) *issues.Issues {

	in_scope_cell :=
		issues_processor.
			in_scope_cell

	cell_check_parameter := new(
		service_parameters.
			CellCheckParameters)

	cell_check_parameter.
		In_scope_cell =
		in_scope_cell

	cell_check_parameter.
		In_scope_issue_type =
		in_scope_issue_type

	cell_check_service_factory :=
		new(
			cell_check_services.
				CellCheckServiceFactory)

	cell_check_service :=
		cell_check_service_factory.
			Create(
				*cell_check_parameter)

	cell_check_service.
		Set_cell_check_result()

	cell_check_issue :=
		issues_processor.
			process_issue_transactions(
				cell_check_service)

	return cell_check_issue
}

func (
	issues_processor *issuesProcessors) process_cell_check_issue(
	cell_check_issue *issues.Issues,
	in_scope_issue_type issues.IssueTypes) {

	there_is_an_issue :=
		cell_check_issue != nil

	if there_is_an_issue {

		cell_check_issue.
			Issue_type =
			in_scope_issue_type

		issues_processor.
			append_cell_check_issue(
				cell_check_issue)
	}
}

func (
	issues_processor *issuesProcessors) append_cell_check_issue(
	cell_check_issue *issues.Issues) {

	if cell_check_issue != nil {

		issues_processor.
			cell_checks_issues = append(
			issues_processor.cell_checks_issues,
			*cell_check_issue)

	}

}

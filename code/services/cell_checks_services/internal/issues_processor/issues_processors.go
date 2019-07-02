package issues_processor

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
)

type issuesProcessors struct {
	in_scope_issue_types []issues.IssueTypes
	in_scope_cell        cells.Cells
	cell_checks_issues   []issues.Issues
}

func (
	issues_processor *issuesProcessors) Set_cell_check_issues() {

	in_scope_issue_types :=
		issues_processor.
			in_scope_issue_types

	for _, in_scope_issue_type := range in_scope_issue_types {

		cell_check_issue :=
			issues_processor.
				get_cell_check_issue(
					in_scope_issue_type)

		issues_processor.
			process_cell_check_issue(
				cell_check_issue,
				in_scope_issue_type)
	}

}

func (
	issues_processor *issuesProcessors) Get_cell_checks_issues() []issues.Issues {

	return issues_processor.cell_checks_issues

}

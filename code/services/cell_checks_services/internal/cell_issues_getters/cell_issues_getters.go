package cell_issues_getters

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_checks_services/internal/issues_processor"
)

//push this to internal types
//split 1: check cell for issues

type CellIssuesGetters struct {
}

func Get_cell_issues(
	in_scope_cell cells.Cells,
	list_of_in_scope_issue_types []issues.IssueTypes) []issues.Issues {

	issues_processor :=
		issues_processor.
			Create(
				in_scope_cell,
				list_of_in_scope_issue_types)

	issues_processor.
		Set_cell_check_issues()

	cell_checks_issues :=
		issues_processor.
			Get_cell_checks_issues()

	return cell_checks_issues
}

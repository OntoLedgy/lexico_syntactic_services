package issues_processor

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
)

func Create(
	in_scope_cell cells.Cells,
	in_scope_issue_types []issues.IssueTypes) *issuesProcessors {

	issues_processor :=
		new(
			issuesProcessors)

	issues_processor.
		in_scope_cell =
		in_scope_cell

	issues_processor.
		in_scope_issue_types =
		in_scope_issue_types

	return issues_processor

}

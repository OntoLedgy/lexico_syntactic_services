package issues_processor

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
)

func Create_issues_processor(
	in_scope_cell object_model.Cells,
	in_scope_issue_types []issues.IssueTypes) *issuesProcessor {

	issues_processor := new(issuesProcessor)

	issues_processor.in_scope_cell = in_scope_cell
	issues_processor.in_scope_issue_types = in_scope_issue_types

	return issues_processor

}

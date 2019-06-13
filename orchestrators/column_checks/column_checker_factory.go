package column_checks

import (
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
)

func Create_column_checker(
	in_scope_cells object_model.InScopeCells,
	issue_types []issues.IssueTypes) IColumnCheckers {

	column_checker_orchestrator :=
		new(
			ColumnCheckers)

	column_checker_orchestrator.In_scope_cells = in_scope_cells
	column_checker_orchestrator.Issue_types = issue_types

	return column_checker_orchestrator
}

package cell_checks

import (
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
)

func Create_cell_checks_orchestrator(
	cell_value object_model.InScopeCell,
	check_types []issues.IssueTypes) ICellCheckOrchestrators {

	cell_check_orchestrator :=
		new(
			CellChecksOrchestrators)

	cell_check_orchestrator.Issue_types = check_types
	cell_check_orchestrator.In_scope_cell = cell_value

	return cell_check_orchestrator

}

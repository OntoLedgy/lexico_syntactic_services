package cell_checks_service

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
)

func Create_cell_checks_service(
	cell_value object_model.InScopeCells,
	check_types []issues.IssueTypes) ICellCheckOrchestrators {

	cell_check_orchestrator :=
		new(
			CellChecksService)

	cell_check_orchestrator.Issue_types = check_types
	cell_check_orchestrator.In_scope_cell = cell_value

	return cell_check_orchestrator

}

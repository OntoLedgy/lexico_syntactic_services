package column_checks

import (
	"syntactic_checker/helpers/configuration_handler"
	"syntactic_checker/object_model"
)

func Create_column_checker(
	in_scope_cells object_model.InScopeCells,
	run_configuration *configuration_handler.Configurations) IColumnCheckers {

	column_checker_orchestrator :=
		new(
			ColumnCheckers)

	column_checker_orchestrator.Run_configuration = run_configuration
	column_checker_orchestrator.In_scope_cells = in_scope_cells

	return column_checker_orchestrator
}

package cell_checks_orchestrator

import (
	"syntactic_checker/cell_checks_results_finaliser"
	"syntactic_checker/cell_checks_runner"
	"syntactic_checker/cells_preparer"
	"syntactic_checker/configuration_handler"
)

func Orchestrate_csv_cell_syntactic_checks() {

	run_configuration :=
		configuration_handler.Get_configuration()

	in_scope_identified_cells :=
		cells_preparer.Get_in_scope_identified_cells(
			run_configuration)

	syntactic_check_result_transaction_set :=
		cell_checks_runner.Run(
			in_scope_identified_cells,
			run_configuration)

	cell_checks_results_finaliser.Report_syntactic_check_outputs(
		syntactic_check_result_transaction_set)
}

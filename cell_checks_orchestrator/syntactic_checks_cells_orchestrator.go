package cell_checks_orchestrator

import (
	"syntactic_checker/cell_checks_orchestrator/cells_preparer"
	"syntactic_checker/cell_checks_orchestrator/checks_runner"
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
	"syntactic_checker/cell_checks_orchestrator/results_finaliser"
)

func Orchestrate_csv_cell_syntactic_checks() {

	run_configuration :=
		configuration_handler.Get_configuration()

	in_scope_identified_cells :=
		cells_preparer.Get_in_scope_identified_cells(
			run_configuration)

	syntactic_check_result_transaction_set :=
		checks_runner.Run(
			in_scope_identified_cells,
			run_configuration)

	results_finaliser.Report_syntactic_check_outputs(
		syntactic_check_result_transaction_set,
		run_configuration)
}

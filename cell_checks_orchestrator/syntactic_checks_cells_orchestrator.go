package cell_checks_orchestrator

import (
	"syntactic_checker/cell_checks_preparer"
	"syntactic_checker/cell_checks_results_finaliser"
	"syntactic_checker/cell_checks_runner"
	"syntactic_checker/configuration_handler"
)

func Orchestrate_csv_cell_syntactic_checks() {

	run_configuration :=
		configuration_handler.Get_configuration()

	csv_dataset :=
		cell_checks_preparer.Get_csv_cells_for_syntactic_checking(
			run_configuration)

	syntactic_check_result_transaction_set :=
		cell_checks_runner.Get_syntactic_checks_results(
			csv_dataset,
			run_configuration)

	//TODO - Stage 1 - add cell_checks_results_finaliser.Report_syntactic_check_issues(check_result_transaction_set)
	//TODO - Stage 1 - add cell_checks_results_finaliser.Report_syntactic_check_issue_parameters()
	cell_checks_results_finaliser.Report_syntactic_check_fixes(
		syntactic_check_result_transaction_set)
}

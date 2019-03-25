package cell_checks_runner

import (
	"syntactic_checker/cell_checks_orchestrator/cell_checks_results_finaliser"
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
)

func Run(
	in_scope_identified_cells [][]interface{},
	run_configuration *configuration_handler.Configurations) map[string][][]string {

	var syntactic_check_result_report map[string][][]string

	column_uuid :=
		run_configuration.Csv_configuration.Check_column_uuid

	syntactic_check_result_transactions :=
		process_syntactic_checks_for_cells(
			in_scope_identified_cells,
			run_configuration)

	syntactic_check_result_report =
		cell_checks_results_finaliser.Prepare_syntactic_checks_results_transactions(
			syntactic_check_result_transactions,
			column_uuid)

	return syntactic_check_result_report
}

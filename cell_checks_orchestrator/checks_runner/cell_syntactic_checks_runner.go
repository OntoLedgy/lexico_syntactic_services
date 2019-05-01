package checks_runner

import (
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
	"syntactic_checker/cell_checks_orchestrator/results_finaliser"
)

func Run(
	in_scope_identified_cells []map[string]interface{},
	run_configuration *configuration_handler.Configurations) map[string][][]string {

	var syntactic_check_result_report map[string][][]string

	column_uuid :=
		run_configuration.
			Csv_configuration.
			Check_column_uuid

	syntactic_check_result_transactions, cells_syntactic_check_fix_transactions :=
		process_syntactic_checks_for_cells(
			in_scope_identified_cells,
			run_configuration)

	syntactic_check_result_report =
		results_finaliser.Prepare_syntactic_checks_results_transactions(
			syntactic_check_result_transactions,
			cells_syntactic_check_fix_transactions,
			column_uuid)

	return syntactic_check_result_report
}

package orchestrators

import (
	"syntactic_checker/check_results_processor/finaliser"
	"syntactic_checker/helpers/cells_preparer"
	"syntactic_checker/helpers/configuration_handler"
	"syntactic_checker/orchestrators/column_checks"
)

func Orchestrate_syntactic_checks(
	configuration_file_path string) {

	run_configuration :=
		configuration_handler.Get_configuration(
			configuration_file_path)

	in_scope_cells :=
		cells_preparer.
			Get_in_scope_identified_cells(
				run_configuration)

	column_checker :=
		column_checks.
			Create_column_checker(
				in_scope_cells,
				run_configuration.Check_configuration.Issue_types)

	syntactic_check_results_transactions :=
		column_checker.
			RunColumnChecks()

	finaliser.
		Report_syntactic_check_outputs(
			syntactic_check_results_transactions,
			run_configuration)
}

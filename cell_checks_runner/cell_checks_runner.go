package cell_checks_runner

import (
	"fmt"
	storage_slices "storage/slices"
	"syntactic_checker/cell_checks_preparer"
	"syntactic_checker/cell_checks_results_finaliser"
	"syntactic_checker/configuration_handler"
)

func Get_syntactic_checks_results( //TODO - Stage 1 - move to syntactic cell check runner package
	in_scope_cells_interface [][]interface{},
	run_configuration *configuration_handler.Configurations) [][]string {

	var syntactic_check_result_transaction_set [][]string
	column_name :=
		run_configuration.Csv_configuration.Check_column_name
	in_scope_checks :=
		cell_checks_preparer.Get_in_scope_checks(run_configuration)

	fmt.Printf(
		"\nStarting processing. \ncolumn: %v\nin_scope_checks: %s\n",
		column_name,
		storage_slices.Pretty_print(in_scope_checks))

	check_transactions :=
		Process_cells_checks(
			in_scope_cells_interface,
			in_scope_checks)

	if check_transactions != nil {

		syntactic_check_result_transaction_set =
			cell_checks_results_finaliser.Prepare_syntactic_check_result_transactions_set(
				check_transactions,
				column_name)

	} else {
		syntactic_check_result_transaction_set = nil
	}

	//#TODO - Stage 1 - change issue report output format (check_uuids, check_type_uuids, object _uuids)
	return syntactic_check_result_transaction_set
}

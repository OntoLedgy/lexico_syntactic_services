package syntactic_check_cell_orchestrator

import (
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/configuration"
	"syntactic_checker/reporter"
)

func Orchestrate_csv_cell_syntactic_checks() {

	configuration, in_scope_checks :=
		configuration.Get_configuration() //#TODO - Stage 1 - include the in_scope_checks to the configuration structure

	fmt.Printf(
		"\nReading CSV Data for columns: %s\n",
		configuration.Csv_configuration.Check_column_name)

	csv_dataset :=
		storage.Read_csv_data(
			configuration.Csv_configuration.Csv_file_name)

	fmt.Printf(
		"\nStarting processing. \ncolumn: %v\nin_scope_checks: %s\n",
		configuration.Csv_configuration.Check_column_name,
		storage_slices.Pretty_print(in_scope_checks))

	syntactic_check_result_transaction_set :=
		Get_syntactic_checks_results(
			csv_dataset,
			configuration.Csv_configuration.Check_column_name,
			in_scope_checks)

	//TODO - Stage 1 - add reporter.Report_syntactic_check_issues(check_result_transaction_set)
	//TODO - Stage 1 - add reporter.Report_syntactic_check_issue_parameters()
	reporter.Report_syntactic_check_fixes(
		syntactic_check_result_transaction_set)

}

func Get_syntactic_checks_results(
	in_scope_cellset_data [][]string,
	column_name string, // TODO - Stage 1 - change to column uuid
	in_scope_checks [][]interface{}) [][]string {

	var syntactic_check_result_transaction_set [][]string

	in_scope_cellset_data_interface :=
		Prepare_cellset_data(
			in_scope_cellset_data)

	check_transactions :=
		Process_cells_checks(
			in_scope_cellset_data_interface,
			in_scope_checks)

	if check_transactions != nil {

		syntactic_check_result_transaction_set =
			Prepare_syntactic_check_result_transaction_set(
				check_transactions,
				column_name)

	} else {
		syntactic_check_result_transaction_set = nil
	}

	//#TODO - Stage 1 - change issue report output format (check_uuids, check_type_uuids, object _uuids)
	return syntactic_check_result_transaction_set
}

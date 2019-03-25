package cells_preparer

import (
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
)

func Get_in_scope_identified_cells(
	run_configuration *configuration_handler.Configurations) [][]interface{} {

	csv_filename :=
		run_configuration.Csv_configuration.Csv_file_name

	in_scope_identified_cells :=
		prepare_in_scope_identified_cells_data(
			csv_filename)

	return in_scope_identified_cells
}

func prepare_in_scope_identified_cells_data(
	csv_filename string) [][]interface{} {

	fmt.Printf(
		"\nReading CSV Data..")

	in_scope_identified_cells_raw :=
		storage.Read_csv_data(
			csv_filename)

	fmt.Printf(
		"Preparing extracted data for checks (converting to interface)")

	in_scope_identified_cells_interface :=
		storage_slices.Convert_2d_string_to_interface(
			in_scope_identified_cells_raw)

	return in_scope_identified_cells_interface
}

package cell_checks_preparer

import (
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/configuration_handler"
)

func Get_csv_cells_for_syntactic_checking( //TODO - Stage 1 - move to syntactic check preparer
	run_configuration *configuration_handler.Configurations) [][]interface{} {

	csv_filename :=
		run_configuration.Csv_configuration.Csv_file_name

	check_column_name :=
		run_configuration.Csv_configuration.Check_column_name

	fmt.Printf(
		"\nReading CSV Data for columns: %s\n",
		check_column_name)

	csv_dataset :=
		storage.Read_csv_data(
			csv_filename)

	in_scope_cellset_data_interface :=
		prepare_cells_data(
			csv_dataset)

	return in_scope_cellset_data_interface

}

func prepare_cells_data(
	in_scope_cells [][]string) [][]interface{} {

	fmt.Printf(
		"Preparing extracted data for checks (converting to interface)")

	in_scope_cells_interface :=
		storage_slices.Convert_2d_string_to_interface(
			in_scope_cells)

	return in_scope_cells_interface
}

func Get_in_scope_checks(
	run_configuration *configuration_handler.Configurations) [][]interface{} {

	var in_scope_check_interface []interface{}
	var in_scope_checks_interface [][]interface{}

	in_scope_checks :=
		run_configuration.Csv_configuration.Issue_types

	for _, in_scope_check := range in_scope_checks {

		in_scope_check_interface = append(in_scope_check_interface,
			in_scope_check.Issue_type_uuid,
			in_scope_check.Issue_type_name,
			in_scope_check.Issue_check_regex,
			in_scope_check.Issue_check_replacement_string)
		in_scope_checks_interface = append(in_scope_checks_interface, in_scope_check_interface)
		in_scope_check_interface = nil
	}

	return in_scope_checks_interface
}

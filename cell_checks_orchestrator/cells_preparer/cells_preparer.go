package cells_preparer

import (
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
)

func Get_in_scope_identified_cells(
	run_configuration *configuration_handler.Configurations) []map[string]interface{} {

	csv_filename :=
		run_configuration.Csv_configuration.Csv_file_name

	in_scope_identified_cells :=
		prepare_in_scope_identified_cells_data(
			csv_filename)

	return in_scope_identified_cells
}

func prepare_in_scope_identified_cells_data(
	csv_filename string) []map[string]interface{} {

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

	in_scope_identified_cells_with_headers :=
		get_csv_with_headers(
			in_scope_identified_cells_interface)

	return in_scope_identified_cells_with_headers
}

func get_csv_with_headers(csv_data_with_headers [][]interface{}) []map[string]interface{} {

	var rows []map[string]interface{}

	var header []interface{}

	for index, csv_data_with_headers_row := range csv_data_with_headers {

		if index == 0 {
			header = csv_data_with_headers_row
		} else {

			dict := make(map[string]interface{})

			for i := range header {
				dict[header[i].(string)] = csv_data_with_headers_row[i]
			}
			rows = append(rows, dict)

		}

	}

	return rows

}

package syntactic_check_processor

import (
	"encoding/json"
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/configuration"
)

//Stip -
func Orchestrate_csv_syntactic_checks(
	configuration configuration.Csv_configuration,
	in_scope_checks [][]interface{}) [][]string {

	in_scope_column_positions :=
		Prepare_csv_check_column_set(
			configuration.Identity_column_position,
			configuration.In_scope_check_column_positions)

	fmt.Printf(
		"\nReading CSV Data for columns: %s\n",
		in_scope_column_positions)

	csv_dataset :=
		storage.Read_csv_data(
			configuration.Csv_file_name)

	fmt.Printf(
		"\nStarting processing. \ncolumns: %v\nin_scope_checks: %s\n",
		in_scope_column_positions,
		Pretty_print(in_scope_checks))

	check_transaction_set :=
		Process_csv_columns(
			csv_dataset,
			in_scope_column_positions,
			in_scope_checks)

	//TODO - Stage 1 - add returns for issue parameters, and fixes

	return check_transaction_set

}

func Prepare_csv_check_column_set(
	identity_column_number int,
	check_columns map[string]int) map[string][]int {

	column_set :=
		make(map[string][]int)

	for check_column_name := range check_columns {

		column_set[check_column_name] =
			append(
				column_set[check_column_name],
				identity_column_number,
				check_columns[check_column_name])
	}

	return column_set
}

func Process_csv_columns(
	csv_dataset [][]string,
	column_set map[string][]int,
	in_scope_checks [][]interface{}) [][]string {

	var checked_columns_transaction_set [][]string

	for column_name := range column_set {

		fmt.Printf(
			"\nExtracting Columns %s\n",
			column_set)

		extracted_data :=
			storage_slices.Extract_columns_from_2d_slices(
				csv_dataset,
				column_set[column_name])

		checked_column_transaction_set :=
			Process_csv_column(
				extracted_data,
				column_name,
				in_scope_checks)

		checked_columns_transaction_set =
			append(
				checked_columns_transaction_set,
				checked_column_transaction_set...)

		checked_column_transaction_set =
			nil

	}

	return checked_columns_transaction_set

}

func Process_csv_column(
	extracted_data [][]string,
	column_name string, // TODO - Stage 1 - change to column uuid
	in_scope_checks [][]interface{}) [][]string {

	var check_transaction_set [][]string

	for _, check := range in_scope_checks {

		fmt.Println(
			"Adding check type uuid to extracted data\n")

		csv_dataset_including_check :=
			storage_slices.Add_single_value_column_to_2d_slice(
				extracted_data,
				check[0].(string)) //

		fmt.Printf(
			"csv_dataset_including_check: %s\n",
			csv_dataset_including_check[0])

		fmt.Print(
			"converting cell data to interface\n")

		csv_dataset_including_check_interface :=
			storage_slices.Convert_2d_string_to_interface(
				csv_dataset_including_check)

		fmt.Printf(
			"processing checks: %s\n",
			in_scope_checks)

		check_transactions :=
			Process_cell_checks(
				csv_dataset_including_check_interface,
				in_scope_checks)

		if check_transactions != nil {

			check_transactions_string :=
				storage_slices.Convert_2d_interface_to_string(
					check_transactions)

			check_transactions_string =
				storage_slices.Add_single_value_column_to_2d_slice(
					check_transactions_string,
					column_name)

			for _, check_transaction := range check_transactions_string {
				check_transaction_set =
					append(
						check_transaction_set,
						check_transaction)
			}

		}

		csv_dataset_including_check = nil

	}
	//#TODO - Stage 1 - change issue report output format (check_uuids, check_type_uuids, object _uuids)
	return check_transaction_set

}

func Pretty_print(json_data [][]interface{}) []byte {

	pretty_printed_json, _ := json.MarshalIndent(
		json_data,
		"",
		"	") //#TODO add pretty_printer to general utilities

	return pretty_printed_json

} //TODO - Stage 1 - move to json utilities

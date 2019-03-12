package syntactic_checks

import (
	"fmt"
	storage_slices "storage/slices"
	"syntactic_checker/reporting"
)

func Process_csv_syntactic_check_columns(
	csv_dataset [][]string,
	column_set map[string][]int,
	in_scope_checks [][]interface{}) {

	var check_transaction_set [][]string

	for column_name := range column_set {

		fmt.Printf("\nExtracting Columns %s\n", column_set)
		extracted_data := storage_slices.Extract_columns_from_2d_slices(csv_dataset, column_set[column_name])

		column_check_transaction_set := Process_csv_syntactic_check_column(extracted_data, column_name, in_scope_checks)

		check_transaction_set = append(check_transaction_set, column_check_transaction_set...)

		column_check_transaction_set = nil

	}

	reporting.Report_syntactic_check_issues(check_transaction_set)

}

func Process_csv_syntactic_check_column(
	extracted_data [][]string,
	column_name string,
	in_scope_checks [][]interface{}) [][]string {

	var check_transaction_set [][]string

	for _, check := range in_scope_checks {

		fmt.Println("Adding check uuid to extracted data\n")
		csv_dataset_including_check := storage_slices.Add_single_value_column_to_2d_slice(
			extracted_data,
			check[0].(string))

		fmt.Printf(
			"csv_dataset_including_check: %s\n",
			csv_dataset_including_check[0])

		fmt.Print("converting cell data to interface\n")
		csv_dataset_including_check_interface :=
			storage_slices.Convert_2d_string_to_interface(
				csv_dataset_including_check)

		fmt.Print("processing checks\n")
		check_transactions :=
			process_cell_sets(
				csv_dataset_including_check_interface,
				in_scope_checks)

		if check_transactions != nil {

			check_transactions_string :=
				storage_slices.Convert_2d_interface_to_string(
					check_transactions)

			check_transactions_string = storage_slices.Add_single_value_column_to_2d_slice(
				check_transactions_string,
				column_name)

			for _, check_transaction := range check_transactions_string {
				check_transaction_set = append(
					check_transaction_set,
					check_transaction)
			}

		}

		csv_dataset_including_check = nil

	}
	//#TODO - Stage 2 - change issue report output format (check_uuids, check_type_uuids, object _uuids)

	return check_transaction_set

}

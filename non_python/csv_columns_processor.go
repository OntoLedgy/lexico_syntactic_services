package non_python

import (
	"fmt"
	storage_slices "storage/slices"
	"syntactic_checker/syntactic_check_cell_orchestrator"
)

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
			syntactic_check_cell_orchestrator.Get_syntactic_checks_results(
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

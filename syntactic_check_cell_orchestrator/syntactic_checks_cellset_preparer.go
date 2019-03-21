package syntactic_check_cell_orchestrator

import (
	"fmt"
	storage_slices "storage/slices"
)

func Prepare_cellset_data(in_scope_cellset_data [][]string) [][]interface{} {
	fmt.Printf(
		"Preparing extracted data for checks (converting to interface)") //TODO - Stage 1 - check if interface conversion is really needed

	in_scope_cellset_data_interface :=
		storage_slices.Convert_2d_string_to_interface(
			in_scope_cellset_data)

	return in_scope_cellset_data_interface
}

func Prepare_syntactic_check_result_transaction_set(
	check_transactions [][]interface{},
	column_name string) [][]string {

	var check_transaction_set [][]string

	check_transactions_string :=
		Prepare_syntactic_check_result_transactions(
			check_transactions,
			column_name)

	for _, check_transaction := range check_transactions_string {
		check_transaction_set =
			append(
				check_transaction_set,
				check_transaction)
	}

	return check_transaction_set
}

func Prepare_syntactic_check_result_transactions(
	check_transactions [][]interface{},
	column_name string) [][]string {

	check_transactions_string :=
		storage_slices.Convert_2d_interface_to_string(
			check_transactions)

	check_transactions_string =
		storage_slices.Add_single_value_column_to_2d_slice(
			check_transactions_string,
			column_name)

	return check_transactions_string

}

package cell_checks_results_finaliser

import "storage/slices"

func Prepare_syntactic_check_result_transactions_set( //TODO - Stage 1 - move to finaliser
	check_transactions [][]interface{},
	column_name string) [][]string {

	var check_transaction_set [][]string

	check_transactions_string :=
		prepare_syntactic_check_result_transactions(
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

func prepare_syntactic_check_result_transactions(
	check_transactions [][]interface{},
	column_name string) [][]string {

	check_transactions_string :=
		storage.Convert_2d_interface_to_string(
			check_transactions)

	check_transactions_string =
		storage.Add_single_value_column_to_2d_slice(
			check_transactions_string,
			column_name)

	return check_transactions_string

}

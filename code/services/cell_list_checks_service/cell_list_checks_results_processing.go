package cell_list_checks_service

func process_syntactic_check_fix_transactions(

	cell_syntactic_check_fix_transaction []interface{},
	cells_syntactic_check_fix_transactions [][]interface{}) [][]interface{} {

	if cell_syntactic_check_fix_transaction != nil {

		cells_syntactic_check_fix_transactions =
			append(
				cells_syntactic_check_fix_transactions,
				cell_syntactic_check_fix_transaction)

		cell_syntactic_check_fix_transaction = nil

	}
	return cells_syntactic_check_fix_transactions
}

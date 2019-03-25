package cell_checks_runner

import (
	"fmt"
	"syntactic_checker/regex_processor"
)

func process_syntactic_checks_for_cells(
	in_scope_identified_cells [][]interface{},
	in_scope_check_types [][]interface{}) [][]interface{} {

	var check_result_transaction_set [][]interface{}
	var cell_check_result_transaction_set [][]interface{}

	fmt.Printf(
		"processing checks: %s\n",
		in_scope_check_types)

	for _, in_scope_cell := range in_scope_identified_cells {

		cell_check_result_transaction_set =
			process_syntactic_checks_for_cell(
				in_scope_cell,
				in_scope_check_types)

		check_result_transaction_set = append(check_result_transaction_set,
			cell_check_result_transaction_set...)

	}

	/*
		if check_result_transaction_set != nil { //#TODO add to logger
			fmt.Printf(
				"\n++++++++++++++++++++++++Issue_types completed: %s issues found.\nExample transaction row: %s ",
				len(check_result_transaction_set),
				check_result_transaction_set[0])
		}*/
	return check_result_transaction_set
}

func process_syntactic_checks_for_cell(
	in_scope_identified_cell []interface{},
	in_scope_checks [][]interface{}) [][]interface{} {

	var cell_check_result_transaction_set [][]interface{}
	var cell_syntactic_check_result_transaction []interface{}

	for _, in_scope_check := range in_scope_checks {

		cell_syntactic_check_result_transaction =
			process_syntactic_check_for_cell(
				in_scope_identified_cell,
				in_scope_check)

		cell_check_result_transaction_set =
			process_syntactic_check_result_transaction(
				in_scope_identified_cell,
				in_scope_check,
				cell_syntactic_check_result_transaction,
				cell_check_result_transaction_set)

		cell_syntactic_check_result_transaction = nil
	}

	return cell_check_result_transaction_set
}

func process_syntactic_check_for_cell(
	in_scope_identified_cell []interface{},
	in_scope_check []interface{}) []interface{} {

	var cell_check_result_transaction []interface{}

	if in_scope_identified_cell[1] != nil { // if cell value is not 'null'

		cell_check_result_transaction = //#TODO - Stage 2 - add switch to include non-regex in_scope_check types.
			regex_processor.
				Process_regex_check( // run regex in_scope_check
					in_scope_check[2].(string),  // in_scope_check regex string
					in_scope_identified_cell[1], // cell value
					in_scope_check[3].(string))  // in_scope_check regex replacement value
	} else {
		cell_check_result_transaction =
			nil
	}

	return cell_check_result_transaction
}

func process_syntactic_check_result_transaction(
	in_scope_identified_cell []interface{},
	in_scope_check []interface{},
	cell_syntactic_check_result_transaction []interface{},
	cell_sytactic_check_result_transaction_set [][]interface{}) [][]interface{} {

	if cell_syntactic_check_result_transaction != nil {
		//if issues are found append information to the transaction
		cell_syntactic_check_result_transaction =
			append(
				cell_syntactic_check_result_transaction,
				in_scope_check[0],           // add in_scope_check type
				in_scope_identified_cell[0]) // add row uuid

		cell_sytactic_check_result_transaction_set = //append to transaction register
			append(cell_sytactic_check_result_transaction_set,
				cell_syntactic_check_result_transaction)

		fmt.Printf(
			"\n\ncell_sytactic_check_result_transaction_set:%v\n\n\n",
			cell_sytactic_check_result_transaction_set)
	}

	return cell_sytactic_check_result_transaction_set
}

package syntactic_check_cell_orchestrator

import (
	"fmt"
	"syntactic_checker/regex_processor"
)

func Process_cells_checks(
	in_scope_cells [][]interface{},
	in_scope_check_types [][]interface{}) [][]interface{} {

	fmt.Printf(
		"processing checks: %s\n",
		in_scope_check_types)

	var check_result_transaction_set [][]interface{}
	var cell_check_result_transaction_set [][]interface{}

	for _, in_scope_cell := range in_scope_cells { // for each cell row

		cell_check_result_transaction_set =
			Process_cell_checks(
				in_scope_cell,
				in_scope_check_types)

		check_result_transaction_set = append(check_result_transaction_set,
			cell_check_result_transaction_set...)

	}

	/*
		if check_result_transaction_set != nil { //#TODO add to logger
			fmt.Printf(
				"\n++++++++++++++++++++++++Checks completed: %s issues found.\nExample transaction row: %s ",
				len(check_result_transaction_set),
				check_result_transaction_set[0])
		}*/
	return check_result_transaction_set
}

func Process_cell_checks(
	in_scope_cell []interface{},
	in_scope_checks [][]interface{}) [][]interface{} {

	var cell_check_result_transaction_set [][]interface{}
	var cell_check_result_transaction []interface{}

	for _, in_scope_check := range in_scope_checks { // for each in_scope_check type in in_scope_check

		cell_check_result_transaction = Process_cell_check(
			in_scope_cell,
			in_scope_check)

		cell_check_result_transaction_set = Process_cell_check_result_transaction(
			in_scope_cell,
			in_scope_check,
			cell_check_result_transaction,
			cell_check_result_transaction_set)

		cell_check_result_transaction = nil

	}

	return cell_check_result_transaction_set

}

func Process_cell_check(
	in_scope_cell []interface{},
	in_scope_check []interface{}) []interface{} {

	var cell_check_result_transaction []interface{}

	if in_scope_cell[1] != nil { // if cell value is not 'null'

		cell_check_result_transaction = //#TODO - Stage 2 - add switch to include non-regex in_scope_check types.
			regex_processor.
				Process_regex_check( // run regex in_scope_check
					in_scope_check[2].(string), // in_scope_check regex string
					in_scope_cell[1],           // cell value
					in_scope_check[3].(string)) // in_scope_check regex replacement value
	} else {
		cell_check_result_transaction =
			nil
	}

	return cell_check_result_transaction
}

func Process_cell_check_result_transaction(
	in_scope_cell []interface{},
	in_scope_check []interface{},
	cell_check_result_transaction []interface{},
	cell_check_result_transaction_set [][]interface{}) [][]interface{} {

	if cell_check_result_transaction != nil { //#TODO - Stage 1 - move this out to separate function ( process result transactions)
		//if issues are found append information to the transaction
		cell_check_result_transaction =
			append(
				cell_check_result_transaction,
				in_scope_check[0], // add in_scope_check type
				in_scope_cell[0])  // add row uuid

		cell_check_result_transaction_set = //append to transaction register
			append(cell_check_result_transaction_set,
				cell_check_result_transaction)

		fmt.Printf(
			"\n\ncell_check_result_transaction_set:%v\n\n\n",
			cell_check_result_transaction_set)
	}

	return cell_check_result_transaction_set
}

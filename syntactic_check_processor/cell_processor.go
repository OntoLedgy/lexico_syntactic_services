package syntactic_check_processor

import (
	"syntactic_checker/regex_processor"
)

func Process_cell_checks(
	in_scope_cell_dataset [][]interface{},
	checks [][]interface{}) [][]interface{} {

	var check_result_transaction_set [][]interface{}
	var check_result_transaction []interface{}

	for _, in_scope_cell_row := range // for each cell row
	in_scope_cell_dataset {

		for _, check := range //#TODO - Stage 1 - remove this loop (iteration by check type already happes in the calling function)
		// for each check type in check register
		checks {
			if check[0] == // if check type uuid is included in check configuration
				in_scope_cell_row[2] &&
				in_scope_cell_row[1] != // AND cell value is not 'null'
					nil {

				/*check_string_string_value := in_scope_cell_row[1].(string)

				for _,char := range check_string_string_value {
					fmt.Printf("\n%c",char)
				}*/

				check_result_transaction = //#TODO - Stage 2 - add switch to include non-regex check types.
					regex_processor.
						Process_regex_check( // run regex check
							check[2].(string),    // check regex string
							in_scope_cell_row[1], // cell value
							check[3].(string))    // check regex replacement value

				if check_result_transaction != nil { //#TODO - Stage 1 - move this out to separate function ( process result transactions)

					//if issues are found append information to the transaction

					check_result_transaction =
						append(
							check_result_transaction,
							in_scope_cell_row[2], // add check type
							in_scope_cell_row[0]) // add row uuid

					check_result_transaction_set = //append to transaction register
						append(check_result_transaction_set,
							check_result_transaction)
				}
				check_result_transaction = nil

			}
		}

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

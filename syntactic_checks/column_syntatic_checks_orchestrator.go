package syntactic_checks

import (
	"fmt"
	"syntactic_checker/regex_management"
)

func Process_column_sets(
	in_scope_columns_dataset [][]interface{},
	checks [][]interface{},
	in_scope_column_check_configurations [][]interface{}) [][]interface{} {

	var in_scope_column_cell_set_including_check_configuration [][]interface{} //#TODO explore alternative data structures for this.  possibly split this into column cell set and check configuration
	var in_scope_column_cell_and_check_configuration []interface{}             //#TODO explore alternative data structures for this.  possibly split this into column cell set and check configuration

	var transaction_rowset [][]interface{}

	for _, in_scope_column_row := range // for each row in the columns dataset
	in_scope_columns_dataset {

		for _, in_scope_column_check_configuration := range // for each entry in check configurations
		in_scope_column_check_configurations {

			if in_scope_column_check_configuration[0] == // check if column uuid in check configuration
				in_scope_column_row[2] { // matches column uuid in column dataset -- add to cell check dataset

				in_scope_column_cell_and_check_configuration := // create cell data for checks (strip column uuid and add check uuid) #TODO - chedck if this is the best way to create the slide or alternatives available
					append(
						in_scope_column_cell_and_check_configuration, // nil
						in_scope_column_row[0],                       // row uuid
						in_scope_column_row[1],                       // cell value
						in_scope_column_check_configuration[1])       // check uuid

				in_scope_column_cell_set_including_check_configuration = // add to cell configuration set (contains a row for each check for the cell value)
					append(
						in_scope_column_cell_set_including_check_configuration,
						in_scope_column_cell_and_check_configuration)

				in_scope_column_cell_and_check_configuration =
					nil
			}

		}

		if in_scope_column_cell_set_including_check_configuration != nil { //if the column is marked for checks

			transaction_column_rowset :=
				process_cell_sets( // run check process
					in_scope_column_cell_set_including_check_configuration,
					checks)

			for _, transaction_row := range transaction_column_rowset { // for each returned check transaction, append the column uuid

				transaction_row = append(
					transaction_row,        // check transaction
					in_scope_column_row[2]) // column uuid

				transaction_rowset = append( // append the transaction (+ column uuid) to transaction set.
					transaction_rowset,
					transaction_row)

				transaction_row =
					nil
			}
			in_scope_column_cell_set_including_check_configuration =
				nil
		}

	}

	fmt.Printf( //#TODO add to logger
		"\n--Total %s transactions generated.\n--------------\nSample Transaction : %s\n-----------\n",
		len(transaction_rowset),
		transaction_rowset[0])

	return transaction_rowset

}

func process_cell_sets(
	in_scope_cell_dataset [][]interface{},
	checks [][]interface{}) [][]interface{} {

	var check_result_transaction_set [][]interface{}
	var check_result_transaction []interface{}

	for _, in_scope_cell_row := range // for each cell row
	in_scope_cell_dataset {

		for _, check := range // for each check type in check register
		checks {

			if check[0] == // if check type uuid is included in check configuration
				in_scope_cell_row[2] &&

				in_scope_cell_row[1] != // AND cell value is not 'null'
					nil {

				/*check_string_string_value := in_scope_cell_row[1].(string)

				for _,char := range check_string_string_value {
					fmt.Printf("\n%c",char)
				}*/

				check_result_transaction = //#TODO this is has be generalised to include non-regex check types.
					regex_management.Process_regex_check( // run regex check
						check[2].(string),    // check regex string
						in_scope_cell_row[1], // cell value
						check[3].(string))    // check regex replacement value

				if check_result_transaction != nil { //if issues are found append information to the transaction

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

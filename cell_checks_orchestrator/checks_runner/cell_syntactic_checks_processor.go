package checks_runner

import (
	"fmt"
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
	"syntactic_checker/cell_checks_orchestrator/regex_processor"
	"syntactic_checker/object_model"
)

func process_syntactic_checks_for_cells(
	in_scope_identified_cells [][]interface{},
	run_configuration *configuration_handler.Configurations) [][]interface{} {

	var check_result_transaction_set [][]interface{}
	var cell_check_result_transaction_set [][]interface{}

	fmt.Printf(
		"processing checks: %s\n",
		run_configuration)

	for _, in_scope_cell := range in_scope_identified_cells {

		cell_check_result_transaction_set =
			process_syntactic_checks_for_cell(
				in_scope_cell,
				run_configuration)

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
	run_configuration *configuration_handler.Configurations) [][]interface{} {

	var cell_check_result_transaction_set [][]interface{}
	var cell_syntactic_check_result_transaction []interface{}

	in_scope_checks :=
		run_configuration.Csv_configuration.Issue_types

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
	in_scope_check object_model.Issue_types) []interface{} {

	var cell_check_result_transaction []interface{}

	if in_scope_identified_cell[1] != nil {

		cell_check_result_transaction = //TODO - Stage 3 - add switch to include non-regex in_scope_check types.
			regex_processor.
				Process_regex_check( //TODO - Stage 2 - replace the check interface with the check type object and pass it through
					in_scope_check.Issue_check_regex,
					in_scope_identified_cell[1], // cell value
					in_scope_check.Issue_check_replacement_string)
	} else {
		cell_check_result_transaction = //TODO - Stage 1 - if cell value is null - report cell is null error
			nil
	}
	return cell_check_result_transaction
}

func process_syntactic_check_result_transaction(
	in_scope_identified_cell []interface{},
	in_scope_check object_model.Issue_types,
	cell_syntactic_check_result_transaction []interface{},
	cell_sytactic_check_result_transaction_set [][]interface{}) [][]interface{} {

	if cell_syntactic_check_result_transaction != nil {
		//if issues are found append information to the transaction
		cell_syntactic_check_result_transaction =
			append(
				cell_syntactic_check_result_transaction,
				in_scope_check.Issue_type_uuid,
				in_scope_identified_cell[0]) //TODO - Stage 2 - replace with human readable column name

		cell_sytactic_check_result_transaction_set = //append to transaction register
			append(cell_sytactic_check_result_transaction_set,
				cell_syntactic_check_result_transaction)

		fmt.Printf(
			"\n\ncell_sytactic_check_result_transaction_set:%v\n\n\n",
			cell_sytactic_check_result_transaction_set)
	}

	return cell_sytactic_check_result_transaction_set
}

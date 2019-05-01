package checks_runner

import (
	"database_manager/utils"
	"fmt"
	//"storage/slices"
	"syntactic_checker/cell_checks_orchestrator/cell_fixer"
	"syntactic_checker/cell_checks_orchestrator/configuration_handler"
	"syntactic_checker/cell_checks_orchestrator/regex_processor"
	"syntactic_checker/object_model"
)

func process_syntactic_checks_for_cells(
	in_scope_identified_cells []map[string]interface{},
	run_configuration *configuration_handler.Configurations) ([][]interface{}, [][]interface{}) {

	var cells_syntactic_check_issues_transactions [][]interface{}
	var cell_syntactic_check_issues_transactions [][]interface{}
	var cell_syntactic_check_fix_transaction []interface{}
	var cells_syntactic_check_fix_transactions [][]interface{}

	fmt.Printf(
		"processing checks: %s\n",
		run_configuration)

	for _, in_scope_cell := range in_scope_identified_cells {

		cell_syntactic_check_issues_transactions, cell_syntactic_check_fix_transaction =
			process_syntactic_checks_for_cell(
				in_scope_cell,
				run_configuration)

		cells_syntactic_check_issues_transactions =
			append(
				cells_syntactic_check_issues_transactions,
				cell_syntactic_check_issues_transactions...)

		if cell_syntactic_check_fix_transaction != nil {

			cells_syntactic_check_fix_transactions =
				append(
					cells_syntactic_check_fix_transactions,
					cell_syntactic_check_fix_transaction)

			cell_syntactic_check_fix_transaction = nil

		}

	}

	/*
		if cells_syntactic_check_issues_transactions != nil { //#TODO add to logger
			fmt.Printf(
				"\n++++++++++++++++++++++++Issue_types completed: %s issues found.\nExample transaction row: %s ",
				len(cells_syntactic_check_issues_transactions),
				cells_syntactic_check_issues_transactions[0])
		}*/
	return cells_syntactic_check_issues_transactions, cells_syntactic_check_fix_transactions
}

func process_syntactic_checks_for_cell(
	in_scope_identified_cell map[string]interface{},
	run_configuration *configuration_handler.Configurations) ([][]interface{}, []interface{}) {

	var cell_syntactic_check_issue_transactions [][]interface{}
	var cell_syntactic_check_aggregated_fixes_transaction []interface{}
	var cell_syntactic_check_issue_results []object_model.Regex_check_results

	in_scope_syntactic_check_types :=
		run_configuration.Csv_configuration.Issue_types

	//TODO - Stage 2 - move the for loop to a function (process syntactic checks issues for cell)
	for _, in_scope_syntactic_check_type := range in_scope_syntactic_check_types {

		cell_syntactic_check_issue_result :=
			process_syntactic_check_for_cell(
				in_scope_identified_cell,
				run_configuration.Csv_configuration.Check_column_name,
				in_scope_syntactic_check_type)

		cell_syntactic_check_issue_results, cell_syntactic_check_issue_transactions =
			aggregate_syntactic_check_issues(
				in_scope_identified_cell,
				run_configuration,
				in_scope_syntactic_check_type,
				cell_syntactic_check_issue_results,
				cell_syntactic_check_issue_result,
				cell_syntactic_check_issue_transactions)
	}

	// TODO - Stage 2 - move the for loop

	if cell_syntactic_check_issue_results != nil {

		fix_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		cell_syntactic_check_aggregated_fixes_transaction =
			aggregate_cell_syntactic_check_fixes(
				cell_syntactic_check_issue_results)

		cell_syntactic_check_aggregated_fixes_transaction =
			append(
				cell_syntactic_check_aggregated_fixes_transaction,
				in_scope_identified_cell[run_configuration.Csv_configuration.Identity_column_name].(string),
				fix_uuid.String())

	}

	fmt.Printf("\n--cell fix transaction: %v \n", cell_syntactic_check_aggregated_fixes_transaction)

	return cell_syntactic_check_issue_transactions, cell_syntactic_check_aggregated_fixes_transaction
}

func aggregate_syntactic_check_issues(
	in_scope_identified_cell map[string]interface{},
	run_configuration *configuration_handler.Configurations,
	in_scope_check object_model.Issue_types,
	cell_syntactic_check_results []object_model.Regex_check_results,
	regex_check_issue_result *object_model.Regex_check_results,
	cell_syntactic_check_issue_transactions [][]interface{}) ([]object_model.Regex_check_results, [][]interface{}) {

	var cell_syntactic_check_issue_transaction []interface{}

	if regex_check_issue_result != nil {

		cell_syntactic_check_issue_transaction =
			cell_fixer.Generate_issue_transaction(
				*regex_check_issue_result)

		cell_syntactic_check_issue_transactions =
			aggregate_syntactic_check_issue_transactions(
				in_scope_identified_cell,
				run_configuration,
				in_scope_check,
				cell_syntactic_check_issue_transaction,
				cell_syntactic_check_issue_transactions)

		cell_syntactic_check_results =
			aggregate_syntactic_check_issue_result_transactions(
				*regex_check_issue_result,
				cell_syntactic_check_results)

	}

	cell_syntactic_check_issue_transaction = nil

	return cell_syntactic_check_results, cell_syntactic_check_issue_transactions
}

func process_syntactic_check_for_cell(
	in_scope_identified_cell map[string]interface{},
	check_column_name string,
	in_scope_check object_model.Issue_types) *object_model.Regex_check_results {

	var regex_check_result *object_model.Regex_check_results

	if in_scope_identified_cell[check_column_name] != nil {

		regex_check_result = //TODO - Stage 3 - add switch to include non-regex in_scope_check types.
			regex_processor.
				Process_regex_check( //TODO - Stage 2 - replace the check interface with the check type object and pass it through
					in_scope_check.Issue_check_regex,
					in_scope_identified_cell[check_column_name], // cell value
					in_scope_check.Issue_check_replacement_string)

	} else {
		fmt.Printf(
			"\nWARNING: cell value for row_id%s is null\n",
			in_scope_identified_cell[check_column_name])

	}
	return regex_check_result
}

func aggregate_cell_syntactic_check_fixes(
	cell_syntactic_check_results []object_model.Regex_check_results) []interface{} {

	var cell_syntactic_check_aggregated_fixes_transaction []interface{}
	var interim_modified_string string
	var interim_marked_string string

	cell_syntactic_check_aggregated_fixes_transaction = nil

	original_string :=
		cell_syntactic_check_results[0].Original_string

	interim_modified_string = original_string

	for _, cell_syntactic_check_result := range cell_syntactic_check_results {

		interim_modified_string, interim_marked_string =
			generate_cell_syntactic_check_fix(
				cell_syntactic_check_result,
				interim_modified_string)

	}

	cell_syntactic_check_aggregated_fixes_transaction =
		append(
			cell_syntactic_check_aggregated_fixes_transaction,
			cell_syntactic_check_results[0].Original_string,
			interim_marked_string,
			interim_modified_string)

	return cell_syntactic_check_aggregated_fixes_transaction
}

func generate_cell_syntactic_check_fix(
	cell_syntactic_check_result object_model.Regex_check_results,
	interim_modified_string string) (string, string) {

	interim_modified_string_next :=
		cell_fixer.Modify_string_by_index(
			interim_modified_string,
			cell_syntactic_check_result.Replacement_string,
			cell_syntactic_check_result.Regex_match_indices)

	interim_marked_string :=
		cell_fixer.Modify_string_by_index(
			interim_modified_string,
			cell_syntactic_check_result.Mark_string,
			cell_syntactic_check_result.Regex_match_indices)

	return interim_modified_string_next, interim_marked_string
}

func aggregate_syntactic_check_issue_result_transactions(
	regex_check_issue_result object_model.Regex_check_results,
	cell_syntactic_check_results []object_model.Regex_check_results) []object_model.Regex_check_results {

	cell_syntactic_check_results = append(
		cell_syntactic_check_results,
		regex_check_issue_result)

	return cell_syntactic_check_results

}

func aggregate_syntactic_check_issue_transactions(
	in_scope_identified_cell map[string]interface{},
	run_configuration *configuration_handler.Configurations,
	in_scope_check object_model.Issue_types,
	cell_syntactic_check_result_transaction []interface{},
	cell_sytactic_check_result_transaction_set [][]interface{}) [][]interface{} {

	if cell_syntactic_check_result_transaction != nil {
		//if issues are found append information to the transaction
		cell_syntactic_check_result_transaction =
			append(
				cell_syntactic_check_result_transaction,
				in_scope_check.Issue_type_uuid,
				in_scope_identified_cell[run_configuration.Csv_configuration.Identity_column_name]) //TODO - Stage 2 - replace with human readable column name

		cell_sytactic_check_result_transaction_set = //append to transaction register
			append(cell_sytactic_check_result_transaction_set,
				cell_syntactic_check_result_transaction)

		fmt.Printf(
			"\n\ncell_sytactic_check_result_transaction_set:%v\n\n\n",
			cell_sytactic_check_result_transaction_set)
	}

	return cell_sytactic_check_result_transaction_set
}

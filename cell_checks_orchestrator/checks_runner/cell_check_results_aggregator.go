package checks_runner

import (
	"fmt"
	"syntactic_checker/cell_checks_orchestrator/cell_fixer"
	"syntactic_checker/object_model"
)

func aggregate_syntactic_check_issues(
	in_scope_identified_cell map[string]interface{},
	identity_column_name string,
	in_scope_check object_model.Issue_types,
	cell_syntactic_check_results []object_model.Regex_check_results,
	regex_check_issue_result *object_model.Regex_check_results,
	cell_syntactic_check_issue_transactions [][]interface{}) ([]object_model.Regex_check_results, [][]interface{}) {

	var cell_syntactic_check_issue_transaction []interface{}

	if regex_check_issue_result != nil {

		fmt.Printf(
			"\nprocessing issues...\n")

		cell_syntactic_check_issue_transaction =
			cell_fixer.Generate_issue_transaction(
				*regex_check_issue_result)

		cell_syntactic_check_issue_transactions =
			aggregate_syntactic_check_issue_transactions(
				in_scope_identified_cell,
				identity_column_name,
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
	identity_column_name string,
	in_scope_check object_model.Issue_types,
	cell_syntactic_check_result_transaction []interface{},
	cell_sytactic_check_result_transaction_set [][]interface{}) [][]interface{} {

	if cell_syntactic_check_result_transaction != nil {
		//if issues are found append information to the transaction
		cell_syntactic_check_result_transaction =
			append(
				cell_syntactic_check_result_transaction,
				in_scope_check.Issue_type_uuid,
				in_scope_identified_cell[identity_column_name]) //TODO - Stage 2 - replace with human readable column name

		cell_sytactic_check_result_transaction_set = //append to transaction register
			append(cell_sytactic_check_result_transaction_set,
				cell_syntactic_check_result_transaction)

		fmt.Printf(
			"\n\ncell_sytactic_check_result_transaction_set:%v\n\n\n",
			cell_sytactic_check_result_transaction_set)
	}

	return cell_sytactic_check_result_transaction_set
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

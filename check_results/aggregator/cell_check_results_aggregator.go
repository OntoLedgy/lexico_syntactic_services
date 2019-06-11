package aggregator

import (
	"fmt"
	"syntactic_checker/cell_checkers"
	"syntactic_checker/cell_checkers/regex_checkers"
	"syntactic_checker/cell_fixer"
)

func Aggregate_syntactic_check_issues(
	cell_checker cell_checkers.CellCheckers,
	regex_check_issue_result *regex_checkers.RegexCheckResults,
	cell_syntactic_check_issue_transactions [][]interface{}) [][]interface{} {

	var cell_syntactic_check_issue_transaction []interface{}

	if regex_check_issue_result != nil {

		fmt.Printf(
			"\nprocessing issues...\n")

		cell_syntactic_check_issue_transaction =
			cell_fixer.Generate_issue_transaction(
				*regex_check_issue_result)

		cell_syntactic_check_issue_transactions =
			aggregate_syntactic_check_issue_transactions(
				cell_checker,
				cell_syntactic_check_issue_transaction,
				cell_syntactic_check_issue_transactions)

	}

	cell_syntactic_check_issue_transaction = nil

	return cell_syntactic_check_issue_transactions
}

func aggregate_syntactic_check_issue_transactions(
	cell_checker cell_checkers.CellCheckers,
	cell_syntactic_check_issue_transaction []interface{},
	cell_syntactic_check_issue_transactions [][]interface{}) [][]interface{} {

	if cell_syntactic_check_issue_transaction != nil {

		cell_syntactic_check_issue_transaction =
			append(
				cell_syntactic_check_issue_transaction,
				cell_checker.Issue_type.Issue_type_uuid,
				cell_checker.Cell.Cell_identifier)

		cell_syntactic_check_issue_transactions =
			append(cell_syntactic_check_issue_transactions,
				cell_syntactic_check_issue_transaction)

		fmt.Printf(
			"\n\ncell_syntactic_check_issue_transactions:%v\n\n\n",
			cell_syntactic_check_issue_transactions)
	}

	return cell_syntactic_check_issue_transactions
}

//#TODO - Deprecate

/*

func aggregate_syntactic_check_issue_result_transactions(
	regex_check_issue_result regex_checkers.RegexCheckResults,
	cell_syntactic_check_results []regex_checkers.RegexCheckResults) []regex_checkers.RegexCheckResults {

	cell_syntactic_check_results = append(
		cell_syntactic_check_results,
		regex_check_issue_result)

	return cell_syntactic_check_results

}


func generate_cell_syntactic_check_fix(
	cell_syntactic_check_result regex_checkers.RegexCheckResults,
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
	cell_syntactic_check_results []regex_checkers.RegexCheckResults) []interface{} {

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
}*/

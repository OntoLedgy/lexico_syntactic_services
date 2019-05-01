package results_finaliser

import (
	"fmt"
	"storage/slices"
	"strconv"
)

func Prepare_syntactic_checks_results_transactions(
	syntactic_checks_results_transactions [][]interface{},
	cells_syntactic_check_fix_transactions [][]interface{},
	column_uuid string) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	if syntactic_checks_results_transactions != nil {

		syntactic_checks_results_transaction_map =
			prepare_syntactic_check_results_transactions(
				syntactic_checks_results_transactions,
				cells_syntactic_check_fix_transactions,
				column_uuid)

	} else {

		syntactic_checks_results_transaction_map =
			nil
	}

	return syntactic_checks_results_transaction_map
}

func prepare_syntactic_check_results_transactions(
	check_transactions [][]interface{},
	cells_syntactic_check_fix_transactions [][]interface{},
	column_uuid string) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	syntactic_checks_results_transactions_string :=
		storage.Convert_2d_interface_to_string(
			check_transactions)

	cells_syntactic_check_fix_transactions_string :=
		storage.Convert_2d_interface_to_string(
			cells_syntactic_check_fix_transactions)

	syntactic_checks_results_transaction_map["syntactic_check_fix_transactions_set"] =
		prepare_syntactic_check_fixes(
			cells_syntactic_check_fix_transactions_string,
			column_uuid)

	syntactic_checks_results_transaction_map["syntactic_check_issues_set"] =
		prepare_syntactic_checks_issues(
			syntactic_checks_results_transactions_string,
			column_uuid)

	syntactic_checks_results_transaction_map["syntactic_check_issue_parameters_set"] =
		prepare_syntactic_checks_issue_parameters(
			syntactic_checks_results_transactions_string,
			column_uuid)

	return syntactic_checks_results_transaction_map

}

func prepare_syntactic_check_fixes(
	syntactic_check_result_transactions [][]string,
	column_uuid string) [][]string {

	var syntactic_check_fix_transactions_set [][]string

	syntactic_check_fixes_transactions :=
		syntactic_check_result_transactions

	for _, syntactic_check_fix := range syntactic_check_fixes_transactions {

		syntactic_check_fix_transactions_set =
			append(
				syntactic_check_fix_transactions_set,
				syntactic_check_fix)
	}

	return syntactic_check_fix_transactions_set
}

func prepare_syntactic_checks_issues(
	syntactic_check_result_transactions [][]string,
	column_uuid string) [][]string {

	var syntactic_check_issues_set [][]string

	columns_to_extract := []int{0, 4, 5} //get 0 - check_uuid, 4 - check_type_uuid, 5 - row_id

	syntactic_check_issues_transactions :=
		storage.Extract_columns_from_2d_slices(
			syntactic_check_result_transactions,
			columns_to_extract)

	//syntactic_check_issues_transactions =
	//	storage.Add_single_value_column_to_2d_slice(
	//		syntactic_check_issues_transactions,
	//		column_uuid)

	for _, syntactic_check_issue := range syntactic_check_issues_transactions {

		syntactic_check_issues_set =
			append(
				syntactic_check_issues_set,
				syntactic_check_issue)
	}
	fmt.Printf("issues:%s", syntactic_check_issues_set)

	return syntactic_check_issues_set

}

func prepare_syntactic_checks_issue_parameters(
	syntactic_check_result_transactions [][]string,
	column_uuid string) [][]string {

	var syntactic_check_issue_parameters_set [][]string

	var syntactic_check_issue_parameters [][]string

	parameter_columns := []int{5, 1, 2, 3}

	for _, syntactic_check_issue := range syntactic_check_result_transactions {

		syntactic_check_issue_parameters =
			get_syntactic_check_issue_parameters(
				syntactic_check_issue,
				parameter_columns)

		syntactic_check_issue_parameters_set =
			append(
				syntactic_check_issue_parameters_set,
				syntactic_check_issue_parameters...)
	}

	return syntactic_check_issue_parameters_set

}

func get_syntactic_check_issue_parameters(
	syntactic_check_issue []string,
	parameter_columns []int) [][]string {

	var syntactic_check_issue_parameters [][]string
	var syntactic_check_issue_parameter []string

	for parameter_index, parameter_column := range parameter_columns {

		syntactic_check_issue_parameter =
			append(syntactic_check_issue_parameter,
				syntactic_check_issue[0],
				strconv.Itoa(parameter_index+1),
				syntactic_check_issue[parameter_column])

		syntactic_check_issue_parameters =
			append(
				syntactic_check_issue_parameters,
				syntactic_check_issue_parameter)

		syntactic_check_issue_parameter = nil

	}

	return syntactic_check_issue_parameters
}

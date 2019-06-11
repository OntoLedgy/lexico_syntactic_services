package finaliser

import (
	"storage/slices"
	"strconv"
)

func Prepare_syntactic_checks_results_transactions(
	syntactic_checks_results_transactions [][]interface{},
	cells_syntactic_check_fix_transactions [][]interface{}) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	if syntactic_checks_results_transactions != nil {

		syntactic_checks_results_transaction_map =
			prepare_syntactic_check_results_transactions(
				syntactic_checks_results_transactions,
				cells_syntactic_check_fix_transactions)

	} else {

		syntactic_checks_results_transaction_map =
			nil
	}

	return syntactic_checks_results_transaction_map
}

func prepare_syntactic_check_results_transactions(
	check_transactions [][]interface{},
	cells_syntactic_check_fix_transactions [][]interface{}) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	syntactic_checks_results_transactions_string :=
		storage.Convert_2d_interface_to_string(
			check_transactions)

	cells_syntactic_check_fix_transactions_string :=
		storage.Convert_2d_interface_to_string(
			cells_syntactic_check_fix_transactions)

	syntactic_checks_results_transaction_map["syntactic_check_fix_transactions_set"] =
		cells_syntactic_check_fix_transactions_string

	syntactic_checks_results_transaction_map["syntactic_check_issues_set"] =
		prepare_syntactic_check_issues(
			syntactic_checks_results_transactions_string)

	syntactic_checks_results_transaction_map["syntactic_check_issue_parameters_set"] =
		prepare_syntactic_checks_issue_parameters(
			syntactic_checks_results_transactions_string)

	return syntactic_checks_results_transaction_map

}

func prepare_syntactic_check_issues(
	syntactic_check_result_transactions [][]string) [][]string {

	var syntactic_check_issues_set [][]string

	issue_columns := []int{0, 4, 5}

	syntactic_check_issues_set = storage.Extract_columns_from_2d_slices(syntactic_check_result_transactions, issue_columns)

	return syntactic_check_issues_set

}

func prepare_syntactic_checks_issue_parameters(
	syntactic_check_result_transactions [][]string) [][]string {

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

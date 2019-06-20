package results_processor

import (
	"syntactic_checker/code/object_model/transactions"
)

func Prepare_syntactic_checks_results_transactions(
	column_check_issues transactions.IssuesTransactions,
	column_check_fixes transactions.FixTransactions) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	if column_check_issues.Issues != nil {

		syntactic_checks_results_transaction_map =
			prepare_syntactic_check_results_transactions(
				column_check_fixes,
				column_check_issues)

	} else {

		syntactic_checks_results_transaction_map =
			nil
	}

	return syntactic_checks_results_transaction_map
}

func prepare_syntactic_check_results_transactions(
	cell_list_fix_transactions transactions.FixTransactions,
	cell_list_issues_transactions transactions.IssuesTransactions) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	syntactic_checks_results_transaction_map["syntactic_check_fix_transactions_set"] =
		prepare_syntactic_check_fixes(
			cell_list_fix_transactions)

	syntactic_checks_results_transaction_map["syntactic_check_issues_set"] =
		prepare_syntactic_check_issues(
			cell_list_issues_transactions)

	syntactic_checks_results_transaction_map["syntactic_check_issue_parameters_set"] =
		prepare_syntactic_checks_issue_parameters(
			cell_list_issues_transactions)

	return syntactic_checks_results_transaction_map

}

func prepare_syntactic_check_fixes(
	cell_list_fix_transactions transactions.FixTransactions) [][]string {

	fix_transaction_string := make([][]string, len(cell_list_fix_transactions.Fixes))

	for index, fix := range cell_list_fix_transactions.Fixes {

		fix_transaction_string[index] = make([]string, 5)

		fix_transaction_string[index][0] = fix.Cell.Cell_value
		fix_transaction_string[index][1] = fix.String_edit_history.Get_marked_string()
		fix_transaction_string[index][2] = fix.String_edit_history.Get_modified_string()
		fix_transaction_string[index][3] = fix.Object_uuid.String()
		fix_transaction_string[index][4] = fix.Cell.Cell_identifier

	}

	return fix_transaction_string
}

func prepare_syntactic_check_issues(
	cell_list_issues_transactions transactions.IssuesTransactions) [][]string {

	syntactic_check_issues_set := make([][]string, len(cell_list_issues_transactions.Issues))

	for index, issue := range cell_list_issues_transactions.Issues {

		syntactic_check_issues_set[index] = make([]string, 3)

		syntactic_check_issues_set[index][0] = issue.Objects.Object_uuid.String()
		syntactic_check_issues_set[index][1] = issue.Issue_type.Issue_type_uuid
		syntactic_check_issues_set[index][2] = issue.Cell.Cell_identifier

	}

	return syntactic_check_issues_set
}

func prepare_syntactic_checks_issue_parameters(

	cell_list_issues_transactions transactions.IssuesTransactions) [][]string {

	var syntactic_check_issue_parameters_set [][]string

	syntactic_check_issue_parameters_set = make([][]string, len(cell_list_issues_transactions.Issues)*4)

	for index, issue := range cell_list_issues_transactions.Issues {

		//TODO - Stage 2 - tidy up this algorithm

		index = index * 4

		syntactic_check_issue_parameters_set[index] = make([]string, 3)
		syntactic_check_issue_parameters_set[index][0] = issue.Object_uuid.String()
		syntactic_check_issue_parameters_set[index][1] = "1"
		syntactic_check_issue_parameters_set[index][2] = issue.Cell.Cell_identifier

		syntactic_check_issue_parameters_set[index+1] = make([]string, 3)
		syntactic_check_issue_parameters_set[index+1][0] = issue.Object_uuid.String()
		syntactic_check_issue_parameters_set[index+1][1] = "2"
		syntactic_check_issue_parameters_set[index+1][2] = issue.Cell.Cell_value

		syntactic_check_issue_parameters_set[index+2] = make([]string, 3)
		syntactic_check_issue_parameters_set[index+2][0] = issue.Object_uuid.String()
		syntactic_check_issue_parameters_set[index+2][1] = "3"
		syntactic_check_issue_parameters_set[index+2][2] = issue.Marked_cell_value

		syntactic_check_issue_parameters_set[index+3] = make([]string, 3)
		syntactic_check_issue_parameters_set[index+3][0] = issue.Object_uuid.String()
		syntactic_check_issue_parameters_set[index+3][1] = "4"
		syntactic_check_issue_parameters_set[index+3][2] = issue.Modified_cell_value

	}
	return syntactic_check_issue_parameters_set

}

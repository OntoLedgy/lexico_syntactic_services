package results_processors

import (
	"syntactic_checker/code/object_model/service_results"
)

func prepare_syntactic_checks_results_transactions(
	cell_list_checks_result service_results.CellListChecksResults) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	there_are_issues :=
		cell_list_checks_result.Cell_list_checks_results != nil

	if there_are_issues {

		syntactic_checks_results_transaction_map =
			prepare_syntactic_check_results_transactions(
				cell_list_checks_result)

	} else {

		syntactic_checks_results_transaction_map =
			nil
	}

	return syntactic_checks_results_transaction_map
}

func prepare_syntactic_check_results_transactions(
	cell_list_checks_result service_results.CellListChecksResults) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	syntactic_checks_results_transaction_map["syntactic_check_fix_transactions_set"] =
		prepare_syntactic_check_fixes(
			cell_list_checks_result)

	syntactic_checks_results_transaction_map["syntactic_check_issues_set"] =
		prepare_syntactic_check_issues(
			cell_list_checks_result)

	syntactic_checks_results_transaction_map["syntactic_check_issue_parameters_set"] =
		prepare_syntactic_checks_issue_parameters(
			cell_list_checks_result)

	return syntactic_checks_results_transaction_map

}

func prepare_syntactic_check_fixes(
	cell_list_checks_result service_results.CellListChecksResults) [][]string {

	cell_list_checks_result_set := cell_list_checks_result.Cell_list_checks_results

	var fix_transaction_string [][]string

	for _, cell_checks_result := range cell_list_checks_result_set {

		cell_fix_slice := make([]string, 5)

		cell_fix_slice[0] = cell_checks_result.In_scope_cell.Cell_value
		cell_fix_slice[1] = cell_checks_result.Cell_checks_fix.Cell_value_edit_history.Get_marked_string()
		cell_fix_slice[2] = cell_checks_result.Cell_checks_fix.Cell_value_edit_history.Get_modified_string()
		cell_fix_slice[3] = cell_checks_result.Cell_checks_fix.Object_uuid.String()
		cell_fix_slice[4] = cell_checks_result.In_scope_cell.Cell_identifier

		fix_transaction_string = append(fix_transaction_string, cell_fix_slice)

	}

	return fix_transaction_string
}

func prepare_syntactic_check_issues(
	cell_list_checks_result service_results.CellListChecksResults) [][]string {

	var syntactic_check_issues_set [][]string

	cell_list_checks_result_set := cell_list_checks_result.Cell_list_checks_results

	for _, cell_checks_result := range cell_list_checks_result_set {

		for _, cell_issue := range cell_checks_result.Cell_checks_issues {

			cell_issue_slice := make([]string, 3)

			cell_issue_slice[0] = cell_issue.Object_uuid.String()
			cell_issue_slice[1] = cell_issue.Issue_type.Issue_type_uuid
			cell_issue_slice[2] = cell_checks_result.In_scope_cell.Cell_identifier
			syntactic_check_issues_set = append(syntactic_check_issues_set, cell_issue_slice)

		}

	}

	return syntactic_check_issues_set
}

func prepare_syntactic_checks_issue_parameters(
	cell_list_checks_result service_results.CellListChecksResults) [][]string {

	var syntactic_check_issue_parameters_set [][]string

	cell_list_checks_result_set := cell_list_checks_result.Cell_list_checks_results

	for _, cell_checks_result := range cell_list_checks_result_set {

		for _, cell_issue := range cell_checks_result.Cell_checks_issues {

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				cell_issue.Object_uuid.String(),
				"1",
				cell_checks_result.In_scope_cell.Cell_identifier,
				syntactic_check_issue_parameters_set)

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				cell_issue.Object_uuid.String(),
				"2",
				cell_checks_result.In_scope_cell.Cell_value,
				syntactic_check_issue_parameters_set)

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				cell_issue.Object_uuid.String(),
				"3",
				cell_issue.Cell_edit_history.Get_marked_string(),
				syntactic_check_issue_parameters_set)

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				cell_issue.Object_uuid.String(),
				"4",
				cell_issue.Cell_edit_history.Get_modified_string(),
				syntactic_check_issue_parameters_set)

		}

	}

	/*
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
			syntactic_check_issue_parameters_set[index+2][2] = issue.Cell_value_edit_history.Get_marked_string()

			syntactic_check_issue_parameters_set[index+3] = make([]string, 3)
			syntactic_check_issue_parameters_set[index+3][0] = issue.Object_uuid.String()
			syntactic_check_issue_parameters_set[index+3][1] = "4"
			syntactic_check_issue_parameters_set[index+3][2] = issue.Cell_value_edit_history.Get_modified_string()

		}
	*/
	return syntactic_check_issue_parameters_set

}

func append_issue_parameter_slice(
	cell_check_issue_uuid string,
	parameter_number string,
	parameter_value string,
	syntactic_check_issue_parameters_set [][]string) [][]string {

	cell_issue_parameter_slice := make([]string, 3)

	cell_issue_parameter_slice[0] = cell_check_issue_uuid // cell_issue.Object_uuid.String()
	cell_issue_parameter_slice[1] = parameter_number      //"1"
	cell_issue_parameter_slice[2] = parameter_value       //cell_checks_result.In_scope_cell.Cell_identifier

	syntactic_check_issue_parameters_set =
		append(
			syntactic_check_issue_parameters_set,
			cell_issue_parameter_slice)

	return syntactic_check_issue_parameters_set
}

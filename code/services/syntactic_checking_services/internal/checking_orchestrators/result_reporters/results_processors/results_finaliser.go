package results_processors

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

//TODO - Add type structure

func prepare_syntactic_checks_results_transactions(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	there_are_issues :=
		identified_string_list_checks_result.Identified_string_checks_results != nil

	if there_are_issues {

		syntactic_checks_results_transaction_map =
			prepare_syntactic_check_results_transactions(
				identified_string_list_checks_result)

	} else {

		syntactic_checks_results_transaction_map =
			nil
	}

	return syntactic_checks_results_transaction_map
}

func prepare_syntactic_check_results_transactions(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults) map[string][][]string {

	syntactic_checks_results_transaction_map :=
		make(map[string][][]string, 3)

	syntactic_checks_results_transaction_map["syntactic_check_fix_transactions_set"] =
		prepare_syntactic_check_fixes(
			identified_string_list_checks_result)

	syntactic_checks_results_transaction_map["syntactic_check_issues_set"] =
		prepare_syntactic_check_issues(
			identified_string_list_checks_result)

	syntactic_checks_results_transaction_map["syntactic_check_issue_parameters_set"] =
		prepare_syntactic_checks_issue_parameters(
			identified_string_list_checks_result)

	syntactic_checks_results_transaction_map["syntactic_check_issue_details_set"] =
		prepare_syntactic_checks_issue_details(
			identified_string_list_checks_result)

	return syntactic_checks_results_transaction_map

}

func prepare_syntactic_check_fixes(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults) [][]string {

	identified_string_list_checks_result_set := identified_string_list_checks_result.Identified_string_checks_results

	var fix_transaction_string [][]string

	for _, identified_string_checks_result := range identified_string_list_checks_result_set {

		if identified_string_checks_result.String_checks_result.String_checks_fix.String_value_edit_history != nil {
			identified_string_fix_slice := make([]string, 5)

			identified_string_fix_slice[0] = identified_string_checks_result.Identified_string.String_identified.String_value
			identified_string_fix_slice[1] = identified_string_checks_result.String_checks_result.String_checks_fix.String_value_edit_history.Get_marked_string()
			identified_string_fix_slice[2] = identified_string_checks_result.String_checks_result.String_checks_fix.String_value_edit_history.Get_modified_string()
			identified_string_fix_slice[3] = identified_string_checks_result.String_checks_result.String_checks_fix.Object_uuid.String()
			identified_string_fix_slice[4] = identified_string_checks_result.Identified_string.String_identifier

			fix_transaction_string = append(fix_transaction_string, identified_string_fix_slice)
		}
	}

	return fix_transaction_string
}

func prepare_syntactic_check_issues(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults) [][]string {

	var syntactic_check_issues_set [][]string

	identified_string_list_checks_result_set := identified_string_list_checks_result.Identified_string_checks_results

	for _, identified_string_checks_result := range identified_string_list_checks_result_set {

		for _, identified_string_issue := range identified_string_checks_result.String_checks_result.String_checks_issues_list.String_checks_issue_results {

			identified_string_issue_slice := make([]string, 3)

			identified_string_issue_slice[0] = identified_string_issue.String_checks_issue.Object_uuid.String()
			identified_string_issue_slice[1] = identified_string_issue.String_checks_issue.Issue_type.Issue_type_uuid
			identified_string_issue_slice[2] = identified_string_checks_result.Identified_string.String_identifier
			syntactic_check_issues_set = append(syntactic_check_issues_set, identified_string_issue_slice)

		}

	}

	return syntactic_check_issues_set
}

func prepare_syntactic_checks_issue_parameters(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults) [][]string {

	var syntactic_check_issue_parameters_set [][]string

	identified_string_list_checks_result_set := identified_string_list_checks_result.Identified_string_checks_results

	for _, identified_string_checks_result := range identified_string_list_checks_result_set {

		for _, identified_string_issue := range identified_string_checks_result.String_checks_result.String_checks_issues_list.String_checks_issue_results {

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				identified_string_issue.String_checks_issue.Object_uuid.String(),
				"1",
				identified_string_checks_result.Identified_string.String_identifier,
				syntactic_check_issue_parameters_set)

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				identified_string_issue.String_checks_issue.Object_uuid.String(),
				"2",
				identified_string_checks_result.Identified_string.String_identified.String_value,
				syntactic_check_issue_parameters_set)

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				identified_string_issue.String_checks_issue.Object_uuid.String(),
				"3",
				identified_string_issue.String_edit_history.Get_marked_string(),
				syntactic_check_issue_parameters_set)

			syntactic_check_issue_parameters_set = append_issue_parameter_slice(
				identified_string_issue.String_checks_issue.Object_uuid.String(),
				"4",
				identified_string_issue.String_edit_history.Get_modified_string(),
				syntactic_check_issue_parameters_set)

		}

	}

	return syntactic_check_issue_parameters_set

}

func append_issue_parameter_slice(
	identified_string_check_issue_uuid string,
	parameter_number string,
	parameter_value string,
	syntactic_check_issue_parameters_set [][]string) [][]string {

	identified_string_issue_parameter_slice := make([]string, 3)

	identified_string_issue_parameter_slice[0] = identified_string_check_issue_uuid
	identified_string_issue_parameter_slice[1] = parameter_number
	identified_string_issue_parameter_slice[2] = parameter_value

	syntactic_check_issue_parameters_set =
		append(
			syntactic_check_issue_parameters_set,
			identified_string_issue_parameter_slice)

	return syntactic_check_issue_parameters_set
}

func prepare_syntactic_checks_issue_details(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults) [][]string {

	var syntactic_check_issues_set [][]string

	identified_string_list_checks_result_set := identified_string_list_checks_result.Identified_string_checks_results

	for _, identified_string_checks_result := range identified_string_list_checks_result_set {

		for _, identified_string_issue := range identified_string_checks_result.String_checks_result.String_checks_issues_list.String_checks_issue_results {

			identified_string_issue_slice := make([]string, 5)

			identified_string_issue_slice[0] = identified_string_checks_result.Identified_string.String_identified.String_value
			identified_string_issue_slice[1] = identified_string_issue.String_edit_history.Get_marked_string()
			identified_string_issue_slice[2] = identified_string_issue.String_edit_history.Get_modified_string()
			identified_string_issue_slice[3] = identified_string_issue.String_checks_issue.Issue_type.Issue_type_uuid
			identified_string_issue_slice[4] = identified_string_checks_result.Identified_string.String_identifier

			syntactic_check_issues_set = append(syntactic_check_issues_set, identified_string_issue_slice)

		}

	}

	return syntactic_check_issues_set
}

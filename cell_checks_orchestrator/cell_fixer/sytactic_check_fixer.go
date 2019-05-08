package cell_fixer

import (
	"syntactic_checker/object_model"
)

func Generate_issue_transaction(
	regex_check_result object_model.Regex_check_results) []interface{} {

	var check_result_transaction []interface{}

	cell_value_marked :=
		Modify_string_by_index(
			regex_check_result.Original_string,
			regex_check_result.Mark_string,
			regex_check_result.Regex_match_indices)

	cell_value_modified :=
		Modify_string_by_index(
			regex_check_result.Original_string,
			regex_check_result.Replacement_string,
			regex_check_result.Regex_match_indices)

	check_result_transaction =
		append(check_result_transaction,
			regex_check_result.Check_uuids,
			regex_check_result.Original_string,
			cell_value_marked,
			cell_value_modified)

	return check_result_transaction
}

//		TODO - Stage 2 - 	separate replacement process from regex check, return sub_match_indicies
// 		for aggregatation first and then call modification function in one go.

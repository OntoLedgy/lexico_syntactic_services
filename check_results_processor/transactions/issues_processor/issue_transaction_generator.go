package issues_processor

import (
	"syntactic_checker/cell_checkers/regex_checkers"
	"syntactic_checker/helpers/cell_editor"
	"syntactic_checker/object_model/issues"
)

func Generate_issue_transaction(
	regex_check_result regex_checkers.RegexCheckResults,
	issue_type issues.IssueTypes) []interface{} {

	var check_result_transaction []interface{}

	cell_value_modified, cell_value_marked :=
		cell_editor.Edit_cell(
			regex_check_result.Original_string,
			issue_type,
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

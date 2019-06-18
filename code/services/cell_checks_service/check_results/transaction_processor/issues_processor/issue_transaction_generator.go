package issues_processor

import (
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_service/regex_checkers"
	"syntactic_checker/code/services/cell_checks_service/cell_editor"
)

func Generate_issue_transaction(
	regex_check_result regex_checkers.RegexCheckResults,
	issue_type issues.IssueTypes) ([]interface{}, issues.Issues) {

	var check_result_transaction []interface{}
	var cell_check_issue issues.Issues

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

	cell_check_issue.Object_uuid = cell_check_issue.Set_object_uuid()
	cell_check_issue.Marked_cell_value = cell_value_marked
	cell_check_issue.Modified_cell_value = cell_value_modified

	return check_result_transaction, cell_check_issue
}

//		TODO - Stage 2 - separate replacement process from regex check, return sub_match_indicies
// 		for aggregation first and then call modification function in one go.

package cell_checks

import (
	"database_manager/utils"
	"fmt"
	"syntactic_checker/cell_checkers"
	"syntactic_checker/cell_fixer"
	"syntactic_checker/object_model"
)

func Process_syntactic_check_fixes_for_cell(
	in_scope_syntactic_check_types []object_model.IssueTypes,
	in_scope_cell object_model.InScopeCell,
	cell_syntactic_check_issue_transactions [][]interface{}) []interface{} {

	var cell_syntactic_check_aggregated_fixes_transaction []interface{}
	var interim_cell_modified, interim_cell_marked object_model.InScopeCell

	if cell_syntactic_check_issue_transactions != nil {

		fmt.Printf(
			"\nProcessing fixes...\n")

		fix_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		interim_cell_modified = in_scope_cell
		interim_cell_marked = in_scope_cell

		for _, in_scope_syntactic_check_type := range in_scope_syntactic_check_types {

			cell_checker := cell_checkers.
				Create_cell_checker(
					interim_cell_modified,
					in_scope_syntactic_check_type)

			cell_syntactic_check_issue_result :=
				cell_checker.
					CheckCell()

			// TODO - Stage 2 - move this out to String Editor function.
			// TODO - Stage 3 - Add link to String Editor here...

			if cell_syntactic_check_issue_result != nil {

				interim_cell_modified.Cell_value =
					cell_fixer.Modify_string_by_index(
						interim_cell_modified.Cell_value,
						cell_syntactic_check_issue_result.Replacement_string,
						cell_syntactic_check_issue_result.Regex_match_indices)

				interim_cell_marked.Cell_value =
					cell_fixer.Modify_string_by_index(
						interim_cell_marked.Cell_value,
						cell_syntactic_check_issue_result.Mark_string,
						cell_syntactic_check_issue_result.Regex_match_indices)

			}

		}

		cell_syntactic_check_aggregated_fixes_transaction =
			append(
				cell_syntactic_check_aggregated_fixes_transaction,
				in_scope_cell.Cell_value,
				interim_cell_marked.Cell_value,
				interim_cell_modified.Cell_value,
				fix_uuid.String(),
				in_scope_cell.Cell_identifier)

		fmt.Printf("\nCell fix transaction: %v \n", cell_syntactic_check_aggregated_fixes_transaction)

	}

	return cell_syntactic_check_aggregated_fixes_transaction
}

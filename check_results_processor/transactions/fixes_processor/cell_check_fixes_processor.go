package fixes_processor

import (
	"database_manager/utils"
	"fmt"
	"syntactic_checker/object_model"
	"syntactic_checker/object_model/issues"
)

func Process_cell_check_fixes(
	issue_types []issues.IssueTypes,
	in_scope_cell object_model.InScopeCell,
	cell_syntactic_check_issue_transactions [][]interface{}) []interface{} {

	var cell_syntactic_check_aggregated_fixes_transaction []interface{}

	if cell_syntactic_check_issue_transactions != nil {

		fmt.Printf(
			"\nProcessing fixes...\n")

		fix_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		interim_cell_modified, interim_cell_marked :=
			Generate_fix_transaction(
				in_scope_cell,
				issue_types)

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

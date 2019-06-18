package fixes_processor

import (
	"database_manager/utils"
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
)

func Get_cell_check_fix(
	issue_types []issues.IssueTypes,
	in_scope_cell object_model.InScopeCells,
	cell_syntactic_check_issue_transactions [][]interface{}) ([]interface{}, fixes.Fixes) {

	var cell_check_fixes_transaction []interface{}
	var cell_check_fix fixes.Fixes

	if cell_syntactic_check_issue_transactions != nil {

		fmt.Printf(
			"\nProcessing fixes...\n")

		fix_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		interim_cell_modified, interim_cell_marked, cell_check_fix :=
			Generate_fix_transaction(
				in_scope_cell,
				issue_types)

		cell_check_fixes_transaction =
			append(
				cell_check_fixes_transaction,
				in_scope_cell.Cell_value,
				interim_cell_marked.Cell_value,
				interim_cell_modified.Cell_value,
				fix_uuid.String(),
				in_scope_cell.Cell_identifier)

		fmt.Printf("\nCell fix transaction: %v \n", cell_check_fixes_transaction, cell_check_fix)

		return cell_check_fixes_transaction, cell_check_fix

	}

	return cell_check_fixes_transaction, cell_check_fix
}

package fix_processors

import (
	"fmt"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
)

type fixProcessors struct {
	issue_types     []issues.IssueTypes
	in_scope_cell   cells.Cells
	cell_checks_fix fixes.Fixes
}

func (fix_processor *fixProcessors) Get_cell_check_fix(
	issue_types []issues.IssueTypes,
	in_scope_cell cells.Cells) fixes.Fixes {

	fmt.Printf(
		"\nProcessing fixes for %s...\n",
		in_scope_cell)

	cell_check_fix :=
		Generate_fix_transaction(
			in_scope_cell,
			issue_types)

	fmt.Printf(
		"\nCell fix transaction: %v \n",
		cell_check_fix)

	return cell_check_fix

}

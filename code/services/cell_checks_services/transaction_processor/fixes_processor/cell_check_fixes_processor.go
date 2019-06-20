package fixes_processor

import (
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
)

func Get_cell_check_fix(
	issue_types []issues.IssueTypes,
	in_scope_cell object_model.Cells,
	cell_checks_issues []issues.Issues) *fixes.Fixes {

	if cell_checks_issues != nil {

		fmt.Printf(
			"\nProcessing fixes...\n")

		cell_check_fix :=
			Generate_fix_transaction(
				in_scope_cell,
				issue_types)

		fmt.Printf(
			"\nCell fix transaction: %v \n",
			cell_check_fix)

		return &cell_check_fix

	}

	return nil
}

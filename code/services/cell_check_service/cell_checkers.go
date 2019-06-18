package cell_check_service

import (
	"fmt"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_service/regex_checkers"
)

type CellCheckers struct {
	Issue_type issues.IssueTypes
	Cell       object_model.InScopeCells
}

func (cell_checker CellCheckers) CheckCell() *regex_checkers.RegexCheckResults {

	var regex_check_result *regex_checkers.RegexCheckResults

	if cell_checker.Cell.Cell_value != "" {

		regex_check_result = //TODO - Stage 3 - add switch to include non-regex in_scope_check types.
			regex_checkers.
				Process_regex_check(
					cell_checker.Issue_type.Issue_check_regex,
					cell_checker.Cell.Cell_value)

	} else {
		fmt.Printf(
			"\nWARNING: Cell value for row_id%s is null\n",
			cell_checker.Cell.Cell_value)

	}
	return regex_check_result

}

//TODO - add constructor

func (cell_checker CellCheckers) construct() {

}

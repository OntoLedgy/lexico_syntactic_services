package cell_checkers

import (
	"fmt"
	"syntactic_checker/cell_checkers/regex_checkers"
	"syntactic_checker/object_model"
)

type CellCheckers struct {
	Issue_type object_model.IssueTypes
	Cell       object_model.InScopeCell
}

func (cell_checker CellCheckers) CheckCell() *regex_checkers.RegexCheckResults {

	var regex_check_result *regex_checkers.RegexCheckResults

	if cell_checker.Cell.Cell_value != "" {

		regex_check_result = //TODO - Stage 3 - add switch to include non-regex in_scope_check types.
			regex_checkers.
				Process_regex_check( //TODO - Stage 2 - replace the check interface with the check type object and pass it through
					cell_checker.Issue_type.Issue_check_regex,
					cell_checker.Cell.Cell_value,
					cell_checker.Issue_type.Issue_check_replacement_string)

	} else {
		fmt.Printf(
			"\nWARNING: Cell value for row_id%s is null\n",
			cell_checker.Cell.Cell_value)

	}
	return regex_check_result

}

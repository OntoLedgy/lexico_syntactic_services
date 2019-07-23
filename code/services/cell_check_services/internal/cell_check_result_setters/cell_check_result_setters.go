package cell_check_result_setters

import (
	"syntactic_checker/code/services/cell_check_services/contract"
)

type CellCheckResultSetters struct {
	contract.ICellCheckServices
}

func (
	cell_check_results_setter *CellCheckResultSetters) Set_cell_check_result() {

	cell_check_parameter :=
		cell_check_results_setter.
			Get_check_parameter()

	cell_value :=
		cell_check_parameter.
			In_scope_cell.
			Cell_value

	check_regex :=
		cell_check_parameter.
			In_scope_issue_type.
			Issue_check_regex

	//TODO - Stage 3 - Add other check types (non - regex)

	cell_check_results_setter.
		Set_cell_check_result_value(
			Get_regex_check_result(
				check_regex,
				cell_value))

}

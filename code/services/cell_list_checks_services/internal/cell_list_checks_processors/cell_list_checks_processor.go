package cell_list_checks_processors

import (
	"fmt"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/services/cell_list_checks_services/contract"
)

type CellListChecksProcessors struct {
	contract.ICellListChecksServices
}

func (cell_list_checks_processor *CellListChecksProcessors) Process_cell_list_for_cell_checks() {

	cell_list_checks_processor.
		iterate_cell_list_for_cell_checks()

}

func (cell_list_checks_processor *CellListChecksProcessors) iterate_cell_list_for_cell_checks() {

	cell_list_checks_parameter :=
		cell_list_checks_processor.
			Get_cell_list_checks_parameter()

	fmt.Printf("Processing checks for cell list: \n",
		cell_list_checks_parameter.
			List_of_in_scope_issue_types)

	in_scope_cells :=
		cell_list_checks_parameter.
			List_of_in_scope_cells.
			Cells

	for _, in_scope_cell := range in_scope_cells {

		cell_list_checks_processor.
			process_and_set_cell_checks_result(
				in_scope_cell)

	}
}

func (
	cell_list_checks_processor *CellListChecksProcessors) process_and_set_cell_checks_result(
	in_scope_cell cells.Cells) {

	cell_list_checks_parameter :=
		cell_list_checks_processor.
			Get_cell_list_checks_parameter()

	issue_types :=
		cell_list_checks_parameter.
			List_of_in_scope_issue_types

	cell_checks_result :=
		Strip_cell_identifier_and_run_string_checks(
			in_scope_cell,
			issue_types)

	cell_list_checks_processor.
		Set_cell_checks_result(
			in_scope_cell,
			cell_checks_result)
}

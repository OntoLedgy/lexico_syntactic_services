package cell_list_checks_processors

import "syntactic_checker/code/services/cell_list_checks_services/contract"

func Create(cell_list_checks_service contract.ICellListChecksServices) *CellListChecksProcessors {

	cell_list_checks_processor :=
		new(
			CellListChecksProcessors)

	cell_list_checks_processor.
		ICellListChecksServices =
		cell_list_checks_service

	return cell_list_checks_processor

}

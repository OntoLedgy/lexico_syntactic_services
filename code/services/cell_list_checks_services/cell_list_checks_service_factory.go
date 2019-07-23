package cell_list_checks_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_list_checks_services/contract"
	"syntactic_checker/code/services/cell_list_checks_services/internal"
)

type CellListChecksServiceFactory struct{}

func (CellListChecksServiceFactory) Create(
	cell_list_checks_parameter service_parameters.CellListChecksParameters) contract.ICellListChecksServices {

	cell_list_checks_service :=
		new(
			internal.CellListChecksServices)

	cell_list_checks_service.
		Cell_list_checks_parameter =
		cell_list_checks_parameter

	return cell_list_checks_service
}

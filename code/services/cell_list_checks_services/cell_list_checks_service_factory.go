package cell_list_checks_services

import (
	"syntactic_checker/code/object_model/service_parameters"
)

type CellListChecksServiceFactory struct{}

func (CellListChecksServiceFactory) Create(
	cell_list_checks_parameter service_parameters.CellListChecksParameters) ICellListChecksService {

	cell_list_checks_service :=
		new(
			cellListChecksService)

	cell_list_checks_service.
		Cell_list_checks_parameter =
		cell_list_checks_parameter

	return cell_list_checks_service
}

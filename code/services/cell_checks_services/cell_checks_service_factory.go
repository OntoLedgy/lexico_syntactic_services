package cell_checks_services

import (
	"syntactic_checker/code/object_model/service_parameters"
)

type CellChecksServiceFactory struct{}

func (
	CellChecksServiceFactory) Create(
	cell_checks_parameter service_parameters.CellChecksParameters) ICellChecksServices {

	cell_checks_service :=
		new(
			cellChecksService)

	cell_checks_service.
		Cell_checks_parameter =
		cell_checks_parameter

	return cell_checks_service
}

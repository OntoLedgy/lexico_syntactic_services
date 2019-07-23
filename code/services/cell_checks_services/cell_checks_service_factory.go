package cell_checks_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_checks_services/contract"
	"syntactic_checker/code/services/cell_checks_services/internal"
)

type CellChecksServiceFactory struct{}

func (
	CellChecksServiceFactory) Create(
	cell_checks_parameter service_parameters.CellChecksParameters) contract.ICellChecksServices {

	cell_checks_service :=
		new(
			internal.CellChecksService)

	cell_checks_service.
		Cell_checks_parameter =
		cell_checks_parameter

	return cell_checks_service
}

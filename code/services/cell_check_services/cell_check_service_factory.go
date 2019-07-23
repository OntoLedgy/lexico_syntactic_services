package cell_check_services

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/cell_check_services/contract"
	"syntactic_checker/code/services/cell_check_services/internal"
)

type CellCheckServiceFactory struct{}

func (
	CellCheckServiceFactory) Create(
	cell_check_parameter service_parameters.CellCheckParameters) contract.ICellCheckServices {

	cell_check_service :=
		new(
			internal.CellCheckServices)

	cell_check_service.
		Cell_check_parameter =
		cell_check_parameter

	return cell_check_service

}

package cell_check_services

import (
	"syntactic_checker/code/object_model/service_parameters"
)

type CellCheckServiceFactory struct{}

func (
	CellCheckServiceFactory) Create(
	cell_check_parameter service_parameters.CellCheckParameters) ICellCheckService {

	cell_check_service :=
		new(
			cellCheckService)

	cell_check_service.
		Cell_check_parameter =
		cell_check_parameter

	return cell_check_service

}

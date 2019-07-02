package cell_list_checks_services

import (
	"fmt"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
)

type cellListChecksService struct {
	Cell_list_checks_parameter service_parameters.CellListChecksParameters
	Cell_list_checks_result    service_results.CellListChecksResults
}

func (
	cell_list_checks_service *cellListChecksService) Set_syntactic_checks_results() {

	fmt.Println("\nRunning cell list checks service...")

	cell_list_checks_service.
		process_cell_list_for_cell_checks()

}

func (
	cell_list_checks_service *cellListChecksService) Get_cell_list_checks_result() service_results.CellListChecksResults {

	return cell_list_checks_service.Cell_list_checks_result
}

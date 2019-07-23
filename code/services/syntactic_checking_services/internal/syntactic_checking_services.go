package internal

import (
	"logger/goinggo_services"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters"
)

type SyntacticCheckingServices struct {
	Run_configuration         configuration_getters.RunConfigurations
	syntactic_checking_result service_results.CellListChecksResults //should this be wrapped into another structure?
	In_scope_cell_list        cells.ListOfCells
	Logger                    *goinggo_services.Logger //use global logging service
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Run_syntactic_checking_service() {

	syntactic_checking_service.
		Logger.
		Info(
			"Starting run checking service")

	syntactic_checking_service_orchestrator :=
		checking_orchestrators.
			Create(syntactic_checking_service)

	syntactic_checking_service_orchestrator.
		Orchestrate_syntactic_checking()

	syntactic_checking_service.
		Logger.
		Info(
			"Exiting run checking service")
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_run_configuration() configuration_getters.RunConfigurations {

	return syntactic_checking_service.Run_configuration
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_in_scope_cell_list() cells.ListOfCells {

	return syntactic_checking_service.In_scope_cell_list
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Set_syntactic_check_results(
	syntactic_checking_result service_results.CellListChecksResults) {

	syntactic_checking_service.
		syntactic_checking_result =
		syntactic_checking_result
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_syntactic_checking_result() service_results.CellListChecksResults {
	return syntactic_checking_service.syntactic_checking_result
}

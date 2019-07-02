package syntactic_checking_services

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/configuration_getters"
)

type syntacticCheckingServices struct {
	run_configuration         configuration_getters.RunConfigurations
	syntactic_checking_result service_results.CellListChecksResults //should this be wrapped into another structure?
	in_scope_cell_list        cells.ListOfCells
}

func (
	syntactic_checking_service *syntacticCheckingServices) Run_syntactic_checking_service() {

	//should this be through a factory?
	syntactic_checking_service_orchestrator :=
		new(
			syntacticCheckingServicesOrchestrators)

	syntactic_checking_service_orchestrator.
		syntacticCheckingServices =
		*syntactic_checking_service

	syntactic_checking_service_orchestrator.
		orchestrate_syntactic_checking()

}

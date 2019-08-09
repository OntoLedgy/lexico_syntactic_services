package internal

import (
	"logger/goinggo_services"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters/object_model"
)

type SyntacticCheckingServices struct {
	syntactic_checking_result service_results.IdentifiedStringListChecksResults //should this be wrapped into another structure?
	Run_configuration         object_model.RunConfigurations
	Identified_string_list    identified_strings.IdentifiedStringLists
	Logger                    *goinggo_services.Logger //use global logging service
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Run_syntactic_checking_service() {

	syntactic_checking_service.
		Logger.
		Info(
			"Starting run checking service")

	syntactic_checking_service_orchestrator_factory :=
		new(
			checking_orchestrators.
				SyntacticCheckingServiceOrchestratorFactory)

	syntactic_checking_service_orchestrator :=
		syntactic_checking_service_orchestrator_factory.
			Create(
				syntactic_checking_service)

	syntactic_checking_service_orchestrator.
		Orchestrate_syntactic_checking()

	syntactic_checking_service.
		Logger.
		Info(
			"Exiting run checking service")
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_run_configuration() object_model.RunConfigurations {

	return syntactic_checking_service.Run_configuration
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_identified_string_list() identified_strings.IdentifiedStringLists {

	return syntactic_checking_service.Identified_string_list
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Set_syntactic_check_results(
	syntactic_checking_result service_results.IdentifiedStringListChecksResults) {

	syntactic_checking_service.
		syntactic_checking_result =
		syntactic_checking_result
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_syntactic_checking_result() service_results.IdentifiedStringListChecksResults {
	return syntactic_checking_service.syntactic_checking_result
}

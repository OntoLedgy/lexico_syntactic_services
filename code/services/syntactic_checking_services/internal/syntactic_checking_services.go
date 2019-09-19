package internal

import (
	"logger/standard_global_logger"
	"syntactic_checker/code/object_model/configurations"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators"
)

type SyntacticCheckingServices struct {
	Syntactic_checking_service_data *configurations.SyntacticCheckingData
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Run_syntactic_checking_service() {

	standard_global_logger.
		Global_logger.
		Print(
			"Starting syntactic checking service")

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

	standard_global_logger.Global_logger.Print(
		"Exiting syntactic checking service")
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_run_configuration() *configurations.RunConfigurations {

	return syntactic_checking_service.Syntactic_checking_service_data.Run_configuration
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_identified_string_list() *identified_strings.IdentifiedStringLists {

	return syntactic_checking_service.Syntactic_checking_service_data.Identified_string_list
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Set_syntactic_check_results(
	syntactic_checking_result service_results.IdentifiedStringListChecksResults) {

	syntactic_checking_service.
		Syntactic_checking_service_data.
		Syntactic_checking_results =
		&syntactic_checking_result
}

func (
	syntactic_checking_service *SyntacticCheckingServices) Get_syntactic_checking_result() service_results.IdentifiedStringListChecksResults {

	return *syntactic_checking_service.Syntactic_checking_service_data.Syntactic_checking_results
}

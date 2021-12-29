package services

import (
	"fmt"
	"github.com/OntoLedgy/syntactic_checker/code/services/service_run_preparers"
	"github.com/OntoLedgy/syntactic_checker/code/services/syntactic_checking_services"
	"github.com/urfave/cli"
)

func Orchestrate_services(command_line_argument_context *cli.Context) {

	fmt.Print(
		"Starting syntactic checking service\n")

	configuration_file_path :=
		command_line_argument_context.Args().Get(0)

	service_run_preparer :=
		new(
			service_run_preparers.ServiceRunPreparers)

	service_run_data :=
		service_run_preparer.
			Get_service_run_data(
				configuration_file_path)

	syntactic_checking_service_factory :=
		new(
			syntactic_checking_services.
				SyntacticCheckingServiceFactory)

	syntactic_checking_service :=
		syntactic_checking_service_factory.
			Create(
				service_run_data)

	syntactic_checking_service.
		Run_syntactic_checking_service()

	fmt.Print(
		"Ending syntactic checking service")

	//standard_global_logger.End_logger()
}

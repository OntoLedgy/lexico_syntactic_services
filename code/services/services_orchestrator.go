package services

import (
	"fmt"
	"logger/standard_global_logger"
	"os"
	"syntactic_checker/code/services/service_run_preparers"
	"syntactic_checker/code/services/syntactic_checking_services"
)

func Orchestrate_services() {

	fmt.Print(
		"Starting syntactic checking service\n")

	configuration_file_path :=
		os.
			Args[1]

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

	standard_global_logger.
		End_logger()
}

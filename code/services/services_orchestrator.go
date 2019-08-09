package services

import (
	log "logger/goinggo_services"
	"os"
	"syntactic_checker/code/services/syntactic_checking_services"
)

func Orchestrate_services() {

	logger :=
		start_logger()

	logger.Info(
		"Starting syntactic checking service")

	configuration_file_path :=
		os.
			Args[1]

	syntactic_checking_service_factory :=
		new(
			syntactic_checking_services.
				SyntacticCheckingServiceFactory)

	syntactic_checking_service :=
		syntactic_checking_service_factory.
			Create(
				configuration_file_path,
				logger)

	syntactic_checking_service.
		Run_syntactic_checking_service()

	end_logger(
		logger)

}

//TODO - Stage 3 - move this out to logger service

func start_logger() *log.Logger {

	logger :=
		log.
			Create_logger(
				"./outputs/logs",
				log.Info_Level)

	logger.
		Started()

	return logger
}

func end_logger(
	logger *log.Logger) {

	logger.Info(
		"Ending syntactic checking service")

	logger.Completed()

	logger.Stop()

}

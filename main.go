package main

import (
	"fmt"
	"os"
	"syntactic_checker/code/services/syntactic_checking_services"
)

//TODO - Stage 2 - add commandline tools
//TODO - Stage 2 - record import dependencies
//TODO - Stage 2 - add logger and error handling

func main() {

	fmt.Println(
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
				configuration_file_path)

	syntactic_checking_service.
		Run_syntactic_checking_service()

	fmt.Println(
		"Ending syntactic checking service")
}

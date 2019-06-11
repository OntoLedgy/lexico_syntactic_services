package main

import (
	"fmt"
	"os"
	"syntactic_checker/orchestrators"
)

//TODO - Stage 2 - add commandline tools
//TODO - Stage 2 - add configuration_handler management tools
//TODO - Stage 2 - record import dependencies
//TODO - Stage 2 - add logger and error handling

func main() {

	fmt.Println(
		"Starting cell set syntactic check orchestrator")

	configuration_file_path := os.Args[1]

	orchestrators.
		Orchestrate_syntactic_checks(
			configuration_file_path)

	fmt.Println(
		"Ending cell set syntactic check orchestrator")
}

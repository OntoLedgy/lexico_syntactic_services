package main

import (
	"fmt"
	"syntactic_checker/cell_checks_orchestrator"
)

//TODO - Stage 2 - add commandline tools
//TODO - Stage 2 - add configuration_handler management tools
//TODO - Stage 2 - record import dependencies
//TODO - Stage 2 - add logger and error handling

func main() {

	fmt.Println(
		"Starting cell set syntactic check orchestrator")

	cell_checks_orchestrator.Orchestrate_csv_cell_syntactic_checks()

	fmt.Println(
		"Ending cell set syntactic check orchestrator")
}

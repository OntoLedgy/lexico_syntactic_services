package main

import (
	"fmt"
	"syntactic_checker/syntactic_check_cell_orchestrator"
)

//TODO - Stage 2 - add commandline tools
//TODO - Stage 2 - add configuration management tools
//TODO - Stage 2 - record import dependencies
//TODO - Stage 2 - add logger and error handling

func main() {

	fmt.Println(
		"Starting cell set syntactic check orchestrator")

	syntactic_check_cell_orchestrator.Orchestrate_csv_cell_syntactic_checks()

	fmt.Println(
		"Ending cell set syntactic check orchestrator")
}

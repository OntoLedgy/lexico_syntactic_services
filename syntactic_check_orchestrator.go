package main

import (
	"fmt"
	"syntactic_checker/configuration"
	"syntactic_checker/reporter"
	//database_checker "syntactic_checker/syntactic_check_processor/database_data_processor"
	"syntactic_checker/syntactic_check_processor"
)

func main() {
	//TODO - Stage 2 - add commandline tools
	//TODO - Stage 2 - add configuration management tools

	fmt.Println(
		"Starting cell set syntactic check orchestrator") //TODO - Stage 2 - to be move to logger

	configuration, in_scope_checks :=
		configuration.Get_configuration() //#TODO - Stage 1 - include the in_scope_checks to the configuration structure

		//csv file based in_scope_checks
		//if configuration.Csv_configuration.Csv_checks_required {

	//TODO - Stage 1 - move to Orchestrator

	check_result_transaction_set :=
		syntactic_check_processor.Orchestrate_csv_syntactic_checks(
			configuration.Csv_configuration,
			in_scope_checks)

	//TODO - Stage 1 - add reporter.Report_syntactic_check_issues(check_result_transaction_set)
	//TODO - Stage 1 - add reporter.Report_syntactic_check_issue_parameters()
	reporter.Report_syntactic_check_fixes(check_result_transaction_set)

	//}

	// database based in_scope_checks - IGNORE BEYOND THIS POINT
	/*
		if configuration.Database_configuration.Database_checks_required {

			configuration_database_filename :=
				configuration.Database_configuration.Database_file_name

			database_checker.Execute_database_syntactic_checks(
				configuration_database_filename)

		}
		//Process_character_distribution(csv_file_name, column_set)
	*/
}

//TODO - Stage 2 - record import dependencies

package main

import (
	"fmt"
	"storage/csv"
	"syntactic_checker/configuration"
	"syntactic_checker/syntactic_checks"
)

func main() {
	//#TODO add commandline tools

	configuration, checks := configuration.Get_configuration()

	//csv file based checks
	if configuration.Csv_configuration.Csv_checks_required {
		fmt.Println("Starting csv checks")
		column_set :=
			make(map[string][]int)

		check_columns := configuration.Csv_configuration.Column_set

		for check_column_name := range check_columns {

			column_set[check_column_name] =
				append(
					column_set[check_column_name],
					configuration.Csv_configuration.Identity_column_number,
					check_columns[check_column_name])

		}

		fmt.Println(column_set)

		fmt.Print("\nReading CSV Data\n")

		csv_file,
			csv_file_data := storage.Open_csv_file(
			configuration.Csv_configuration.Csv_file_name)

		csv_dataset := storage.Read_csv_to_slice(
			csv_file,
			csv_file_data,
			"")

		syntactic_checks.Process_csv_syntactic_check_columns(
			csv_dataset,
			column_set,
			checks)

	}

	// database based checks

	if configuration.Databse_configuration.Database_checks_required {

		configuration_database_filename :=
			configuration.Databse_configuration.Database_file_name

		syntactic_checks.Execute_database_syntactic_checks(
			configuration_database_filename)

	}
	//Process_character_distribution(csv_file_name, column_set)

}

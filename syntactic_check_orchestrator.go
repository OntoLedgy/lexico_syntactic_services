package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	configuration_structures "syntactic_checker/configuration"
	"syntactic_checker/syntactic_checks"
)

func main() {

	configuration := get_configuration()

	//csv file based checks
	if configuration.Csv_configuration.Csv_checks_required {

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

		syntactic_checks.Execute_csv_syntactic_checks(
			configuration.Csv_configuration.Csv_file_name,
			column_set)

	}
	//
	// database based checks

	if configuration.Databse_configuration.Database_checks_required {

		configuration_database_filename :=
			configuration.Databse_configuration.Database_file_name
		syntactic_checks.Execute_database_syntactic_checks(configuration_database_filename)

	} //Process_character_distribution(csv_file_name, column_set)

}

func get_configuration() *configuration_structures.Configuration {

	configuration_file, err := os.Open("configuration.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened configuration.json")
	// defer the closing of our jsonFile so that we can parse it later on

	defer configuration_file.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(configuration_file)

	// we initialize our Configuration
	var configuration configuration_structures.Configuration

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &configuration)

	return &configuration

}

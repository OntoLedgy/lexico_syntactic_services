package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//TODO - Stage 1 - Clean names
//TODO - Stage 2 - replace configuration management framework

func Get_configuration() (*Configuration, [][]interface{}) {

	configuration_file, err := os.Open(`configuration.json`)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened configuration.json")

	defer configuration_file.Close()

	// read opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(
		configuration_file)

	// initialize our Configuration
	var configuration Configuration

	// unmarshal our byteArray into our configuration structure
	json.Unmarshal(
		byteValue,
		&configuration)

	var check []interface{}
	var checks [][]interface{}

	in_scope_checks := configuration.Csv_configuration.Checks

	for _, in_scope_check := range in_scope_checks {

		check = append(check,
			in_scope_check.Issue_type_uuid,
			in_scope_check.Issue_type_name,
			in_scope_check.Issue_check_regex,
			in_scope_check.Issue_check_replacement_string)
		checks = append(checks, check)
		check = nil
	}

	return &configuration, checks
}

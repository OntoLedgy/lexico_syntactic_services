package configuration_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//TODO - Stage 1 - Clean names
//TODO - Stage 2 - replace configuration_handler management framework

func Get_configuration() *Configurations {

	var run_configuration Configurations

	run_configuration_byte_array :=
		Read_json_to_byte_array(
			`run_configuration.json`) //TODO - Stage 1 - add this as a commandline argument

	json.Unmarshal(
		run_configuration_byte_array,
		&run_configuration)

	return &run_configuration
}

//TODO - Stage 1 - move this out to external utility package
func Read_json_to_byte_array(json_file_name string) []byte {

	json_data, json_file_read_error :=
		os.Open(
			json_file_name)

	if json_file_read_error != nil {
		fmt.Println(
			json_file_read_error)
	}
	fmt.Println(
		"Successfully Opened run_configuration.json")

	defer json_data.Close()

	json_data_byte_array, _ :=
		ioutil.ReadAll(
			json_data)

	return json_data_byte_array
}

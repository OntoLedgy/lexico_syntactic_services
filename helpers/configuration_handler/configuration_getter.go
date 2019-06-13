package configuration_handler

import (
	"encoding/json"
	storage_json "storage/json"
)

//TODO - Stage 3 - replace configuration_handler management framework

func Get_configuration(configuration_file_path string) *Configurations {

	var run_configuration Configurations

	run_configuration_byte_array :=
		storage_json.Read_json_to_byte_array(
			configuration_file_path)

	json.Unmarshal(
		run_configuration_byte_array,
		&run_configuration)

	Modification_marker = run_configuration.Check_configuration.Modification_marker

	return &run_configuration
}

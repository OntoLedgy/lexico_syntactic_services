package configuration_getters

import (
	"encoding/json"
	storage_json "storage/json"
)

//TODO - Stage 3 - replace with configuration_getters management framework

type configurationGetters struct{}

func (configuraiton_getter *configurationGetters) Get_configuration(
	configuration_file_path string) *RunConfigurations {

	var run_configuration RunConfigurations

	run_configuration_byte_array :=
		storage_json.Read_json_to_byte_array(
			configuration_file_path)

	json.Unmarshal(
		run_configuration_byte_array,
		&run_configuration)

	Modification_marker =
		run_configuration.
			Check_configuration.
			Modification_marker

	return &run_configuration
}

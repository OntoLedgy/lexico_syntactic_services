package configuration_getters

import (
	"encoding/json"
	storage_json "storage/json"
	"syntactic_checker/code/object_model"
	configueation_object_model "syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters/object_model"
)

//TODO - Stage 3 - replace with configuration_getters management framework

type configurationGetters struct{}

func (configuraiton_getter *configurationGetters) Get_configuration(
	configuration_file_path string) *configueation_object_model.RunConfigurations {

	var run_configuration configueation_object_model.RunConfigurations

	run_configuration_byte_array :=
		storage_json.
			Read_json_to_byte_array(
				configuration_file_path)

	json.Unmarshal(
		run_configuration_byte_array,
		&run_configuration)

	object_model.Modification_marker =
		run_configuration.
			Check_configuration.
			Modification_marker

	return &run_configuration
}

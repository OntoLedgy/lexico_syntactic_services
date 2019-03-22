package configuration_handler

import "database_manager/utils"

type Csv_configuration struct {
	Csv_checks_required      bool          `json:"csv_checks_required"`
	Csv_file_name            string        `json:"csv_file_name"`
	Identity_column_position int           `json:"identity_column_number"` //TODO - Stage 1 - change json tag name to identity_column_position
	Check_column_name        string        `json:"check_column_name"`
	Check_column_uuid        string        `json:"check_column_uuid"`
	Issue_types              []Issue_types `json:"issue_types"`
}

type Column_details map[string]Columns

type Columns struct {
	Column_uuids string `json:"column_uuids"`
	Column_name  string `json:"column_name"`
	Column_index int    `json:"column_index"`
}

type Database_configuration struct {
	Database_checks_required bool   `json:"database_checks_required"`
	Database_file_name       string `json:"csv_file_name"`
}

type Configurations struct {
	Csv_configuration      Csv_configuration
	Database_configuration Database_configuration
}

type Issue_types struct {
	Issue_type_uuid                string `json:"issue_type_uuid"`
	Issue_type_name                string `json:"issue_type_name"`
	Issue_check_type               string `json:"issue_check_type"`         //TODO - Stage 1 - move checks out from issues
	Issue_check_regex              string `json:"check_regex"`              //TODO - Stage 1 - move regex checks from checks
	Issue_check_replacement_string string `json:"check_replacement_string"` //

}

//TODO - Stage 1 - integrate this into the issue type structure

type Checks struct {
	uuids       utils.UUID
	check_names string
	check_type  Check_Types
}

type Check_Types struct {
	uuids            utils.UUID
	check_type_names string
}

type Non_parameterised_checks struct {
	check                    Checks
	check_regex_string       string
	check_replacement_string string
}

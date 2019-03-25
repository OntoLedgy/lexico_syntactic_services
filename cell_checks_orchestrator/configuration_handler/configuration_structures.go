package configuration_handler

import "syntactic_checker/object_model"

type Configurations struct {
	Configuration_cotext string

	Csv_configuration      Csv_configuration
	Database_configuration Database_configuration
}

type Csv_configuration struct {
	Csv_checks_required      bool                       `json:"csv_checks_required"`
	Csv_file_name            string                     `json:"csv_file_name"`
	Identity_column_position int                        `json:"identity_column_number"` //TODO - Stage 1 - change json tag name to identity_column_position
	Check_column_name        string                     `json:"check_column_name"`
	Check_column_uuid        string                     `json:"check_column_uuid"`
	Issue_types              []object_model.Issue_types `json:"issue_types"`
}

type Database_configuration struct {
	Database_checks_required bool   `json:"database_checks_required"`
	Database_file_name       string `json:"csv_file_name"`
}

type Column_details map[string]Columns

type Columns struct {
	Column_uuids string `json:"column_uuids"`
	Column_name  string `json:"column_name"`
	Column_index int    `json:"column_index"`
}
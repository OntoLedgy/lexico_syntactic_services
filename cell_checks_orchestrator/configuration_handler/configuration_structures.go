package configuration_handler

import "syntactic_checker/object_model"

type Configurations struct {
	Configuration_cotext string

	Csv_configuration      Csv_configuration
	Database_configuration Database_configuration
	Output_configuration   Output_configurations
}

type Csv_configuration struct {
	Csv_checks_required  bool                       `json:"csv_checks_required"`
	Csv_file_name        string                     `json:"csv_file_name"`
	Identity_column_name string                     `json:"identity_column_name"`
	Check_column_name    string                     `json:"check_column_name"`
	Check_column_uuid    string                     `json:"check_column_uuid"`
	Issue_types          []object_model.Issue_types `json:"issue_types"`
}

type Output_configurations struct {
	Root_folder_path                           string `json:"root_folder_path"`
	Output_fixes_file_relative_path            string `json:"fixes_file_relative_path"`
	Output_issues_file_relative_path           string `json:"issues_file_relative_path"`
	Output_issue_parameters_file_relative_path string `json:"issue_parameters_file_relative_path"`
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

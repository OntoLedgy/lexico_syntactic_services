package configuration_getters

import "syntactic_checker/code/object_model/issues"

type CheckConfigurations struct {
	Checks_required      bool                `json:"csv_checks_required"`
	Input_csv_file_name  string              `json:"csv_file_name"`
	Identity_column_name string              `json:"identity_column_name"`
	Check_column_name    string              `json:"check_column_name"`
	Check_column_uuid    string              `json:"check_column_uuid"`
	Modification_marker  string              `json:"modification_marker"`
	Issue_types          []issues.IssueTypes `json:"issue_types"`
}

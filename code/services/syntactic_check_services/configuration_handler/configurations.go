package configuration_handler

import (
	"syntactic_checker/code/object_model/issues"
)

var Modification_marker string

type Configurations struct {
	Configuration_cotext string
	Check_configuration  Check_configuration
	Output_configuration Output_configurations
}

type Check_configuration struct {
	Csv_checks_required  bool                `json:"csv_checks_required"`
	Csv_file_name        string              `json:"csv_file_name"`
	Identity_column_name string              `json:"identity_column_name"`
	Check_column_name    string              `json:"check_column_name"`
	Check_column_uuid    string              `json:"check_column_uuid"`
	Modification_marker  string              `json:"modification_marker"`
	Issue_types          []issues.IssueTypes `json:"issue_types"`
}

type Output_configurations struct {
	Output_fixes_file_absolute_path            string `json:"output_fixes_file_absolute_path"`
	Output_issues_file_absolute_path           string `json:"output_issues_file_absolute_path"`
	Output_issue_parameters_file_absolute_path string `json:"output_issue_parameter_file_absolute_path"`
}

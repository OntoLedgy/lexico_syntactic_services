package configurations

type OutputConfigurations struct {
	Output_log_folder_absolute_path            string `json:"output_log_folder_absolute_path"`
	Output_fixes_file_absolute_path            string `json:"output_fixes_file_absolute_path"`
	Output_issues_file_absolute_path           string `json:"output_issues_file_absolute_path"`
	Output_issues_details_absolute_path        string `json:"output_issues_details_file_absolute_path"`
	Output_issue_parameters_file_absolute_path string `json:"output_issue_parameter_file_absolute_path"`
}

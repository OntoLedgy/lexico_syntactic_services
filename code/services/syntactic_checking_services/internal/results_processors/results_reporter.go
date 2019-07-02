package results_processors

import (
	"storage/csv"
)

func (
	results_processor *resultsProcessors) Report_syntactic_check_outputs() {

	//TODO - Split this method

	cell_list_checks_result :=
		results_processor.
			syntactic_checking_results

	syntactic_check_result_report :=
		prepare_syntactic_checks_results_transactions(
			cell_list_checks_result)

	output_configuration :=
		results_processor.output_configuration

	issues_file_name :=
		output_configuration.
			Output_issues_file_absolute_path

	issue_parameters_file_name :=
		output_configuration.
			Output_issue_parameters_file_absolute_path

	fixes_file_name :=
		output_configuration.
			Output_fixes_file_absolute_path

	syntactic_check_issues_set :=
		syntactic_check_result_report["syntactic_check_issues_set"]

	report_syntactic_check_issues(
		syntactic_check_issues_set,
		issues_file_name)

	report_syntactic_check_issue_parameters(
		syntactic_check_result_report["syntactic_check_issue_parameters_set"],
		issue_parameters_file_name)

	report_syntactic_check_fixes(
		syntactic_check_result_report["syntactic_check_fix_transactions_set"],
		fixes_file_name)
}

func report_syntactic_check_issues(
	issue_transactions [][]string,
	output_csv_filename string) {

	output_header := []string{
		"check_uuids",
		"check_type_uuids",
		"identifiers"}

	storage.Write_slice_with_header_to_csv(
		issue_transactions,
		output_header,
		output_csv_filename)

}

func report_syntactic_check_issue_parameters(
	issue_parameter_transactions [][]string,
	output_csv_filename string) {

	output_header := []string{
		"check_uuids",
		"parameter_sequence_number",
		"parameter_values"}

	storage.Write_slice_with_header_to_csv(
		issue_parameter_transactions,
		output_header,
		output_csv_filename)

}

func report_syntactic_check_fixes(
	fix_transactions [][]string,
	output_csv_filename string) {

	output_header := []string{
		"cell_values_original",
		"cell_values_marked",
		"cell_values_fixed",
		"fix_uuids",
		"identifiers"}

	storage.Write_slice_with_header_to_csv(
		fix_transactions,
		output_header,
		output_csv_filename)

}

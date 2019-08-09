package result_reporters

import (
	"storage/csv"
)

//TODO - Add type structure

func write_syntactic_checking_result_to_csvs(
	syntactic_check_result_report map[string][][]string,
	issues_file_name string,
	issue_parameters_file_name string,
	fixes_file_name string) {

	report_syntactic_check_issues(
		syntactic_check_result_report["syntactic_check_issues_set"],
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

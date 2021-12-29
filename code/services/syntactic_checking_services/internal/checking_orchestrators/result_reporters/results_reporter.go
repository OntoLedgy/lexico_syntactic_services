package result_reporters

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/disk/files"
	"github.com/OntoLedgy/storage_interop_services/code/services/documents/csv"
)

//TODO - Add type structure

func write_syntactic_checking_result_to_csvs(
	syntactic_check_result_report map[string][][]string,
	issues_file_name string,
	issue_parameters_file_name string,
	fixes_file_name string,
	issue_details_file_name string) {

	syntactic_check_issues_output_header := []string{
		"check_uuids",
		"check_type_uuids",
		"identifiers"}

	write_check_results(
		syntactic_check_result_report["syntactic_check_issues_set"],
		issues_file_name,
		syntactic_check_issues_output_header)

	syntactic_check_issue_parameters_output_header := []string{
		"check_uuids",
		"parameter_sequence_number",
		"parameter_values"}

	write_check_results(
		syntactic_check_result_report["syntactic_check_issue_parameters_set"],
		issue_parameters_file_name,
		syntactic_check_issue_parameters_output_header)

	syntactic_check_fixes_output_header := []string{
		"cell_values_original",
		"cell_values_marked",
		"cell_values_fixed",
		"fix_uuids",
		"identifiers"}

	write_check_results(
		syntactic_check_result_report["syntactic_check_fix_transactions_set"],
		fixes_file_name,
		syntactic_check_fixes_output_header)

	syntactic_check_issue_details_output_header := []string{
		"cell_values_original",
		"cell_values_marked",
		"cell_values_fixed",
		"check_type_uuids",
		"identifiers"}

	write_check_results(
		syntactic_check_result_report["syntactic_check_issue_details_set"],
		issue_details_file_name,
		syntactic_check_issue_details_output_header)

}

func write_check_results(
	issue_transactions [][]string,
	output_csv_filename string,
	output_header []string) {

	files.Delete_file_it_already_exists(
		output_csv_filename)

	csv.
		Write_slice_with_header_to_csv(
			issue_transactions,
			output_header,
			output_csv_filename)
}

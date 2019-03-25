package results_finaliser

import "storage/csv"

func Report_syntactic_check_outputs(
	syntactic_check_results map[string][][]string) {

	syntactic_check_issues_set :=
		syntactic_check_results["syntactic_check_issues_set"]

	Report_syntactic_check_issues(
		syntactic_check_issues_set) //TODO - Stage 1 - send just the dictionary to the called function.

	Report_syntactic_check_issue_parameters(
		syntactic_check_results["syntactic_check_issue_parameters_set"])

	Report_syntactic_check_fixes(
		syntactic_check_results["syntactic_check_fix_transactions_set"])
}

func Report_syntactic_check_issues(
	issue_transactions [][]string) {

	output_csv_filename := "outputs\\issues.csv" //TODO - Stage 2 - change to folder path (outputs/issues.csv)

	output_header := []string{
		"check_uuids",
		"check_type_uuids",
		"object_id",
		"column_uuids"}

	storage.Write_slice_with_header_to_csv(
		issue_transactions,
		output_header,
		output_csv_filename)

}

func Report_syntactic_check_issue_parameters(
	issue_parameter_transactions [][]string) {

	output_csv_filename := "outputs\\issue_parameters.csv"

	output_header := []string{
		"check_uuids",
		"parameter_sequence_number",
		"parameter_values"}

	storage.Write_slice_with_header_to_csv(
		issue_parameter_transactions,
		output_header,
		output_csv_filename)

}

func Report_syntactic_check_fixes(
	fix_transactions [][]string) {

	output_csv_filename := "outputs\\fixes.csv"

	output_header := []string{
		"issue_uuids",
		"cell_values_original",
		"cell_values_marked",
		"cell_values_modified",
		"check_type_uuids",
		"row_id",
		"column_uuids"}

	storage.Write_slice_with_header_to_csv(
		fix_transactions,
		output_header,
		output_csv_filename)

}

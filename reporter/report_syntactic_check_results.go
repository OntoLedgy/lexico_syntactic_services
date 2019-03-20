package reporter

import "storage/csv"

func Report_syntactic_check_outputs(syntactic_check_results map[string][][]string) {

	Report_syntactic_check_issues(syntactic_check_results["issues"])
	Report_syntactic_check_issue_parameters(syntactic_check_results["issue parameters"])
	Report_syntactic_check_fixes(syntactic_check_results["fixes"])
}

func Report_syntactic_check_issues(issue_transactions [][]string) {

	output_csv_filename := "issues.csv" //TODO - Stage 2 - change to folder path (outputs/issues.csv)

	output_header := []string{
		"check_uuids",
		"check_type_uuids",
		"object_id",
		"column_uudis"}

	storage.Write_slice_with_header_to_csv(
		issue_transactions,
		output_header,
		output_csv_filename)

}

//TODO - Stage 1 - too WET either split or generalise.
func Report_syntactic_check_issue_parameters(issue_parameter_transactions [][]string) {

	output_csv_filename := "issue_parameters.csv"

	output_header := []string{
		"check_uuids",
		"parameter_sequence_number",
		"parameter_values"}

	storage.Write_slice_with_header_to_csv(
		issue_parameter_transactions,
		output_header,
		output_csv_filename)

}

func Report_syntactic_check_fixes(fix_transactions [][]string) {

	output_csv_filename := "outputs\\fixes.csv"

	output_header := []string{
		"issue_uuids",
		"cell_values_original",
		"cell_values_marked",
		"cell_values_modified",
		"check_type_uuids",
		"row_id",
		//"column_uuids", 	//TOOD - Stage 1 - add to fix transaction data
		//"table_names",    //TOOD - Stage 1 - add to fix transaction data
		"column_names"}

	storage.Write_slice_with_header_to_csv(
		fix_transactions,
		output_header,
		output_csv_filename)

}

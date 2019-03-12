package reporting

import "storage/csv"

func Report_syntactic_check_issues(issue_transactions [][]string) {

	//#TODO - Stage 2 - change issue report output format (check_uuids, check_type_uuids, object _uuids)

	output_filename := "transactions.csv"

	output_header := []string{
		"check_uuids",
		"cell_values_original",
		"cell_values_marked",
		"cell_values_modified",
		"check_type_uuids",
		"row_id",
		"column_name"}

	output_csv_file, _ := storage.Open_csv_file(output_filename)
	storage.Write_1d_slice_to_csv(output_header, output_csv_file)
	storage.Write_2d_slice_set_to_csv(issue_transactions, output_csv_file)

}

package syntactic_checks

import (
	"fmt"
	storage_csv "storage/csv"
	storage_slices "storage/slices"
)

func Execute_csv_syntactic_checks(csv_file_name string, column_set map[string][]int) {

	//Load Configuratoin+++++++++++++++++++++++++++
	//#TODO Move to Load Configuration
	//#TODO - check configuration should be read from external configuration source.

	var check []interface{}
	var checks [][]interface{}

	//#TODO add check type property for sub-string matching.

	//special characters check
	/*check = append(check,
	`{DCCD33EC-7443-4666-8626-9C2B60ED82EF}`,
	`Assai Character Range`,
	`[\x3B\x40\x5C\x7E\x81\xA0\xB0\xB7\xC2\xC3\xF8\x{2013}\x{201D}\x{2019}\x{2026}]`,
	`STRING.EMPTY`)
	*/
	check = append(check,
		`{A9C658C8-C304-400E-BEEA-96D0DB809AC8}`,
		`Trailing Spaces`,
		`\S{1,}(\s+)$`,
		`STRING.EMPTY`)

	var check_transaction_set [][]string
	checks = append(checks, check)

	//#TODO columns to extract should be read from configuration

	//END OF Load Configuratoin+++++++++++++++++++++++++++

	//Load Data+++++++++++++++++++++
	//#TODO Move to Load data
	var csv_dataset [][]string

	csv_file, csv_file_data := storage_csv.Open_csv_file(csv_file_name)

	fmt.Print("Reading CSV Data\n")
	csv_dataset = storage_csv.Read_csv_to_slice(csv_file, csv_file_data, "")

	//END OF Load Data++++++++++++++++++++

	//Process Columsn+++++++++++++
	//#TODO Move to Process column sets
	// Process Columns
	//check_transactions_string :=
	// [][]string{{
	// "check_uuids",
	// "cell_values_original" ,
	// "cell_values_marked",
	// "cell_values_modified",
	// "check_type_uuids",
	// "row_id",
	// "column_name" }}

	for column_name := range column_set {

		fmt.Printf("\nExtracting Columns %s\n", column_set[column_name])
		extracted_data := storage_slices.Extract_columns_from_2d_slices(csv_dataset, column_set[column_name])

		fmt.Printf("Adding check uuid to extracted data %s\n", column_set[column_name])
		csv_dataset_including_check := storage_slices.Add_single_value_column_to_2d_slice(
			extracted_data,
			check[0].(string))

		fmt.Print("converting cell data to interface %s\n", csv_dataset_including_check)
		csv_dataset_including_check_interface :=
			storage_slices.Convert_2d_string_to_interface(
				csv_dataset_including_check)

		fmt.Print("processing checks\n")
		check_transactions :=
			process_cell_sets(
				csv_dataset_including_check_interface,
				checks)

		//#TODO Add Column information to Transactions.

		if check_transactions != nil {

			check_transactions_string :=
				storage_slices.Convert_2d_interface_to_string(
					check_transactions)

			check_transactions_string = storage_slices.Add_single_value_column_to_2d_slice(check_transactions_string, column_name)

			for _, check_transaction := range check_transactions_string {
				check_transaction_set = append(check_transaction_set, check_transaction)
			}

			//Output
			//fmt.Printf("data :%s\n transactions: \n %s", csv_dataset_including_check_interface, check_transactions)
		}

	}

	output_filename := "transactions.csv"
	output_csv_file, _ := storage_csv.Open_csv_file(output_filename)
	storage_csv.Write_2d_slice_set_to_csv(check_transaction_set, output_csv_file)

}

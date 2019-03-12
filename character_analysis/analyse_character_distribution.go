package character_analysis

import (
	"fmt"
	storage_csv "storage/csv"
	storage_slices "storage/slices"
)

func Process_character_distribution(csv_file_name string, column_set map[string][]int) {

	//#TODO Move to Load data
	var csv_dataset [][]string

	csv_file, csv_file_data := storage_csv.Open_csv_file(csv_file_name)

	fmt.Print("Reading CSV Data\n")
	csv_dataset = storage_csv.Read_csv_to_slice(csv_file, csv_file_data, "")

	//END OF Load Data++++++++++++++++++++

	//Process Columns+++++++++++++

	for column_name := range column_set {

		fmt.Printf("Extracting Columns %s\n", column_set[column_name])

		extracted_data := storage_slices.Extract_columns_from_2d_slices(csv_dataset, column_set[column_name])

		character_map := Report_character_distribution(extracted_data)

		fmt.Printf("column name: %s\n", column_name)
		for character_key, occurance_count := range character_map {

			fmt.Printf("character: %v, [%c], %v\n", character_key, character_key, occurance_count)
		}

	}

}

func Report_character_distribution(extracted_data [][]string) map[rune]int {
	//------------------------------ check special charactars -------------------------------
	character_map := make(map[rune]int)

	for _, row := range extracted_data {

		for _, character := range row[1] {
			character_map[character] += 1

		}

	}
	return character_map
}

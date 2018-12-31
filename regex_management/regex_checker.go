package regex_management

import (
	"regexp"
)

// take : db and in scope table name, table column name, check name
// return : column check results (tranasction list)+ row uuids

/*func Execute_sytactic_checks (accessDatabase *database.MsAccessDatabase, column_data [][]interface{}, check_type string) [][]interface{}{

	//data storage for return
	var check_transaction_dataset_for_column_check_type [][] interface{}
	var check_transaction_dataset_for_cell_check_type []interface{}

	//#TODO read regex check string directly from tables for regex

	for _,column_data_item := range column_data {

		check_transaction_dataset_for_cell_check_type = Syntactically_check_string(check_type, column_data_item)

		if check_transaction_dataset_for_cell_check_type != nil {
			check_transaction_dataset_for_column_check_type = append(check_transaction_dataset_for_column_check_type, check_transaction_dataset_for_cell_check_type)
		}
	}

	// for each cell, run the check

	return check_transaction_dataset_for_column_check_type
}


func Syntactically_check_string (check_type string, check_string []interface{}, regex_string string) []interface{} {

	var check_result_transaction []interface{}

	if check_string[1] == nil {
		return nil
	}

	//check_result_transaction = append(check_result_transaction,check_string[1].(string))

	//#TODO replace switch with a single regex function

	switch check_type {

	case "Leading Spaces":
		regex_string = `^(\s+)\S{1,}`

		check_result_transaction = Process_regex_check(regex_string,check_string[1])

		if check_result_transaction != nil {
			check_result_transaction = append(check_result_transaction,check_string[0].(string))
			check_result_transaction = append(check_result_transaction,"Leading Spaces")

			return check_result_transaction

		}

	case "Trailing Spaces":

		regex_string := `\S{1,}(\s+)$`

		check_result_transaction = Process_regex_check(regex_string,check_string[1])

		if check_result_transaction != nil {

			check_result_transaction = append(check_result_transaction,check_string[0].(string))
			check_result_transaction = append(check_result_transaction,"Trailing Spaces")

			return check_result_transaction

		}

	}

	return nil

}

*/
func Process_regex_check(regex_string string, check_string interface{}, replacement_string_type string) []interface{} {

	var check_result_transaction []interface{}
	var mark_string rune
	var replacement_string rune

	syntactic_check_regex_object, _ := regexp.Compile(regex_string)

	regex_find_result := syntactic_check_regex_object.FindString(check_string.(string))
	regex_fine_result_with_index := syntactic_check_regex_object.FindStringIndex(check_string.(string))

	mark_string = '~'
	switch replacement_string_type {
	case "STRING.EMPTY":
		replacement_string = 0
	}

	//fmt.Printf("\n-->looking for issue using regex :%s, string: %s\n ", regex_string, check_string.(string))

	if regex_find_result != "" {
		//fmt.Printf("\n-->found issue using regex :%s, string: %s\n ", regex_string, check_string.(string))
		check_result_transaction = append(check_result_transaction, check_string.(string))

		//check_result_transaction = append(check_result_transaction,regex_find_result)

		//fmt.Printf("\n-->creating replacement mark at :%s for string: %s\n ", regex_fine_result_with_index[0], check_string.(string))
		marked_text := replaceAtIndex(check_string.(string), mark_string, regex_fine_result_with_index[0])

		check_result_transaction = append(check_result_transaction, marked_text)

		modified_text := replaceAtIndex(check_string.(string), replacement_string, regex_fine_result_with_index[0])

		check_result_transaction = append(check_result_transaction, modified_text)

		return check_result_transaction

	}

	return nil

}

func replaceAtIndex(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

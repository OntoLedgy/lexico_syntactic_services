package regex_processor

import (
	"database_manager/utils"
	"fmt"
	"regexp"
)

func Process_regex_check(
	regex_string string,
	cell_value_original interface{},
	replacement_string_type string) []interface{} {

	var check_result_transaction []interface{}
	var replacement_string string

	mark_string := "~" //TODO - Stage 2 - move to general config

	cell_value_original_string :=
		cell_value_original.(string)

	switch replacement_string_type {

	case "STRING.EMPTY":
		replacement_string = ""

	case "SPACE":
		replacement_string = " "
		//#TODO - Stage 2 - add other replacement string type cases
	}

	syntactic_check_regex_object := //TODO - Stage 2 - add error handling
		regexp.MustCompile(regex_string)

	regex_match_indices :=
		syntactic_check_regex_object.
			FindAllStringSubmatchIndex(
				cell_value_original_string,
				-1)

	if len(regex_match_indices) > 0 {

		check_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		//TODO - Stage 2 - 	separate replacement process from regex check, return sub_match_indicies
		// 					for aggregatation first and then call modification function in one go.

		cell_value_marked :=
			modify_string_by_index(
				cell_value_original_string,
				mark_string,
				regex_match_indices)

		cell_value_modified :=
			modify_string_by_index(
				cell_value_original_string,
				replacement_string,
				regex_match_indices)

		check_result_transaction =
			append(check_result_transaction,
				check_uuid.String(),
				cell_value_original_string,
				cell_value_marked,
				cell_value_modified)

		return check_result_transaction

	} else {

		return nil
	}

}

//TODO - Stage 1 - move out to differnet file/pacakge

func modify_string_by_index(
	string_for_replacement string,
	replacement_string string,
	replacement_indicies [][]int) string {

	var modified_string string
	var replacement_length int
	var replacement_start_position int
	var replacement_end_position int

	fmt.Printf(
		"\nstring_for_repalcement: %s, replacement_indicies : %v",
		string_for_replacement,
		replacement_indicies)

	replacement_string_length :=
		len(replacement_string)

	replacement_offset :=
		0

	for _, index_value := range replacement_indicies {

		//TODO - Stage 1 - break out the if into separate function
		if len(index_value) > 2 {

			replacement_start_position = index_value[2]
			replacement_end_position = index_value[3]

		} else {

			replacement_start_position = index_value[0]
			replacement_end_position = index_value[1]
		}

		replacement_length =
			replacement_end_position -
				replacement_start_position

		modified_string =
			string_for_replacement[:replacement_start_position+replacement_offset] +
				replacement_string +
				string_for_replacement[replacement_end_position+replacement_offset:]

		replacement_offset =
			replacement_offset +
				replacement_string_length -
				replacement_length

		string_for_replacement =
			modified_string

	}

	return modified_string
}

package regex_processor

import (
	"database_manager/utils"
	"github.com/satori/go.uuid"
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
		//TODO - Stage 2 - add other replacement string type cases
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

		check_result_transaction =
			generate_fix_transactions(
				check_uuid,
				cell_value_original_string,
				replacement_string,
				mark_string,
				regex_match_indices)

		return check_result_transaction

	} else {

		return nil
	}

}

func generate_fix_transactions(
	check_uuid uuid.UUID,
	cell_value_original_string string,
	mark_string string,
	replacement_string string,
	regex_match_indices [][]int) []interface{} {

	var check_result_transaction []interface{}

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
}

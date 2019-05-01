package regex_processor

import (
	"database_manager/utils"
	"regexp"
	"syntactic_checker/object_model"
)

func Process_regex_check(
	regex_string string,
	cell_value_original interface{},
	replacement_string_type string) *object_model.Regex_check_results {

	//var check_result_transaction []interface{}
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

	syntactic_check_regex_object :=
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

		regex_check_result := object_model.Regex_check_results{
			check_uuid.String(),
			cell_value_original_string,
			mark_string,
			replacement_string,
			regex_match_indices,
		}

		return &regex_check_result

	} else {

		return nil
	}

}

package regex_management

import (
	"database_manager/utils"
	"fmt"
	"regexp"
)

func Process_regex_check(
	regex_string string,
	check_string interface{},
	replacement_string_type string) []interface{} { //#TODO separate replacement process from regex check

	var check_result_transaction []interface{}
	var mark_string string
	var replacement_string string

	mark_string = "~" //#TODO move to general config

	switch replacement_string_type {

	case "STRING.EMPTY":
		replacement_string = ""

	case "SPACE":
		replacement_string = " "
		//#TODO add other replacement string type cases
	}

	syntactic_check_regex_object := //#TODO add error handling
		regexp.MustCompile(regex_string)

	check_string_string_value := check_string.(string)

	regex_find_result :=
		syntactic_check_regex_object.FindString(
			check_string_string_value)

	if regex_find_result != "" {

		check_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		sub_match_indices :=
			syntactic_check_regex_object.
				FindAllStringSubmatchIndex(
					check_string.(string),
					-1)

		marked_text :=
			modify_string_by_index(
				check_string.(string),
				mark_string,
				sub_match_indices)

		modified_text :=
			modify_string_by_index(
				check_string.(string),
				replacement_string,
				sub_match_indices)

		check_result_transaction =
			append(check_result_transaction,
				check_uuid.String(),
				check_string.(string),
				marked_text,
				modified_text)

		return check_result_transaction
	}

	return nil
}

func modify_string_by_index(string_for_replacement string, replacement_string string, indicies [][]int) string {

	var modified_string string

	//#TODO this should iterate through the index for multiple matches will require recalculating the resultant string's length for each modification. (shift back (subtract) from the current index by previous range - length of replacement string)

	fmt.Printf("\nstring_for_repalcement: %s, indicies : %v", string_for_replacement, indicies)

	//for single substring replacement
	/*modified_string = string_for_replacement[:indicies[0][2]] + replacement_string + string_for_replacement[indicies[0][3]:]*/

	//for main match multiple replacement

	replacement_string_length := len(replacement_string)
	replacement_offset := 0

	for _, index_value := range indicies {

		change_length := index_value[1] - index_value[0]

		modified_string = string_for_replacement[:index_value[0]+replacement_offset] + replacement_string + string_for_replacement[index_value[1]+replacement_offset:]
		replacement_offset = replacement_offset + replacement_string_length - change_length
		string_for_replacement = modified_string

	}

	return modified_string
}

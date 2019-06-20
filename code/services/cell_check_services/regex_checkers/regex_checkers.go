package regex_checkers

import (
	"database_manager/utils"
	"regexp"
)

func Process_regex_check(
	regex_string string,
	in_scope_cell string) *RegexCheckResults {

	cell_value_original_string :=
		in_scope_cell

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

		regex_check_result := RegexCheckResults{
			check_uuid.String(),
			cell_value_original_string,
			regex_match_indices,
		}

		return &regex_check_result

	} else {

		return nil
	}

}

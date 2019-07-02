package regex_checkers

import (
	"database_manager/utils"
	"regexp"
	"string_editor/object_model"
)

type regexCheckers struct {
	regex_check_result RegexCheckResults
	String_edit_ranges []object_model.StringEditRanges
}

func (
	regex_checker *regexCheckers) Process_regex_check(
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

		regex_checker.
			regex_check_result =
			regex_check_result

		regex_checker.
			convert_regex_result_to_string_edit_ranges()

		return &regex_checker.regex_check_result

	} else {

		return nil
	}

}

func (regex_checker *regexCheckers) convert_regex_result_to_string_edit_ranges() {

	replacement_indicies :=
		regex_checker.
			regex_check_result.
			Regex_match_indices

	regex_checker.
		String_edit_ranges =
		make([]object_model.StringEditRanges, len(replacement_indicies))

	for index, replacement_index := range replacement_indicies {

		replacement_start_position, replacement_end_position :=
			regex_checker.get_replacement_positions(replacement_index)

		regex_checker.
			String_edit_ranges[index].
			Constructor(
				replacement_start_position,
				replacement_end_position-replacement_start_position)
	}

}

func (regex_checker *regexCheckers) get_replacement_positions(
	replacement_index []int) (int, int) {

	var replacement_start_position int
	var replacement_end_position int

	if len(replacement_index) > 2 {

		replacement_start_position = replacement_index[2]
		replacement_end_position = replacement_index[3]

	} else {

		replacement_start_position = replacement_index[0]
		replacement_end_position = replacement_index[1]
	}

	return replacement_start_position, replacement_end_position

}

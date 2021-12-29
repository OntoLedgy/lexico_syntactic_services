package regex_checkers

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
	"github.com/OntoLedgy/string_editing_services/object_model"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
	"regexp"
)

type regexCheckers struct {
	check_regex_pattern string
	string_value        string
	regex_check_result  service_results.RegexCheckResults
	String_edit_ranges  []object_model.StringEditRanges
}

func (
	regex_checker *regexCheckers) Process_regex_check() *service_results.RegexCheckResults {

	string_value_original_string :=
		regex_checker.string_value

	regex_object :=
		regexp.MustCompile(
			regex_checker.check_regex_pattern)

	regex_match_indices :=
		regex_object.
			FindAllStringSubmatchIndex(
				string_value_original_string,
				-1)

	regex_pattern_found := len(regex_match_indices) > 0

	if regex_pattern_found {

		check_uuid, _ :=
			utils.GetUUID(
				1,
				"")

		regex_check_result := service_results.RegexCheckResults{
			check_uuid.UUID.String(),
			string_value_original_string,
			regex_match_indices,
		}

		regex_checker.
			regex_check_result =
			regex_check_result

		regex_checker.
			append_replacement_indicies_to_string_edit_range()

		return &regex_checker.regex_check_result

	} else {

		return nil
	}

}

//TODO - Stage 2 - move to a separate type (adaptor / convertor from regex results to string edit ranges)
func (
	regex_checker *regexCheckers) append_replacement_indicies_to_string_edit_range() {

	replacement_indicies :=
		regex_checker.
			regex_check_result.
			Regex_match_indices

	for _, replacement_index := range replacement_indicies {

		regex_checker.
			append_replacement_index_to_string_edit_range(
				replacement_index)

	}

}

func (
	regex_checker *regexCheckers) append_replacement_index_to_string_edit_range(
	replacement_index []int) {

	var replacement_index_position int

	if len(replacement_index) > 2 {

		replacement_index_position = 2

		regex_checker.append_replacement_index(
			replacement_index,
			replacement_index_position)

	} else {

		replacement_index_position = 0

		regex_checker.append_replacement_index(
			replacement_index,
			replacement_index_position)

	}
}

func (
	regex_checker *regexCheckers) append_replacement_index(
	replacement_index []int,
	replacement_index_position int) {

	var replacement_start_position int
	var replacement_end_position int

	for {

		replacement_start_position =
			replacement_index[replacement_index_position]

		replacement_end_position =
			replacement_index[replacement_index_position+1]

		regex_checker.
			append_string_edit_range(
				replacement_start_position,
				replacement_end_position)

		replacement_index_position += 2

		if replacement_index_position >= len(replacement_index) {

			break
		}
	}

}

func (
	regex_checker *regexCheckers) append_string_edit_range(
	replacement_start_position int,
	replacement_end_position int) {

	var string_edit_range object_model.StringEditRanges

	string_edit_range.
		Constructor(
			replacement_start_position,
			replacement_end_position-replacement_start_position)

	regex_checker.String_edit_ranges =
		append(
			regex_checker.String_edit_ranges,
			string_edit_range)
}

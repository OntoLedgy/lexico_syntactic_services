package regex_processor

import "fmt"

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

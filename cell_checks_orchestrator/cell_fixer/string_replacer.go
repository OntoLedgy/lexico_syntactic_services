package cell_fixer

import "fmt"

func modify_string_by_index(
	string_for_replacement string,
	replacement_string string,
	replacement_indicies [][]int) string {

	var modified_string string
	var replacement_length int

	fmt.Printf(
		"\nstring_for_repalcement: %s, replacement_indicies : %v",
		string_for_replacement,
		replacement_indicies)

	replacement_string_length :=
		len(replacement_string)

	replacement_offset :=
		0

	for _, replacement_index := range replacement_indicies {

		replacement_start_position, replacement_end_position :=
			get_replacement_positions(replacement_index)

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

func get_replacement_positions(replacement_index []int) (int, int) {

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

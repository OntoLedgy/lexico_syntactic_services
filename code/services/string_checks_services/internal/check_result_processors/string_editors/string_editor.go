package string_editors

import (
	"logger/standard_global_logger"
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

//TODO - Stage 3 - replace with String Editor functionality

type stringEditors struct {
	string_value       string
	replacement_string string
	check_results      *service_results.StringCheckResults
}

func (
	string_value_editor *stringEditors) Edit_string(
	string_edit_history *string_editor_object_model.StringEditHistories) *string_editor_object_model.StringEditHistories {

	var modified_string, marked_string string

	string_to_edit :=
		string_edit_history.
			GetCurrentString()

	string_to_edit = string_edit_history.GetCurrentString()

	standard_global_logger.Global_logger.Printf(
		"\nString_for_repalcement: %s, replacement_char(s): [%s], replacement_indicies: %v",
		string_to_edit,
		string_value_editor.
			replacement_string,
		string_value_editor.check_results)

	modified_string =
		string_to_edit

	marked_string =
		string_to_edit

	//TODO - add stringeditranges to replacement indices

	modified_string =
		string_value_editor.
			edit_string_using_indicies(
				string_value_editor.
					replacement_string,
				modified_string)

	marked_string =
		string_value_editor.
			edit_string_using_indicies(
				object_model.
					Modification_marker,
				marked_string)

	string_edit_history.
		Set_string_changes(
			modified_string,
			marked_string)

	return string_edit_history
}

func (
	string_value_editor *stringEditors) edit_string_using_indicies(
	replacement_string string,
	string_to_edit string) string {

	replacement_offset :=
		0

	string_edit_ranges := string_value_editor.check_results.Check_result_string_edit_ranges

	for _, edit_range := range string_edit_ranges {

		string_to_edit, replacement_offset =
			string_value_editor.
				edit_string_using_index(
					edit_range,
					string_to_edit,
					replacement_offset,
					replacement_string)

	}
	return string_to_edit
}

func (
	string_value_editor *stringEditors) edit_string_using_index(
	edit_range string_editor_object_model.StringEditRanges,
	original_string string,
	replacement_offset int,
	replacement_string string) (
	string,
	int) {

	edit_start_position :=
		edit_range.
			Start_position

	edit_end_position :=
		edit_start_position +
			edit_range.Range_length

	replacement_length :=
		edit_end_position -
			edit_start_position

	replacement_string_length :=
		len(replacement_string)

	modified_string :=
		original_string[:edit_start_position+replacement_offset] +
			replacement_string +
			original_string[edit_end_position+replacement_offset:]

	replacement_offset =
		replacement_offset +
			replacement_string_length -
			replacement_length

	return modified_string, replacement_offset
}

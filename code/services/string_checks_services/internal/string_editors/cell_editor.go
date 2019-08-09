package string_editors

import (
	"fmt"
	"string_editor/factories"
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/identified_strings"
)

//TODO - Stage 3 - replace with String Editor functionality

type stringEditors struct {
	identified_string_to_edit identified_strings.IdentifiedStrings //TODO - deprecate
	string_value              string
	replacement_string        string
	check_results             *check_results.CheckResults
}

func (
	string_value_editor *stringEditors) Edit_string() *string_editor_object_model.StringEditHistory {

	var modified_string, marked_string string
	var string_editor_factory factories.StringEditorFactory

	string_to_edit :=
		string_value_editor.
			string_value

	//TODO Deprecate
	string_to_edit =
		string_value_editor.
			identified_string_to_edit.
			String_value

	fmt.Printf(
		"\nString_for_repalcement: %s, replacement_char(s): [%s], replacement_indicies: %v",
		string_to_edit,
		string_value_editor.
			replacement_string,
		string_value_editor.check_results)

	string_editor :=
		string_editor_factory.
			CreateStringEditor(
				string_to_edit)

	modified_string =
		string_to_edit

	marked_string =
		string_to_edit

	//TODO - add stringeditranges to replacement indices

	modified_string =
		string_value_editor.
			modify_string_using_indicies(
				string_value_editor.
					replacement_string,
				modified_string)

	marked_string =
		string_value_editor.
			modify_string_using_indicies(
				object_model.
					Modification_marker,
				marked_string)

	string_edit_history :=
		string_editor.
			Get_string_edit_history()

	string_edit_history.
		Set_string_changes(
			modified_string,
			marked_string)

	return string_edit_history
}

func (string_value_editor *stringEditors) modify_string_using_indicies(
	replacement_string string,
	string_to_modify string) string {

	replacement_offset :=
		0

	string_edit_ranges := string_value_editor.check_results.Check_result_string_edit_ranges

	for _, edit_range := range string_edit_ranges {

		string_to_modify, replacement_offset =
			string_value_editor.
				modify_string_using_index(
					edit_range,
					string_to_modify,
					replacement_offset,
					replacement_string)

	}
	return string_to_modify
}

func (
	string_value_editor *stringEditors) modify_string_using_index(
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

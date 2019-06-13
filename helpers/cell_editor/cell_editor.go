package cell_editor

import (
	"fmt"
	"string_editor/factories"
	"string_editor/interfaces"
	"string_editor/object_model"
	"syntactic_checker/helpers"
	"syntactic_checker/helpers/configuration_handler"
	"syntactic_checker/object_model/issues"
)

//TODO - Stage 3 - replace with String Editor functionality

func Edit_cell(
	string_to_edit string,
	issue_type issues.IssueTypes,
	replacement_indicies [][]int) (string, string) {

	var modified_string, marked_string string
	var string_editor_factory factories.StringEditorFactory

	replacement_string :=
		helpers.
			Get_replacement_string(
				issue_type)

	fmt.Printf(
		"\nstring_for_repalcement: %s, replacement_char(s): [%s], replacement_indicies : %v",
		string_to_edit,
		replacement_string,
		replacement_indicies)

	string_editor :=
		string_editor_factory.CreateStringEditor(
			string_to_edit)

	modified_string =
		string_to_edit

	marked_string =
		string_to_edit

	modified_string =
		modify_string_using_indicies(
			replacement_string,
			replacement_indicies,
			modified_string,
			string_editor,
			issue_type)

	marked_string =
		modify_string_using_indicies(
			configuration_handler.Modification_marker,
			replacement_indicies,
			marked_string,
			string_editor,
			issue_type)

	return modified_string, marked_string
}

func modify_string_using_indicies(
	replacement_string string,
	replacement_indicies [][]int,
	modified_string string,
	string_editor interfaces.IStringEditors,
	issue_type issues.IssueTypes) string {

	replacement_offset :=
		0
	for _, replacement_index := range replacement_indicies {

		modified_string, replacement_offset =
			modify_string_using_index(
				replacement_index,
				string_editor,
				issue_type,
				modified_string,
				replacement_offset,
				replacement_string)

	}
	return modified_string
}

func modify_string_using_index(
	replacement_index []int,
	string_editor interfaces.IStringEditors,
	issue_type issues.IssueTypes,
	original_string string,
	replacement_offset int,
	replacement_string string) (string, int) {

	edit_start_position, edit_end_position := //TODO - Stage 2 - using a single variable with two values.
		get_replacement_positions(replacement_index)

	replacement_length :=
		edit_end_position -
			edit_start_position

	replacement_string_length :=
		len(replacement_string)

	edit_range :=
		factories.
			CreateStringEditRange(
				edit_start_position,
				replacement_length)

	execute_string_edit_transaction(
		string_editor.(*object_model.StringEditors),
		issue_type.Issue_check_replacement_string,
		edit_range.(*object_model.StringEditRanges))

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

//TODO - Stage 3 - replace with String Editor Range functionality

func get_replacement_positions(
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

func execute_string_edit_transaction(
	string_editor *object_model.StringEditors,
	operation_type string,
	edit_range *object_model.StringEditRanges) {

	switch operation_type {

	case "STRING.EMPTY":
		string_editor.Delete(*edit_range)

	case "SPACE":
		string_editor.Insert(*edit_range, " ")
	}

}

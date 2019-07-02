package cell_editors

import (
	"fmt"
	"string_editor/factories"
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/services/syntactic_checking_services/configuration_getters"
)

//TODO - Stage 3 - replace with String Editor functionality

type cellEditor struct {
	cell_to_edit       cells.Cells
	replacement_string string
	check_results      *check_results.CheckResults
}

func (
	cell_editor *cellEditor) Edit_cell() *string_editor_object_model.StringEditHistory {

	var modified_string, marked_string string
	var string_editor_factory factories.StringEditorFactory

	string_to_edit :=
		cell_editor.
			cell_to_edit.
			Cell_value

	fmt.Printf(
		"\nString_for_repalcement: %s, replacement_char(s): [%s], replacement_indicies: %v",
		string_to_edit,
		cell_editor.
			replacement_string,
		cell_editor.check_results)

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
		cell_editor.
			modify_string_using_indicies(
				cell_editor.
					replacement_string,
				modified_string)

	marked_string =
		cell_editor.
			modify_string_using_indicies(
				configuration_getters.
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

func (cell_editor *cellEditor) modify_string_using_indicies(
	replacement_string string,
	string_to_modify string) string {

	replacement_offset :=
		0

	string_edit_ranges := cell_editor.check_results.Check_result_string_edit_ranges

	for _, edit_range := range string_edit_ranges {

		string_to_modify, replacement_offset =
			cell_editor.
				modify_string_using_index(
					edit_range,
					string_to_modify,
					replacement_offset,
					replacement_string)

	}
	return string_to_modify
}

func (
	cell_editor *cellEditor) modify_string_using_index(
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

package deprecate_cell_editors

import (
	"fmt"
	"string_editor/factories"
	"string_editor/interfaces"
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_services/internal/regex_checkers"
	"syntactic_checker/code/services/syntactic_checking_services/configuration_getters"
	//"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters"
)

//TODO - Stage 3 - replace with String Editor functionality
//TODO - Stage 2 - remove dependence on issue type and use replacement string value instead

type cellEditor struct {
	cell_to_edit       cells.Cells
	issue_type         issues.IssueTypes
	replacement_string string
	regex_result       regex_checkers.RegexCheckResults
}

func (
	cell_editor *cellEditor) Edit_cell() *string_editor_object_model.StringEditHistory {

	var modified_string, marked_string string
	var string_editor_factory factories.StringEditorFactory

	string_to_edit :=
		cell_editor.
			cell_to_edit.
			Cell_value

	issue_type :=
		cell_editor.
			issue_type

	replacement_indicies :=
		cell_editor.
			regex_result.
			Regex_match_indices

	replacement_string :=
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
		cell_editor.modify_string_using_indicies(
			replacement_string,
			replacement_indicies,
			modified_string,
			string_editor,
			issue_type)

	marked_string =
		cell_editor.
			modify_string_using_indicies(
				configuration_getters.Modification_marker,
				replacement_indicies,
				marked_string,
				string_editor,
				issue_type)

	string_edit_history :=
		string_editor.
			Get_string_edit_history()

	string_edit_history.
		Set_string_changes(
			modified_string,
			marked_string)

	return string_edit_history
}

//TODO - Stage 3 - replace with String Editor Range functionality - strip out regex results to string edit range functionality
//TODO - Stage 2 - move this to regex checkers

func (cell_editor *cellEditor) modify_string_using_indicies(
	replacement_string string, //remove
	replacement_indicies [][]int,
	string_to_modify string,
	string_editor interfaces.IStringEditors,
	issue_type issues.IssueTypes, //remove
) string {

	replacement_offset :=
		0
	for _, replacement_index := range replacement_indicies {

		string_to_modify, replacement_offset =
			modify_string_using_index(
				replacement_index,
				string_editor,
				issue_type,
				string_to_modify,
				replacement_offset,
				replacement_string)

	}
	return string_to_modify
}

func modify_string_using_index(
	replacement_index []int,
	string_editor interfaces.IStringEditors,
	issue_type issues.IssueTypes, //remove
	original_string string,
	replacement_offset int,
	replacement_string string, //remove
) (string, int) {

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
		string_editor,
		issue_type.Issue_check_replacement_string,
		edit_range.(*string_editor_object_model.StringEditRanges))

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
	string_editor interfaces.IStringEditors,
	operation_type string,
	edit_range *string_editor_object_model.StringEditRanges) {

	switch operation_type {

	case "STRING.EMPTY":
		string_editor.Delete(*edit_range)

	case "SPACE":
		string_editor.Insert(*edit_range, " ")
	}

}

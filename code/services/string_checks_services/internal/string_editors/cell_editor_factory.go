package string_editors

import (
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/identified_strings"
)

func Create(
	identified_string_to_edit identified_strings.IdentifiedStrings,
	string_value string,
	check_results *check_results.CheckResults,
	replacement_string string) *stringEditors {

	string_editor :=
		new(
			stringEditors)

	string_editor.
		identified_string_to_edit =
		identified_string_to_edit

	string_editor.
		string_value =
		string_value

	string_editor.
		replacement_string =
		replacement_string

	string_editor.
		check_results =
		check_results

	return string_editor

}

package string_editors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

func Create(
	string_value string,
	check_results *service_results.StringCheckResults,
	replacement_string string) *stringEditors {

	string_editor :=
		new(
			stringEditors)

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

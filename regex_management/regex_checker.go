package regex_management

import (
	"regexp"
)

func Process_regex_check(
	regex_string string,
	check_string interface{},
	replacement_string_type string) []interface{} { //#TODO separate replacement process from regex check

	var check_result_transaction []interface{}
	var mark_string rune
	var replacement_string rune

	syntactic_check_regex_object, _ := //#TODO add error handling
		regexp.Compile(regex_string)

	regex_find_result :=
		syntactic_check_regex_object.FindString(
			check_string.(string))

	regex_find_result_with_index :=
		syntactic_check_regex_object.FindStringIndex(
			check_string.(string))

	mark_string = '~'

	switch replacement_string_type {

	case "STRING.EMPTY":
		replacement_string = 0
		//#TODO add other replacement string type cases
	}

	if regex_find_result != "" {

		check_result_transaction =
			append(check_result_transaction,
				check_string.(string))

		marked_text :=
			replaceAtIndex(
				check_string.(string),
				mark_string,
				regex_find_result_with_index[0])

		check_result_transaction =
			append(
				check_result_transaction,
				marked_text)

		modified_text :=
			replaceAtIndex(
				check_string.(string),
				replacement_string,
				regex_find_result_with_index[0])

		check_result_transaction =
			append(check_result_transaction,
				modified_text)

		return check_result_transaction

	}

	return nil

}

func replaceAtIndex(
	string_to_be_modified string,
	replacement_rune rune,
	replacement_index int) string {

	modified_string :=
		[]rune(
			string_to_be_modified)

	modified_string[replacement_index] =
		replacement_rune

	return string(modified_string)
}

package string_fix_getters

import (
	"fmt"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type fixProcessors struct {
	issue_types       []issues.IssueTypes
	identified_string identified_strings.IdentifiedStrings
	string_value      string
	string_checks_fix fixes.Fixes
}

func (fix_processor *fixProcessors) Get_string_check_fix(
	issue_types []issues.IssueTypes,
	identified_string identified_strings.IdentifiedStrings) fixes.Fixes {

	fmt.Printf(
		"\nProcessing fixes for %s...\n",
		identified_string)

	string_check_fix :=
		Generate_fix_transaction(
			fix_processor.string_value,
			identified_string, //TODO - renaming strings - Deprecate
			fix_processor.issue_types)

	fmt.Printf(
		"\nFix transaction: %v \n",
		string_check_fix)

	return string_check_fix

}

package identified_string_checks_processors

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

func Create(identified_string identified_strings.IdentifiedStrings, issue_types []issues.IssueTypes) *identifiedStringChecksProcessor {

	identified_string_checks_processor := new(identifiedStringChecksProcessor)

	identified_string_checks_processor.identified_string = identified_string
	identified_string_checks_processor.issue_types = issue_types

	return identified_string_checks_processor
}

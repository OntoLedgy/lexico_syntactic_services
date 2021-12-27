package identified_string_checks_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
)

func Create(
	identified_string *identified_strings.IdentifiedStrings,
	issue_types []issues.IssueTypes) *identifiedStringChecksProcessor {

	identified_string_checks_processor := new(identifiedStringChecksProcessor)

	//TODO - tidy this up, null strings should not be processed/ should be reported
	if identified_string != nil {
		identified_string_checks_processor.identified_string = identified_string
	} else {
		identified_string_checks_processor.identified_string = new(identified_strings.IdentifiedStrings)
		identified_string_checks_processor.identified_string.String_identified = new(identified_strings.Strings)
	}

	identified_string_checks_processor.issue_types = issue_types

	return identified_string_checks_processor
}

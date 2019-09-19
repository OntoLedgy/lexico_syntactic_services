package service_results

import (
	"syntactic_checker/code/object_model/identified_strings"
)

type IdentifiedStringChecksResults struct {
	Identified_string    *identified_strings.IdentifiedStrings
	String_checks_result *StringChecksResults
}

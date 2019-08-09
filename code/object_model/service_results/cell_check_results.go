package service_results

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type StringChecksResults struct {
	Identified_string               identified_strings.IdentifiedStrings
	Identified_string_checks_issues []issues.Issues
	Identified_string_checks_fix    fixes.Fixes
}

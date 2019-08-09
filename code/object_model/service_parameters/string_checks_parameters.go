package service_parameters

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type StringChecksParameters struct {
	Identified_string    identified_strings.IdentifiedStrings
	String_value         string
	In_scope_issue_types []issues.IssueTypes
}

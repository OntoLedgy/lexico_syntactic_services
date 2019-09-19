package service_inputs

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type StringCheckInputs struct {
	String_to_check     *identified_strings.Strings
	In_scope_issue_type issues.IssueTypes
}

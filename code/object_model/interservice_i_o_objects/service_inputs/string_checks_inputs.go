package service_inputs

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type StringChecksInputs struct {
	String_to_check *identified_strings.Strings
	Issue_types     []issues.IssueTypes
}

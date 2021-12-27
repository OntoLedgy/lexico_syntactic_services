package service_inputs

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
)

type StringChecksInputs struct {
	String_to_check *identified_strings.Strings
	Issue_types     []issues.IssueTypes
}

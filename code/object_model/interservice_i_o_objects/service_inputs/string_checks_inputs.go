package service_inputs

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/identified_strings"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
)

type StringChecksInputs struct {
	StringToCheck *identified_strings.Strings
	IssueTypes    []issues.IssueTypes
}

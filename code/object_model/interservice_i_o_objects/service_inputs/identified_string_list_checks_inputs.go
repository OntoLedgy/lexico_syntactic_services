package service_inputs

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type IdentifiedStringListChecksInputs struct {
	Identified_string_list       identified_strings.IdentifiedStringLists
	List_of_in_scope_issue_types []issues.IssueTypes
}

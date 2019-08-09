package service_parameters

import (
	"syntactic_checker/code/object_model/identified_strings"
	"syntactic_checker/code/object_model/issues"
)

type IdentifiedStringListChecksParameters struct {
	Identified_string_list       identified_strings.IdentifiedStringLists
	List_of_in_scope_issue_types []issues.IssueTypes
}

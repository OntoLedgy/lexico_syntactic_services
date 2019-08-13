package service_parameters

import (
	"syntactic_checker/code/object_model/issues"
)

type StringChecksParameters struct {
	String_value         string
	In_scope_issue_types []issues.IssueTypes
}

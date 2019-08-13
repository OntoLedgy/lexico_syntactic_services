package service_parameters

import (
	"syntactic_checker/code/object_model/issues"
)

type StringCheckParameters struct {
	String_value        string
	In_scope_issue_type issues.IssueTypes
}

package service_parameters

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
)

type CellCheckParameters struct {
	In_scope_cell       cells.Cells
	In_scope_issue_type issues.IssueTypes
}

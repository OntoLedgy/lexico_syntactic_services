package service_parameters

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
)

type CellChecksParameters struct {
	In_scope_cell                cells.Cells
	List_of_in_scope_issue_types []issues.IssueTypes
}

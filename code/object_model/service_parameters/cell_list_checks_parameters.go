package service_parameters

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/issues"
)

type CellListChecksParameters struct {
	List_of_in_scope_cells       cells.ListOfCells
	List_of_in_scope_issue_types []issues.IssueTypes
}

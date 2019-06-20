package transactions

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/issues"
)

type CellIssuesTransactions struct {
	In_scope_cell object_model.Cells
	Cell_issues   []issues.Issues
}

package service_results

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
)

type CellChecksResults struct {
	In_scope_cell      cells.Cells
	Cell_checks_issues []issues.Issues
	Cell_checks_fix    fixes.Fixes
}

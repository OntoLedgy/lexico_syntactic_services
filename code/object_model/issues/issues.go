package issues

import (
	"core_foundation/core_object_model"
	"syntactic_checker/code/object_model"
)

type Issues struct {
	core_object_model.Objects
	Issue_type          IssueTypes
	Cell                object_model.InScopeCells
	Marked_cell_value   string
	Modified_cell_value string
}

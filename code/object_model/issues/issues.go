package issues

import (
	"core_foundation/core_object_model"
	"string_editor/object_model"
)

type Issues struct {
	core_object_model.Objects
	Issue_type        IssueTypes
	Cell_edit_history object_model.StringEditHistory
}

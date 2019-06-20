package fixes

import (
	"core_foundation/core_object_model"
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model"
)

type Fixes struct {
	core_object_model.Objects
	Cell                object_model.Cells
	String_edit_history *string_editor_object_model.StringEditHistory
}

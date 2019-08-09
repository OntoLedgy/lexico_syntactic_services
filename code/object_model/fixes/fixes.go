package fixes

import (
	"core_foundation/core_object_model"
	"string_editor/object_model"
)

type Fixes struct {
	core_object_model.Objects
	String_value_edit_history object_model.StringEditHistory
}

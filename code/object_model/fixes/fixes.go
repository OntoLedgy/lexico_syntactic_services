package fixes

import (
	"core_foundation/core_object_model"
	"syntactic_checker/code/object_model"
)

type Fixes struct {
	core_object_model.Objects
	Cell            object_model.InScopeCells
	Marked_string   string
	Modified_string string
}

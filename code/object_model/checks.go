package object_model

import (
	"core_foundation/core_object_model"
	"syntactic_checker/code/object_model/check_types"
)

type Checks struct {
	core_object_model.Objects
	check_type check_types.CheckTypes
}

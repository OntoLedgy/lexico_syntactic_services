package object_model

import (
	"core_foundation/core_object_model"
	"syntactic_checker/object_model/check_types"
)

type Checks struct {
	core_object_model.Objects
	check_names string
	check_type  check_types.CheckTypes
}

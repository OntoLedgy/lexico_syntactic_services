package object_model

import (
	"core_foundation/object_model"
)

type Checks struct {
	object_model.Objects
	check_names string
	check_type  CheckTypes
}

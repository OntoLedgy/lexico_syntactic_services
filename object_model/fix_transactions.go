package object_model

import "core_foundation/core_object_model"

type FixTransactions struct {
	core_object_model.Objects
	check_uuids     string
	original_string string
	marked_string   string
	modified_string string
}

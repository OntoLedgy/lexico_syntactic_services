package object_model

import "core_foundation/object_model"

type FixTransactions struct {
	object_model.Objects
	check_uuids     string
	original_string string
	marked_string   string
	modified_string string
}

package service_results

import (
	"core_foundation/core_object_model"
	"string_editor/object_model"
)

type FixChecksResults struct {
	core_object_model.Objects
	String_value_edit_history *object_model.StringEditHistories
	Issue_check_result_list   *IssueChecksResultLists
}

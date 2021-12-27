package service_results

import (
	"github.com/OntoLedgy/core_ontology/core_object_model"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

type FixChecksResults struct {
	core_object_model.Objects
	String_value_edit_history *object_model.StringEditHistories
	Issue_check_result_list   *IssueChecksResultLists
}

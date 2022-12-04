package service_results

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

type FixChecksResults struct {
	objects.BnogObjects
	StringValueEditHistory  *object_model.StringEditHistories
	Issue_check_result_list *IssueChecksResultLists
}

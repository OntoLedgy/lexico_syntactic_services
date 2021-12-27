package issues

import (
	"github.com/OntoLedgy/core_ontology/core_object_model"
)

type Issues struct {
	core_object_model.Objects
	Issue_type IssueTypes
}

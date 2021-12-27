package service_results

import (
	"github.com/OntoLedgy/string_editing_services/object_model"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/issues"
)

type IssueCheckResults struct {
	String_checks_issue *issues.Issues
	String_edit_history *object_model.StringEditHistories
}

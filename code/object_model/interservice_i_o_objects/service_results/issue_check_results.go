package service_results

import (
	"string_editor/object_model"
	"syntactic_checker/code/object_model/issues"
)

type IssueCheckResults struct {
	String_checks_issue *issues.Issues
	String_edit_history *object_model.StringEditHistories
}

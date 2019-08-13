package issues_processors

import (
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
)

type issuesProcessors struct {
	string_checks_parameter service_parameters.StringChecksParameters
	string_checks_issues    []issues.Issues
}

func (
	issues_processor *issuesProcessors) Set_string_check_issues() {

	in_scope_issue_types :=
		issues_processor.
			string_checks_parameter.
			In_scope_issue_types

	for _, in_scope_issue_type := range in_scope_issue_types {

		string_check_issue :=
			issues_processor.
				get_string_check_issue(
					in_scope_issue_type)

		issues_processor.
			process_string_check_issue(
				string_check_issue,
				in_scope_issue_type)
	}

}

func (
	issues_processor *issuesProcessors) Get_string_checks_issues() []issues.Issues {

	return issues_processor.string_checks_issues

}

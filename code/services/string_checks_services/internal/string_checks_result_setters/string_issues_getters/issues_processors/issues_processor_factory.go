package issues_processors

import (
	"syntactic_checker/code/object_model/service_parameters"
)

func Create(
	string_checks_parameter service_parameters.StringChecksParameters) *issuesProcessors {

	issues_processor :=
		new(
			issuesProcessors)

	issues_processor.
		string_checks_parameter =
		string_checks_parameter

	return issues_processor

}

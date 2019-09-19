package issues_processors

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
)

func Create(
	string_checks_parameter service_inputs.StringChecksInputs) *issuesProcessors {

	issues_processor :=
		new(
			issuesProcessors)

	issues_processor.
		string_checks_parameter =
		string_checks_parameter

	return issues_processor

}

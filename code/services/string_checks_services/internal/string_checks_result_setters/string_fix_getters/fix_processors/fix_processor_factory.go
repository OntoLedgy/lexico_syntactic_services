package fix_processors

import (
	"syntactic_checker/code/object_model/service_parameters"
)

func Create(
	string_checks_parameter service_parameters.StringChecksParameters) *FixProcessors {

	fix_processor := new(FixProcessors)

	fix_processor.
		string_checks_parameter =
		string_checks_parameter

	return fix_processor
}

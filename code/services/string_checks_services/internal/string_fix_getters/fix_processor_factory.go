package string_fix_getters

import "syntactic_checker/code/object_model/service_parameters"

type FixProcessorsFactory struct {
}

func (fix_processor_factory *FixProcessorsFactory) Create(
	string_checks_parameter service_parameters.StringChecksParameters) *fixProcessors {

	fix_processor := new(fixProcessors)

	fix_processor.
		string_value =
		string_checks_parameter.
			String_value

	fix_processor.issue_types =
		string_checks_parameter.
			In_scope_issue_types

	return fix_processor
}

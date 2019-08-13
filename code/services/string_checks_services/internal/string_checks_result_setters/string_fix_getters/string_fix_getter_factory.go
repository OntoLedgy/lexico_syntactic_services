package string_fix_getters

import "syntactic_checker/code/object_model/service_parameters"

type StringFixGetterFactory struct {
}

func (
	fix_processor_factory *StringFixGetterFactory) Create(
	string_checks_parameter service_parameters.StringChecksParameters) *StringFixGetters {

	string_fix_getter :=
		new(
			StringFixGetters)

	string_fix_getter.
		string_checks_parameter =
		string_checks_parameter

	return string_fix_getter
}

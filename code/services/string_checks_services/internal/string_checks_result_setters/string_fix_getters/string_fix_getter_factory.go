package string_fix_getters

import "syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"

type StringFixGetterFactory struct {
}

func (
	fix_processor_factory *StringFixGetterFactory) Create(
	string_checks_parameter service_inputs.StringChecksInputs) *StringFixGetters {

	string_fix_getter :=
		new(
			StringFixGetters)

	string_fix_getter.
		string_checks_input =
		string_checks_parameter

	return string_fix_getter
}

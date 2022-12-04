package fix_processors

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
)

func Create(
	string_checks_parameter service_inputs.StringChecksInputs) *FixProcessors {

	fix_processor := new(FixProcessors)

	fix_processor.
		stringChecksInput =
		string_checks_parameter

	return fix_processor
}

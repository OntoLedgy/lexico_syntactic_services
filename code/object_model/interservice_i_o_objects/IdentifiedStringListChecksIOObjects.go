package interservice_i_o_objects

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type IdentifiedStringListChecksIOObjects struct {
	Identified_string_list_checks_input   *service_inputs.IdentifiedStringListChecksInputs
	Identified_string_list_checks_results *service_results.IdentifiedStringListChecksResults
}

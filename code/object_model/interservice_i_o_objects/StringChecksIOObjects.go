package interservice_i_o_objects

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type StringCheckIOObjects struct {
	String_check_input  *service_inputs.StringCheckInputs
	String_check_result *service_results.StringCheckResults
}

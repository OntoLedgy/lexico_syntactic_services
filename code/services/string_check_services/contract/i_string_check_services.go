package contract

import (
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type IStringCheckServices interface {
	Set_string_check_result()
	Get_check_parameter() *service_inputs.StringCheckInputs
	Get_string_check_result() *service_results.StringCheckResults
	Set_string_check_result_value(*service_results.StringCheckResults)
}

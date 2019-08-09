package contract

import (
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/object_model/service_parameters"
)

type IStringCheckServices interface {
	Set_string_check_result()
	Set_string_check_result_value(check_result *check_results.CheckResults)
	Get_check_parameter() *service_parameters.StringCheckParameters
	Get_check_result() *check_results.CheckResults
}

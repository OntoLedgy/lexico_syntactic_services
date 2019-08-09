package contract

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
)

type IStringChecksServices interface {
	Set_string_checks_result()
	Get_string_checks_result() service_results.StringChecksResults
	Get_string_checks_parameter() service_parameters.StringChecksParameters
	Set_issues_result([]issues.Issues)
	Set_string_fixes_result(fixes.Fixes)
}

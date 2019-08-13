package service_results

import (
	"syntactic_checker/code/object_model/fixes"
	"syntactic_checker/code/object_model/issues"
)

type StringChecksResults struct {
	String_checks_issues []issues.Issues
	String_checks_fix    fixes.Fixes
}

package deprecate_cell_editors

import "syntactic_checker/code/object_model/issues"

func Get_replacement_string(issue_type issues.IssueTypes) string {

	var replacement_string string

	switch issue_type.Issue_check_replacement_string {

	case "STRING.EMPTY":
		replacement_string = ""

	case "SPACE":
		replacement_string = " "
		//TODO - Stage 2 - add other replacement string type cases
	}
	return replacement_string
}

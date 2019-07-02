package issues

type IssueTypes struct {
	Issue_type_uuid                string `json:"issue_type_uuid"`
	Issue_type_name                string `json:"issue_type_name"`
	Issue_check_type               string `json:"issue_check_type"`         //TODO - Stage 1 - move checks out from issues
	Issue_check_regex              string `json:"check_regex"`              //TODO - Stage 1 - move regex checks from checks
	Issue_check_replacement_string string `json:"check_replacement_string"` //

}

func (issue_type *IssueTypes) Get_replacement_string() string {

	replacement_string :=
		issue_type.
			get_replacement_string()

	return replacement_string
}

func (issue_type *IssueTypes) get_replacement_string() string {

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

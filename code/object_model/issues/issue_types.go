package issues

type (
	IssueTypes struct {
		Issue_type_uuid              string                `json:"issue_type_uuid"`
		Issue_type_name              string                `json:"issue_type_name"`
		Issue_severity_level         string                `json:"severity_level"`
		Issue_is_molecular           string                `json:"is_molecular"`
		Issue_check_type             string                `json:"issue_check_type"` //TODO - Stage 1 - move checks out from issues
		Issue_check_regex            string                `json:"check_regex"`      //TODO - Stage 1 - move regex checks from checks
		Issue_replacement_parameters ReplacementParameters `json:"replacement_parameters"`
	}
)

type ReplacementParameters struct {
	Is_replaceable                 string `json:"is_replaceable"`
	Issue_check_replacement_string string `json:"replacement_pattern"`
}

func (issue_type *IssueTypes) Get_replacement_string() string {

	replacement_string :=
		issue_type.
			get_replacement_string()

	return replacement_string
}

func (issue_type *IssueTypes) get_replacement_string() string {

	var replacement_string string

	switch issue_type.Issue_replacement_parameters.Issue_check_replacement_string {

	case "STRING.EMPTY":
		replacement_string = ""

	case "SPACE":
		replacement_string = " "

	case "UPPERCASE":
		replacement_string = "UPPERCASE"
		//TODO - Stage 2 - add other replacement string type cases
	}
	return replacement_string
}

type IssueTypesLists struct {
	issue_types_list []IssueTypes
}

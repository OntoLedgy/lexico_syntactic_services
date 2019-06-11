package object_model

type IssueTypes struct {
	Issue_type_uuid                string `json:"issue_type_uuid"`
	Issue_type_name                string `json:"issue_type_name"`
	Issue_check_type               string `json:"issue_check_type"`         //TODO - Stage 1 - move checks out from issues
	Issue_check_regex              string `json:"check_regex"`              //TODO - Stage 1 - move regex checks from checks
	Issue_check_replacement_string string `json:"check_replacement_string"` //

}

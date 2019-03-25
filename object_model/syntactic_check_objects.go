package object_model

import "database_manager/utils"

type Issue_types struct {
	Issue_type_uuid                string `json:"issue_type_uuid"`
	Issue_type_name                string `json:"issue_type_name"`
	Issue_check_type               string `json:"issue_check_type"`         //TODO - Stage 1 - move checks out from issues
	Issue_check_regex              string `json:"check_regex"`              //TODO - Stage 1 - move regex checks from checks
	Issue_check_replacement_string string `json:"check_replacement_string"` //

}

//TODO - Stage 1 - integrate this into the issue type structure

type Checks struct {
	uuids       utils.UUID
	check_names string
	check_type  Check_Types
}

type Check_Types struct {
	uuids            utils.UUID
	check_type_names string
}

type Non_parameterised_checks struct {
	check                    Checks
	check_regex_string       string
	check_replacement_string string
}

type Fix_transactions struct {
	check_uuids     string
	original_string string
	marked_string   string
	modified_string string
}

type Regex_check_results struct {
	Check_uuids         string
	Original_string     string
	Mark_string         string
	Replacement_string  string
	Regex_match_indices [][]int
}

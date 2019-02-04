package syntactic_checks

import "database_manager/utils"

type Checks struct {
	uuids            utils.UUID
	check_names      string
	check_type_uuids utils.UUID
}

type Check_Types struct {
	uuids            utils.UUID
	check_type_names string
}

type Non_parameterised_checks struct {
	check              Checks
	check_regex_string string
	check_type_level   string //level of pattern matching (substrings or main pattern)
}

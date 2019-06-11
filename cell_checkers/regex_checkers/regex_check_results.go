package regex_checkers

type RegexCheckResults struct {
	Check_uuids         string
	Original_string     string
	Mark_string         string //TODO - should not be part of regex check_results
	Replacement_string  string //TODO - should not be part of regex check_results
	Regex_match_indices [][]int
}

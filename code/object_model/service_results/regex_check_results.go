package service_results

type RegexCheckResults struct {
	Check_uuids         string
	Original_string     string
	Regex_match_indices [][]int
}

package cell_checkers

import (
	"syntactic_checker/cell_checkers/regex_checkers"
)

type ICellCheckers interface {
	CheckCell() *regex_checkers.RegexCheckResults
}

package cell_check_service

import (
	"syntactic_checker/code/services/cell_check_service/regex_checkers"
)

type ICellCheckers interface {
	CheckCell() *regex_checkers.RegexCheckResults
}

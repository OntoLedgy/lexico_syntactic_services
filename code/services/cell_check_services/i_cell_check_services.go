package cell_check_services

import (
	string_editor_object_model "string_editor/object_model"
	"syntactic_checker/code/object_model/check_results"
	"syntactic_checker/code/services/cell_check_services/internal/regex_checkers"
)

type ICellCheckService interface {
	Set_cell_check_result()
	Get_cell_regex_check_result() *regex_checkers.RegexCheckResults
	Get_check_result() *check_results.CheckResults
	Get_cell_edit_history() *string_editor_object_model.StringEditHistory
}

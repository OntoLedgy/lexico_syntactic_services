package cell_list_checks_service

type iCellListChecksService interface {
	Get_syntactic_check_results() map[string][][]string
}

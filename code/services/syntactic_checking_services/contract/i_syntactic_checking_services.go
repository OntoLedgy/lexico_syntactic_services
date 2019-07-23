package contract

import (
	"syntactic_checker/code/object_model/cells"
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters"
)

type ISyntacticCheckingServices interface {
	Run_syntactic_checking_service()
	Get_run_configuration() configuration_getters.RunConfigurations
	Set_syntactic_check_results(syntactic_checking_result service_results.CellListChecksResults)
	Get_in_scope_cell_list() cells.ListOfCells
	Get_syntactic_checking_result() service_results.CellListChecksResults
}

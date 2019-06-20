package syntactic_check_services

import (
	"syntactic_checker/code/object_model"
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/syntactic_check_services/configuration_handler"
)

type syntacticChecksService struct {
	run_configuration                    configuration_handler.Configurations
	in_scope_cell_list                   object_model.ListOfCells
	Syntactic_check_results_transactions transactions.SyntacticCheckTransactions
}

func (
	syntactic_checks_service *syntacticChecksService) Get_syntactic_check_run_configuration() configuration_handler.Configurations {

	return syntactic_checks_service.run_configuration

}
func (
	syntactic_checks_service *syntacticChecksService) Get_syntactic_check_transactions() transactions.SyntacticCheckTransactions {

	return syntactic_checks_service.Syntactic_check_results_transactions
}

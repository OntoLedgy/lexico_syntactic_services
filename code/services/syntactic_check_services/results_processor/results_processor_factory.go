package results_processor

import (
	"syntactic_checker/code/object_model/transactions"
	"syntactic_checker/code/services/syntactic_check_services/configuration_handler"
)

func Create_results_processor(
	syntactic_check_transactions transactions.SyntacticCheckTransactions,
	output_configuration configuration_handler.Output_configurations,
) *resultsProcessor {

	results_processor :=
		new(
			resultsProcessor)

	results_processor.
		syntactic_check_transactions =
		syntactic_check_transactions

	results_processor.
		output_configuration =
		output_configuration

	return results_processor

}

package checking_orchestrators

import (
	"syntactic_checker/code/services/syntactic_checking_services/contract"
	"syntactic_checker/code/services/syntactic_checking_services/internal/result_reporters"
	"syntactic_checker/code/services/syntactic_checking_services/internal/result_setters"
)

type SyntacticCheckingServiceOrchestrators struct {
	contract.ISyntacticCheckingServices
}

func (
	syntactic_checking_service_orchestrator *SyntacticCheckingServiceOrchestrators) Orchestrate_syntactic_checking() {

	syntactic_checking_service :=
		syntactic_checking_service_orchestrator.
			ISyntacticCheckingServices

	syntactic_checking_service_setter :=
		new(
			result_setters.SyntacticCheckingServicesResultSetters)

	syntactic_checking_service_setter.
		ISyntacticCheckingServices =
		syntactic_checking_service

	syntactic_checking_service_setter.
		Set_syntactic_checking_result()

	syntactic_checking_results_reporter :=
		new(
			result_reporters.SyntacticCheckingResultReporters)

	syntactic_checking_results_reporter.
		ISyntacticCheckingServices =
		syntactic_checking_service_setter.
			ISyntacticCheckingServices

	syntactic_checking_results_reporter.
		Report_syntactic_checking_result()
}

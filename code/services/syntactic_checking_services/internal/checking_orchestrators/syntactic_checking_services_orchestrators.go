package checking_orchestrators

import (
	"syntactic_checker/code/services/syntactic_checking_services/contract"
	"syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators/result_reporters"
	"syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators/result_setters"
)

type SyntacticCheckingServiceOrchestrators struct {
	contract.ISyntacticCheckingServices
}

func (
	syntactic_checking_service_orchestrator *SyntacticCheckingServiceOrchestrators) Orchestrate_syntactic_checking() {

	syntactic_checking_service_setter :=
		syntactic_checking_service_orchestrator.
			get_syntactic_checking_result() //TODO - should just return the result.

	syntactic_checking_service_orchestrator.
		report_syntactic_checking_result(
			syntactic_checking_service_setter) //TODO - should just pass the result
}

func (
	syntactic_checking_service_orchestrator *SyntacticCheckingServiceOrchestrators) get_syntactic_checking_result() *result_setters.SyntacticCheckingServicesResultSetters {

	syntactic_checking_service :=
		syntactic_checking_service_orchestrator.
			ISyntacticCheckingServices

	syntactic_checking_service_setter :=
		new(
			result_setters.
				SyntacticCheckingServicesResultSetters)

	syntactic_checking_service_setter.
		ISyntacticCheckingServices =
		syntactic_checking_service

	syntactic_checking_service_setter.
		Set_syntactic_checking_result()

	return syntactic_checking_service_setter
}

func (
	syntactic_checking_service_orchestrator *SyntacticCheckingServiceOrchestrators) report_syntactic_checking_result(
	syntactic_checking_service_setter *result_setters.SyntacticCheckingServicesResultSetters) {

	syntactic_checking_results_reporter :=
		new(
			result_reporters.
				SyntacticCheckingResultReporters)

	syntactic_checking_results_reporter.
		ISyntacticCheckingServices =
		syntactic_checking_service_setter.
			ISyntacticCheckingServices

	syntactic_checking_results_reporter.
		Report_syntactic_checking_result()
}

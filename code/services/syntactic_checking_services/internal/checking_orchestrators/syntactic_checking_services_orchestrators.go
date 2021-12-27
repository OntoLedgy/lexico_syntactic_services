package checking_orchestrators

import (
	"github.com/OntoLedgy/syntactic_checker/code/services/syntactic_checking_services/contract"
	"github.com/OntoLedgy/syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators/result_reporters"
	"github.com/OntoLedgy/syntactic_checker/code/services/syntactic_checking_services/internal/checking_orchestrators/result_setters"
)

type SyntacticCheckingServiceOrchestrators struct {
	contract.ISyntacticCheckingServices
}

func (
	syntactic_checking_service_orchestrator *SyntacticCheckingServiceOrchestrators) Orchestrate_syntactic_checking() {

	syntactic_checking_service_orchestrator.
		set_syntactic_checking_result()

	syntactic_checking_service_orchestrator.
		report_syntactic_checking_result()
}

func (
	syntactic_checking_service_orchestrator *SyntacticCheckingServiceOrchestrators) set_syntactic_checking_result() {

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

}

func (
	syntactic_checking_service_orchestrator *SyntacticCheckingServiceOrchestrators) report_syntactic_checking_result() {

	syntactic_checking_results_reporter :=
		new(
			result_reporters.
				SyntacticCheckingResultReporters)

	syntactic_checking_results_reporter.
		ISyntacticCheckingServices =
		syntactic_checking_service_orchestrator.
			ISyntacticCheckingServices

	syntactic_checking_results_reporter.
		Report_syntactic_checking_result()
}

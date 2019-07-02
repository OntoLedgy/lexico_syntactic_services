package syntactic_checking_services

type syntacticCheckingServicesOrchestrators struct {
	syntacticCheckingServices
}

func (
	syntactic_checking_service_orchestrator *syntacticCheckingServicesOrchestrators) orchestrate_syntactic_checking() {

	syntactic_checking_service :=
		&syntactic_checking_service_orchestrator.
			syntacticCheckingServices

	syntactic_checking_service_setter :=
		new(
			syntacticCheckingServicesResultSetters)

	syntactic_checking_service_setter.
		syntacticCheckingServices =
		*syntactic_checking_service

	syntactic_checking_service_setter.
		set_syntactic_checking_result()

	syntactic_checking_results_reporter :=
		new(
			syntacticCheckingResultReporters)

	syntactic_checking_results_reporter.
		syntacticCheckingServices =
		syntactic_checking_service_setter.
			syntacticCheckingServices

	syntactic_checking_results_reporter.
		report_syntactic_checking_result()
}

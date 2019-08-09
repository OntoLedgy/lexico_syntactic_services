package checking_orchestrators

import (
	"syntactic_checker/code/services/syntactic_checking_services/contract"
)

type SyntacticCheckingServiceOrchestratorFactory struct {
}

func (
	factory SyntacticCheckingServiceOrchestratorFactory) Create(
	syntactic_checking_service contract.ISyntacticCheckingServices) *SyntacticCheckingServiceOrchestrators {

	syntactic_checking_service_orchestrator :=
		new(
			SyntacticCheckingServiceOrchestrators)
	//pass the parameter into the constructor

	syntactic_checking_service_orchestrator.
		ISyntacticCheckingServices =
		syntactic_checking_service

	return syntactic_checking_service_orchestrator
}

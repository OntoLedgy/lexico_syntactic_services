package fix_processors

import (
	string_editor_object_model "github.com/OntoLedgy/string_editing_services/object_model"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type FixProcessors struct {
	stringChecksInput service_inputs.StringChecksInputs
	FixChecksResults  service_results.FixChecksResults
}

func (
	fixProcessor *FixProcessors) SetStringChecksFix() {

	fixProcessor.
		FixChecksResults.
		StringValueEditHistory =
		new(
			string_editor_object_model.
				StringEditHistories)

	fixProcessor.
		FixChecksResults.
		StringValueEditHistory.
		Create(fixProcessor.
			stringChecksInput.
			StringToCheck.StringValue)

	fixProcessor.FixChecksResults.
		ObjectUuid =
		fixProcessor.FixChecksResults.
			SetObjectUuid()

	fixProcessor.
		iterateThroughIssueTypes()

	fixProcessor.
		FixChecksResults.
		StringValueEditHistory =
		fixProcessor.
			FixChecksResults.
			StringValueEditHistory

}

func (
	fixProcessor *FixProcessors) iterateThroughIssueTypes() {

	issueTypes :=
		fixProcessor.
			stringChecksInput.
			IssueTypes

	for _, issueType := range issueTypes {

		fixProcessor.
			getStringCheckFix(
				issueType)

	}
}

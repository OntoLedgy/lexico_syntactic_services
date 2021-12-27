package fix_processors

import (
	string_editor_object_model "github.com/OntoLedgy/string_editing_services/object_model"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type FixProcessors struct {
	string_checks_input service_inputs.StringChecksInputs
	Fix_check_result    service_results.FixChecksResults
}

func (
	fix_processor *FixProcessors) Set_string_checks_fix() {

	fix_processor.
		Fix_check_result.
		String_value_edit_history =
		new(
			string_editor_object_model.
				StringEditHistories)

	fix_processor.
		Fix_check_result.
		String_value_edit_history.
		Create(fix_processor.
			string_checks_input.
			String_to_check.String_value)

	fix_processor.Fix_check_result.
		Object_uuid =
		fix_processor.Fix_check_result.
			Objects.
			Set_object_uuid()

	fix_processor.
		iterate_through_issue_types()

	fix_processor.
		Fix_check_result.
		String_value_edit_history =
		fix_processor.
			Fix_check_result.
			String_value_edit_history

}

func (
	fix_processor *FixProcessors) iterate_through_issue_types() {

	issue_types :=
		fix_processor.
			string_checks_input.
			Issue_types

	for _, issue_type := range issue_types {

		fix_processor.
			get_string_check_fix(
				issue_type)

	}
}

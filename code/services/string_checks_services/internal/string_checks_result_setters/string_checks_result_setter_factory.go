package string_checks_result_setters

import "github.com/OntoLedgy/syntactic_checker/code/services/string_checks_services/contract"

func Create(string_checks_service contract.IStringChecksServices) *StringChecksResultSetters {

	string_checks_result_setter :=
		new(StringChecksResultSetters)

	string_checks_result_setter.
		IStringChecksServices =
		string_checks_service

	return string_checks_result_setter
}

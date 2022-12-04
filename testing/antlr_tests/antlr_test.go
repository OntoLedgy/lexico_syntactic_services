package antlr_tests

import (
	"github.com/OntoLedgy/ol_common_services/code/services/operating_system_service"
	"testing"
)

func TestStandard(t *testing.T) {

	grammarFilePath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\syntactic_checker\\testing\\antlr_tests\\grammars\\golang\\GoParser.g4"

	targetGrammarCodeFolderPath := "S\\go\\src\\github.com\\OntoLedgy\\syntactic_checker\\testing\\antlr_tests\\LexerParserCode"

	antlrToolCommand := "java org.antlr.v4.Tool "
	antlrTestCommand := "java org.antlr.v4.gui.TestRig "

	var applicationToolRunner = operating_system_service.ApplicationRunner{
		CommandString:               "cmd",
		CommandArguments:            antlrToolCommand + grammarFilePath,
		CommandEnvironmentDrive:     "D:\\",
		CommandEnvironmentDirectory: targetGrammarCodeFolderPath}

	applicationToolRunner.RunCommand()

	grammarName := "GoLexer"

	var applicationTestRunner = operating_system_service.ApplicationRunner{
		CommandString:               "cmd",
		CommandArguments:            antlrTestCommand + grammarName,
		CommandEnvironmentDrive:     "D:\\",
		CommandEnvironmentDirectory: targetGrammarCodeFolderPath}

	applicationTestRunner.RunCommand()

}

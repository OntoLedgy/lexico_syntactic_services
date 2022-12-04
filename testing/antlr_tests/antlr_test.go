package antlr_tests

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/operating_system_service"
	parser "github.com/OntoLedgy/ol_common_services/testing/antlr_tests/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"testing"
)

func TestRunAntlrTool(t *testing.T) {

	lexerGrammarFilePath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\golang\\GoLexer.g4"

	parserGrammarFilePath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\golang\\GoParser.g4"

	targetGrammarCodeFolderPath := "S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\antlr_tests"

	antlrToolCommand := "java org.antlr.v4.Tool "

	targetLanguage := "Go"

	antlrLanguageParameter :=
		"-Dlanguage=" + targetLanguage

	delegatedGrammarFilePath :=
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\DelegateGrammars"

	delegatedGrammarParameter :=
		" -lib " + delegatedGrammarFilePath + " "

	outputFilePath :=
		"D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\golang\\parser"

	outputFolderPathParameter :=
		" -Xexact-output-dir -o " + outputFilePath + " "

	grammarFilePathParameter := lexerGrammarFilePath

	antlrLexerGeneratorCommand :=

		antlrToolCommand +
			antlrLanguageParameter +
			delegatedGrammarParameter +
			outputFolderPathParameter +
			grammarFilePathParameter

	var antlrToolRunnerLexer = operating_system_service.ApplicationRunner{
		CommandString:               "cmd",
		CommandArguments:            antlrLexerGeneratorCommand,
		CommandEnvironmentDrive:     "D:\\",
		CommandEnvironmentDirectory: targetGrammarCodeFolderPath}

	antlrToolRunnerLexer.RunCommand()

	grammarFilePathParameter = parserGrammarFilePath

	antlrParserGeneratorCommand :=

		antlrToolCommand +
			antlrLanguageParameter +
			delegatedGrammarParameter +
			outputFolderPathParameter +
			grammarFilePathParameter

	var antlrToolRunnerParser = operating_system_service.ApplicationRunner{
		CommandString:               "cmd",
		CommandArguments:            antlrParserGeneratorCommand,
		CommandEnvironmentDrive:     "D:\\",
		CommandEnvironmentDirectory: targetGrammarCodeFolderPath}

	antlrToolRunnerParser.RunCommand()

}

func TestAntlrGoLexer(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

func TestAntlrLexer(t *testing.T) {

	targetGrammarCodeFolderPath := "S\\go\\src\\github.com\\OntoLedgy\\syntactic_checker\\testing\\antlr_tests\\LexerParserCode"

	antlrTestCommand := "java org.antlr.v4.gui.TestRig "

	grammarName := "GoLexer"

	var applicationTestRunner = operating_system_service.ApplicationRunner{
		CommandString:               "cmd",
		CommandArguments:            antlrTestCommand + grammarName,
		CommandEnvironmentDrive:     "D:\\",
		CommandEnvironmentDirectory: targetGrammarCodeFolderPath}

	applicationTestRunner.RunCommand()

}

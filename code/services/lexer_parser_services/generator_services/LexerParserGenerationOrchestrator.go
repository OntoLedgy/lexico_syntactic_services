package generator_services

import (
	"github.com/OntoLedgy/ol_common_services/code/services/operating_system_service"
)

func GenerateLexerParserCode(
	lexerGrammarFilePath string,
	parserGrammarFilePath string,
	targetGrammarCodeFolderPath string,
	targetLanguage string) {

	lexerGrammarFilePath = "D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\golang\\GoLexer.g4"

	parserGrammarFilePath = "D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\golang\\GoParser.g4"

	targetGrammarCodeFolderPath = "S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\antlr_tests"

	antlrToolCommand := "java org.antlr.v4.Tool "

	targetLanguage = "Go"

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

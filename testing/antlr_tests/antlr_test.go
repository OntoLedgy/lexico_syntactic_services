package antlr_tests

import (
	"fmt"
	"github.com/OntoLedgy/ol_common_services/code/services/operating_system_service"
	"github.com/OntoLedgy/syntactic_checker/code/services/lexer_parser_services/generator_services"
	"github.com/OntoLedgy/syntactic_checker/testing/antlr_tests/grammars/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"testing"
)

func TestRunAntlrTool(t *testing.T) {

	lexerGrammarFilePath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\golang\\GoLexer.g4"

	parserGrammarFilePath := "D:\\S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\data\\grammars\\golang\\GoParser.g4"

	targetGrammarCodeFolderPath := "S\\go\\src\\github.com\\OntoLedgy\\lexico_syntactic_services\\testing\\antlr_tests"

	targetLanguage := "Go"

	generator_services.GenerateLexerParserCode(
		lexerGrammarFilePath,
		parserGrammarFilePath,
		targetGrammarCodeFolderPath,
		targetLanguage)

}

func TestAntlrGoLexer(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("import \"fmt\"")

	// Create the Lexer
	lexer := parser.NewGoLexer(is)

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

type goListener struct {
	*parser.BaseGoParserListener
	packageName string
}

func (goListenerInstance *goListener) EnterPackageClause(c *parser.PackageClauseContext) {

	c.GetText()

	fmt.Printf("entering package: %s", c.GetText())
}

func (goListenerInstance *goListener) ExitImportClause() {

	fmt.Printf("entering pacakge")
}

func TestAntlrGoParser(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("package test\n import \"fmt\"")

	// Create the Lexer
	goLexer := parser.NewGoLexer(is)

	stream := antlr.NewCommonTokenStream(goLexer, antlr.TokenDefaultChannel)

	goParser := parser.NewGoParser(stream)
	var listener goListener
	antlr.ParseTreeWalkerDefault.Walk(
		&listener,
		goParser.SourceFile())

	//fmt.Printf("parser output: \"%s \n\"", goParser,)

}

func TestAntlrTestRig(t *testing.T) {

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

package antlr_tests

import (
	"os"
	"os/exec"
)

func GenerateParserAndLexer(grammarFile string) error {
	// Execute the ANTLR tool to generate parser and lexer code based on the provided G4 grammar file
	cmd := exec.Command("antlr4", "-Dlanguage=Go", grammarFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

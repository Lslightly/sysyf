package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

/*
GenParseTree generate ParseTree from `inputFile`
*/
func GenParseTree(inputFile string) *SysYFParser {
	inputStream, err := antlr.NewFileStream(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open input file %s\n", inputFile)
		os.Exit(1)
	}
	lexer := NewSysYFLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewSysYFParser(stream)
	return parser
}

/*
PrintParseTree print ParseTree parsed from `inputFile` by `rule`
*/
func PrintParseTree(inputFile string, rule string) string {
	p := GenParseTree(inputFile)
	printer := NewParseTreePrinter()
	var tree antlr.ParseTree
	switch rule {
	case "Comp_unit":
		tree = p.Comp_unit()
	default:
		tree = p.Comp_unit()
	}
	return printer.Visit(tree).(string)
}

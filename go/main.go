package main

import (
	"SysYFCompiler/parser"
	"flag"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	flag.Parse()
	inputFile := flag.Args()[0]
	p := parser.GenParseTree(inputFile)
	parseTreePrinter := parser.ParseTreePrinter{
		BaseSysYFVisitor: parser.BaseSysYFVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		CurIndent: 0,
	}
	content := parseTreePrinter.Visit(p.Comp_unit())
	fmt.Println(content.(string))
}

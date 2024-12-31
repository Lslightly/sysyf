package ast

import (
	"SysYFCompiler/parser"

	"github.com/antlr4-go/antlr/v4"
)

func PrintAST(inputFile string, rule string) string {
	p := parser.GenParseTree(inputFile)
	var tree antlr.ParseTree
	switch rule {
	case "Comp_unit":
		tree = p.Comp_unit()
	default:
		tree = p.Comp_unit()
	}
	astBuilder := NewASTBuilder()
	ast := astBuilder.Visit(tree).(*Assembly)
	ctx := &Context{
		indentTimes: 0,
	}
	return ast.Print(ctx)
}

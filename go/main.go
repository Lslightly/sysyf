package main

import (
	"SysYFCompiler/ast"
	"SysYFCompiler/parser"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	inputFile := flag.Args()[0]
	p := parser.GenParseTree(inputFile)
	astBuilder := ast.Builder{}
	ast := astBuilder.Visit(p.Comp_unit()).(ast.Assembly)
	fmt.Println(ast.Print(nil))
}

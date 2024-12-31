package ast

import (
	"fmt"
	"strings"
)

type Context struct {
	indentTimes     int
	ExprWithBracket map[Expr]bool
}

func printIndent(times int) string {
	return strings.Repeat("    ", times)
}

func (node *Type) Print(ctx *Context) string {
	switch *node {
	case INT:
		return "int"
	case FLOAT:
		return "float"
	case VOID:
		return "void"
	default:
		panic("unreachable")
	}
}

func (node IntLit) Print(ctx *Context) string { return fmt.Sprintf("%d", node) }
func (node FloatLit) Print(ctx *Context) string {
	return fmt.Sprintf("%f", node)
}
func (node *BinaryExpr) Print(ctx *Context) string { panic("unreachable") }
func (node *UnaryExpr) Print(ctx *Context) string {
	var op string
	switch node.Op {
	case POS:
		op = "+"
	case NEG:
		op = "-"
	case NOT:
		op = "!"
	default:
		panic("unreachable")
	}
	if ctx.ExprWithBracket[node] {
		return fmt.Sprintf("(%s%s)", op, node.Exp.Print(ctx))
	}
	return fmt.Sprintf("%s%s", op, node.Exp.Print(ctx))
}
func (node *BracketExpr) Print(ctx *Context) string {
	return fmt.Sprintf("(%s)", node.Exp.Print(ctx))
}
func (node *EmptyExpr) Print(ctx *Context) string    { panic("unreachable") }
func (node *LVal) Print(ctx *Context) string         { panic("unreachable") }
func (node *FuncCallExpr) Print(ctx *Context) string { panic("unreachable") }
func (node *AssignStmt) Print(ctx *Context) string   { panic("unreachable") }
func (node *ExprStmt) Print(ctx *Context) string     { panic("unreachable") }
func (node *EmptyStmt) Print(ctx *Context) string    { panic("unreachable") }
func (node *VarDefStmt) Print(ctx *Context) string   { panic("unreachable") }
func (node *IfStmt) Print(ctx *Context) string       { panic("unreachable") }
func (node *WhileStmt) Print(ctx *Context) string    { panic("unreachable") }
func (node *BreakStmt) Print(ctx *Context) string    { panic("unreachable") }
func (node *ContinueStmt) Print(ctx *Context) string { panic("unreachable") }
func (node *ReturnStmt) Print(ctx *Context) string {
	format := "return %s;"
	if node.Exp != nil {
		return fmt.Sprintf(format, node.Exp.Print(ctx))
	} else {
		return fmt.Sprintf(format, "")
	}
}
func (node *BlockStmt) Print(ctx *Context) string {
	items := make([]string, 0, len(node.Stmts))
	ctx.indentTimes++
	for _, item := range node.Stmts {
		items = append(items, printIndent(ctx.indentTimes)+item.Print(ctx))
	}
	ctx.indentTimes--
	return fmt.Sprintf("{\n%s\n}", strings.Join(items, "\n"))
}
func (node *FuncParam) Print(ctx *Context) string {
	return ""
}
func (node *FuncDef) Print(ctx *Context) string {
	format := `%s %s(%s) %s`
	typ := node.Typ.Print(ctx)
	name := node.Name
	params := make([]string, 0, len(node.Params))
	for _, param := range node.Params {
		params = append(params, param.Print(ctx))
	}
	paramStr := strings.Join(params, ", ")
	return fmt.Sprintf(format, typ, name, paramStr, node.Body.Print(ctx))
}
func (node *Assembly) Print(ctx *Context) string {
	defStrs := make([]string, 0, len(node.Defs))
	for _, def := range node.Defs {
		defStrs = append(defStrs, def.Print(ctx))
	}
	return strings.Join(defStrs, "\n")
}

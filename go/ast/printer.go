package ast

import (
	"fmt"
	"strings"
)

type Context struct {
	indentTimes int
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
	return fmt.Sprintf("%f", node.Val)[:node.Len]
}
func (node *BinaryExpr) Print(ctx *Context) string {
	var op string
	switch node.Op {
	case OR:
		op = "||"
	case AND:
		op = "&&"
	case EQ:
		op = "=="
	case NEQ:
		op = "!="
	case LT:
		op = "<"
	case LTE:
		op = "<="
	case GT:
		op = ">"
	case GTE:
		op = ">="
	case ADD:
		op = "+"
	case SUB:
		op = "-"
	case MUL:
		op = "*"
	case DIV:
		op = "/"
	case MOD:
		op = "%"
	default:
		panic("unreachable")
	}
	return fmt.Sprintf("%s %s %s", node.Lhs.Print(ctx), op, node.Rhs.Print(ctx))
}
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
	return fmt.Sprintf("%s%s", op, node.Exp.Print(ctx))
}
func (node *BracketExpr) Print(ctx *Context) string {
	return fmt.Sprintf("(%s)", node.Exp.Print(ctx))
}
func (node *EmptyExpr) Print(ctx *Context) string { panic("unreachable") }
func (node *LVal) Print(ctx *Context) string {
	if node.ArrayIdx != nil {
		return fmt.Sprintf("%s[%s]", node.Name, node.ArrayIdx.Print(ctx))
	}
	return node.Name
}
func (node *FuncCallExpr) Print(ctx *Context) string { panic("unreachable") }
func (node *AssignStmt) Print(ctx *Context) string {
	format := "%s = %s;"
	return fmt.Sprintf(format, node.Lhs.Print(ctx), node.Rhs.Print(ctx))
}
func (node *ExprStmt) Print(ctx *Context) string  { panic("unreachable") }
func (node *EmptyStmt) Print(ctx *Context) string { panic("unreachable") }
func (node *VarDeclStmt) Print(ctx *Context) string {
	format := "%s %s;" // type
	if node.IsConst {
		format = "const " + format
	}
	defs := make([]string, 0, len(node.Defs))
	for _, def := range node.Defs {
		defs = append(defs, def.Print(ctx))
	}
	return fmt.Sprintf(format, node.Typ.Print(ctx), strings.Join(defs, ", "))
}
func (node *VarDef) Print(ctx *Context) string {
	format := "%s" // identity
	printInitValsArray := func(initVals []Expr) string {
		initFormat := "{%s}"
		inits := make([]string, 0, len(initVals))
		for _, init := range initVals {
			inits = append(inits, init.Print(ctx))
		}
		return fmt.Sprintf(initFormat, strings.Join(inits, ", "))
	}
	hasArrayLen := node.ArrayLen != nil
	hasInitVals := node.InitVals != nil
	if hasArrayLen {
		arrLen := ""
		if *node.ArrayLen != nil {
			arrLen = (*node.ArrayLen).Print(ctx)
		}
		if hasArrayLen {
			format += "[%s] = %s"
			return fmt.Sprintf(format, node.Name, arrLen, printInitValsArray(*node.InitVals))
		}
		format += "[%s]"
		return fmt.Sprintf(format, node.Name, arrLen)
	}
	if hasInitVals {
		format += " = %s"
		if len(*node.InitVals) != 1 {
			panic(fmt.Errorf("%s has more than one init vals", node.Name))
		}
		return fmt.Sprintf(format, node.Name, (*node.InitVals)[0].Print(ctx))
	}
	return fmt.Sprintf(format, node.Name)
}
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
	return fmt.Sprintf("{\n%s\n%s}", strings.Join(items, "\n"), printIndent(ctx.indentTimes))
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

package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

const INDENT = "    "

type ParseTreePrinter struct {
	BaseSysYFVisitor
	CurIndent int
}

func NewParseTreePrinter() *ParseTreePrinter {
	return &ParseTreePrinter{
		BaseSysYFVisitor: BaseSysYFVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		CurIndent: 0,
	}
}

func printIndent(times int) string {
	return strings.Repeat(INDENT, times)
}

func (v *ParseTreePrinter) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ParseTreePrinter) VisitOneNonTerminalChild(ctx *antlr.BaseParserRuleContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if childRule, ok := child.(antlr.RuleContext); ok {
			return v.Visit(childRule)
		}
	}
	panic("unreachable")
}

func (v *ParseTreePrinter) VisitComp_unit(ctx *Comp_unitContext) interface{} {
	globalDefs := make([]string, 0, len(ctx.AllGlobal_def()))
	for _, globalDef := range ctx.AllGlobal_def() {
		globalDefs = append(globalDefs, v.Visit(globalDef).(string))
	}
	return strings.Join(globalDefs, "\n")
}

func (v *ParseTreePrinter) VisitGlobal_def(ctx *Global_defContext) interface{} {
	return v.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

func (v *ParseTreePrinter) VisitConst_decl(ctx *Const_declContext) interface{} {
	typeStr := v.Visit(ctx.Btype()).(string)
	defStrs := make([]string, 0, len(ctx.AllConst_def()))
	for _, constDef := range ctx.AllConst_def() {
		defStrs = append(defStrs, v.Visit(constDef).(string))
	}
	return fmt.Sprintf("const %s %s;", typeStr, strings.Join(defStrs, ", "))
}

func (v *ParseTreePrinter) VisitBtype(ctx *BtypeContext) interface{} {
	if ctx.INT_T() != nil {
		return "int"
	}
	return "float"
}

func (v *ParseTreePrinter) VisitConst_def(ctx *Const_defContext) interface{} {
	ident := ctx.Ident().GetText()
	arrayLen := ""
	if ctx.LBRACKET() != nil {
		arrayLen = "[]"
		if ctx.INT() != nil {
			arrayLen = fmt.Sprintf("[%s]", ctx.INT().GetText())
		}
	}
	constInitVal := v.Visit(ctx.Const_init_val())
	return fmt.Sprintf("%s%s = %s", ident, arrayLen, constInitVal)
}

func (v *ParseTreePrinter) VisitSingle(ctx *SingleContext) interface{} {
	return v.Visit(ctx.Exp())
}

func (v *ParseTreePrinter) VisitArrayEmpty(ctx *ArrayEmptyContext) interface{} {
	return "{}"
}

func (v *ParseTreePrinter) VisitArrayMul(ctx *ArrayMulContext) interface{} {
	exps := make([]string, 0, len(ctx.AllExp()))
	for _, exp := range ctx.AllExp() {
		exps = append(exps, v.Visit(exp).(string))
	}
	return fmt.Sprintf("{ %s }", strings.Join(exps, ", "))
}

func (v *ParseTreePrinter) VisitVar_decl(ctx *Var_declContext) interface{} {
	btype := v.Visit(ctx.Btype()).(string)
	varDefs := make([]string, 0, len(ctx.AllVar_def()))
	for _, varDef := range ctx.AllVar_def() {
		varDefs = append(varDefs, v.Visit(varDef).(string))
	}
	return fmt.Sprintf("%s %s;", btype, strings.Join(varDefs, ", "))
}

func (v *ParseTreePrinter) VisitVar_def(ctx *Var_defContext) interface{} {
	ident := ctx.Ident().GetText()
	arrayLen := ""
	if ctx.LBRACKET() != nil {
		arrayLen = "[]"
		if ctx.INT() != nil {
			arrayLen = fmt.Sprintf("[%s]", ctx.INT().GetText())
		}
	}
	init := ""
	if ctx.Const_init_val() != nil {
		initVal := v.Visit(ctx.Const_init_val()).(string)
		init = fmt.Sprintf("= %s", initVal)
	}
	if init != "" {
		return fmt.Sprintf("%s%s %s", ident, arrayLen, init)
	}
	return fmt.Sprintf("%s%s", ident, arrayLen)
}

func (v *ParseTreePrinter) VisitFunc_def(ctx *Func_defContext) interface{} {
	btype := "void"
	if ctx.VOID_T() == nil {
		btype = v.Visit(ctx.Btype()).(string)
	}
	ident := ctx.Ident().GetText()
	fparams := make([]string, 0, len(ctx.AllFunc_fparam()))
	for _, fparam := range ctx.AllFunc_fparam() {
		fparams = append(fparams, v.Visit(fparam).(string))
	}
	fparamsStr := strings.Join(fparams, ", ")
	block := v.Visit(ctx.Block())
	return fmt.Sprintf("%s %s(%s) %s", btype, ident, fparamsStr, block)
}

func (v *ParseTreePrinter) VisitFunc_fparam(ctx *Func_fparamContext) interface{} {
	btype := v.Visit(ctx.Btype()).(string)
	ident := ctx.Ident().GetText()
	arrayLen := ""
	if ctx.LBRACKET() != nil {
		arrayLen = "[]"
	}
	return fmt.Sprintf("%s %s%s", btype, ident, arrayLen)
}

func (v *ParseTreePrinter) VisitBlock(ctx *BlockContext) interface{} {
	blockItems := make([]string, 0, len(ctx.AllBlock_item()))
	for _, blockItem := range ctx.AllBlock_item() {
		v.CurIndent += 1
		result := v.Visit(blockItem).(string)
		v.CurIndent -= 1
		blockItems = append(blockItems, printIndent(v.CurIndent+1)+result)
	}
	return fmt.Sprintf("{\n%s\n%s}", strings.Join(blockItems, "\n"), printIndent(v.CurIndent))
}

func (v *ParseTreePrinter) VisitBlock_item(ctx *Block_itemContext) interface{} {
	return v.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

func (v *ParseTreePrinter) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

func (v *ParseTreePrinter) VisitAssign_stmt(ctx *Assign_stmtContext) interface{} {
	lval := v.Visit(ctx.Lval())
	exp := v.Visit(ctx.Exp())
	return fmt.Sprintf("%s = %s;", lval, exp)
}

func (v *ParseTreePrinter) VisitEmpty_stmt(ctx *Empty_stmtContext) interface{} {
	return ";"
}

func (v *ParseTreePrinter) VisitExp_stmt(ctx *Exp_stmtContext) interface{} {
	return fmt.Sprintf("%s;", v.Visit(ctx.Exp()))
}

/*
stmt is not block:

	\n
	{indent}stmt

stmt is block:

	_{

	}
*/
func (v *ParseTreePrinter) printStmtWithIndent(stmt IStmtContext) (res string, isBlock bool) {
	if stmt.Block() == nil {
		isBlock = false
		v.CurIndent++
		res = v.Visit(stmt).(string)
		v.CurIndent--
		res = "\n" + printIndent(v.CurIndent+1) + res
	} else {
		isBlock = true
		res = " " + v.Visit(stmt).(string)
	}
	return
}

func (v *ParseTreePrinter) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	condExp := v.Visit(ctx.Cond_exp())
	ifBranch, ifIsBlock := v.printStmtWithIndent(ctx.Stmt(0))
	if len(ctx.AllStmt()) == 1 {
		/*
			if (cond) {

			}

			if (cond)\n
			{indent}stmt\n
		*/
		return fmt.Sprintf("if (%s)%s", condExp, ifBranch)
	}
	// else stmt
	/*
		 	if (cond) {

			} else if ...

			if (cond) {

			} else ...

			if (cond)\n
			{indent}stmt\n
			{indent}else
	*/
	var prefix, suffix string
	if ifIsBlock {
		prefix = " "
	} else {
		prefix = fmt.Sprintf("\n%s", printIndent(v.CurIndent))
	}
	if ctx.Stmt(1).If_stmt() != nil {
		suffix = fmt.Sprintf(" %s", v.Visit(ctx.Stmt(1)).(string))
	} else {
		suffix, _ = v.printStmtWithIndent(ctx.Stmt(1))
	}
	return fmt.Sprintf("if (%s)%s%selse%s", condExp, ifBranch, prefix, suffix)
}

func (v *ParseTreePrinter) VisitWhile_stmt(ctx *While_stmtContext) interface{} {
	condExp := v.Visit(ctx.Cond_exp())
	stmt, _ := v.printStmtWithIndent(ctx.Stmt())
	return fmt.Sprintf("while (%s)%s", condExp, stmt)
}

func (v *ParseTreePrinter) VisitBreak_stmt(ctx *Break_stmtContext) interface{} {
	return "break;"
}

func (v *ParseTreePrinter) VisitContinue_stmt(ctx *Continue_stmtContext) interface{} {
	return "continue;"
}

func (v *ParseTreePrinter) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	exp := ""
	if ctx.Exp() != nil {
		exp = v.Visit(ctx.Exp()).(string)
	}
	return fmt.Sprintf("return %s;", exp)
}

func (v *ParseTreePrinter) VisitCond_exp(ctx *Cond_expContext) interface{} {
	if ctx.Exp() != nil {
		return v.Visit(ctx.Exp())
	}
	condExps := make([]string, 0, len(ctx.AllCond_exp()))
	for _, condExp := range ctx.AllCond_exp() {
		condExps = append(condExps, v.Visit(condExp).(string))
	}
	return strings.Join(condExps, " "+ctx.GetChild(1).(antlr.TerminalNode).GetText()+" ")
}

func (v *ParseTreePrinter) VisitOtherexpAlt(ctx *OtherexpAltContext) interface{} {
	return v.Visit(ctx.Otherexp())
}

func (v *ParseTreePrinter) VisitOtherexp(ctx *OtherexpContext) interface{} {
	if ctx.Exp() != nil {
		return fmt.Sprintf("(%s)", v.Visit(ctx.Exp()).(string))
	}
	return v.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

func (v *ParseTreePrinter) VisitUnary(ctx *UnaryContext) interface{} {
	exp := v.Visit(ctx.Exp())
	return fmt.Sprintf("%s%s", ctx.GetChild(0).(antlr.TerminalNode).GetText(), exp)
}

func (v *ParseTreePrinter) VisitBinaryLow(ctx *BinaryLowContext) interface{} {
	left := v.Visit(ctx.Exp(0))
	right := v.Visit(ctx.Exp(1))
	return fmt.Sprintf("%s %s %s", left, ctx.GetChild(1).(antlr.TerminalNode).GetText(), right)
}

func (v *ParseTreePrinter) VisitBinaryHigh(ctx *BinaryHighContext) interface{} {
	left := v.Visit(ctx.Exp(0))
	right := v.Visit(ctx.Exp(1))
	return fmt.Sprintf("%s %s %s", left, ctx.GetChild(1).(antlr.TerminalNode).GetText(), right)
}

func (v *ParseTreePrinter) VisitFunc_call(ctx *Func_callContext) interface{} {
	ident := ctx.Ident().GetText()
	exps := make([]string, 0, len(ctx.AllExp()))
	for _, exp := range ctx.AllExp() {
		exps = append(exps, v.Visit(exp).(string))
	}
	return fmt.Sprintf("%s(%s)", ident, strings.Join(exps, ", "))
}

func (v *ParseTreePrinter) VisitLval(ctx *LvalContext) interface{} {
	ident := ctx.Ident().GetText()
	arrayLen := ""
	if ctx.LBRACKET() != nil {
		arrayLen = "[]"
		if ctx.Exp() != nil {
			arrayLen = fmt.Sprintf("[%s]", v.Visit(ctx.Exp()).(string))
		}
	}
	return fmt.Sprintf("%s%s", ident, arrayLen)
}

func (v *ParseTreePrinter) VisitIdent(ctx *IdentContext) interface{} {
	return ctx.GetText()
}

func (v *ParseTreePrinter) VisitNumber(ctx *NumberContext) interface{} {
	return ctx.GetText()
}

package ast

import (
	"SysYFCompiler/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

// Builder is a builder for AST.
type Builder struct {
}

func NewASTBuilder() *Builder {
	return &Builder{}
}

func (builder *Builder) VisitOneNonTerminalChild(ctx *antlr.BaseParserRuleContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if childRule, ok := child.(antlr.RuleContext); ok {
			return builder.Visit(childRule)
		}
	}
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#comp_unit.
func (builder *Builder) VisitComp_unit(ctx *parser.Comp_unitContext) interface{} {
	ast := &Assembly{}
	for _, gd := range ctx.AllGlobal_def() {
		ast.Defs = append(ast.Defs, builder.Visit(gd).(GlobalDef))
	}
	return ast
}

// Visit a parse tree produced by SysYFParser#global_def.
func (builder *Builder) VisitGlobal_def(ctx *parser.Global_defContext) interface{} {
	return builder.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

// Visit a parse tree produced by SysYFParser#const_decl.
func (builder *Builder) VisitConst_decl(ctx *parser.Const_declContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#btype.
func (builder *Builder) VisitBtype(ctx *parser.BtypeContext) interface{} {
	if ctx.INT_T() != nil {
		return INT
	}
	return FLOAT
}

// Visit a parse tree produced by SysYFParser#const_def.
func (builder *Builder) VisitConst_def(ctx *parser.Const_defContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#single.
func (builder *Builder) VisitSingle(ctx *parser.SingleContext) interface{} {
	return []Expr{builder.Visit(ctx.Exp()).(Expr)}
}

// Visit a parse tree produced by SysYFParser#arrayEmpty.
func (builder *Builder) VisitArrayEmpty(ctx *parser.ArrayEmptyContext) interface{} {
	return []Expr{}
}

// Visit a parse tree produced by SysYFParser#arrayMul.
func (builder *Builder) VisitArrayMul(ctx *parser.ArrayMulContext) interface{} {
	initExprs := make([]Expr, 0, len(ctx.AllExp()))
	for _, exp := range ctx.AllExp() {
		initExprs = append(initExprs, builder.Visit(exp).(Expr))
	}
	return initExprs
}

// Visit a parse tree produced by SysYFParser#var_decl.
func (builder *Builder) VisitVar_decl(ctx *parser.Var_declContext) interface{} {
	decl := &VarDeclStmt{}
	typ := builder.Visit(ctx.Btype()).(Type)
	defs := make([]*VarDef, 0, len(ctx.AllVar_def()))
	for _, def := range ctx.AllVar_def() {
		def := builder.Visit(def).(*VarDef)
		def.Parent = decl
		defs = append(defs, def)
	}
	decl.IsConst = false
	decl.Typ = typ
	decl.Defs = defs
	return decl
}

// Visit a parse tree produced by SysYFParser#var_def.
func (builder *Builder) VisitVar_def(ctx *parser.Var_defContext) interface{} {
	defStmt := &VarDef{}
	defStmt.Name = ctx.Ident().GetText()
	var arrLen *Expr = nil
	if ctx.LBRACKET() != nil {
		arrLen = new(Expr)
		if ctx.INT() != nil {
			val, err := strconv.Atoi(ctx.INT().GetText())
			if err != nil {
				panic(err)
			}
			*arrLen = IntLit(val)
		} else {
			*arrLen = nil
		}
	}
	defStmt.ArrayLen = arrLen
	defStmt.InitVals = nil
	if ctx.Const_init_val() != nil {
		initVals := builder.Visit(ctx.Const_init_val()).([]Expr)
		defStmt.InitVals = &initVals
	}
	return defStmt
}

// Visit a parse tree produced by SysYFParser#func_def.
func (builder *Builder) VisitFunc_def(ctx *parser.Func_defContext) interface{} {
	fn := &FuncDef{}
	typ := builder.Visit(ctx.Btype()).(Type)
	fn.Typ = typ
	fn.Name = ctx.Ident().GetText()
	fn.Params = make([]*FuncParam, 0, len(ctx.AllFunc_fparam()))
	for _, param := range ctx.AllFunc_fparam() {
		fn.Params = append(fn.Params, builder.Visit(param).(*FuncParam))
	}
	fn.Body = builder.Visit(ctx.Block()).(*BlockStmt)
	return fn
}

// Visit a parse tree produced by SysYFParser#func_fparam.
func (builder *Builder) VisitFunc_fparam(ctx *parser.Func_fparamContext) interface{} {
	param := &FuncParam{}
	param.Typ = builder.Visit(ctx.Btype()).(Type)
	param.Name = ctx.Ident().GetText()
	param.IsArray = ctx.LBRACKET() != nil
	return param
}

// Visit a parse tree produced by SysYFParser#block.
func (builder *Builder) VisitBlock(ctx *parser.BlockContext) interface{} {
	block := &BlockStmt{}
	for _, item := range ctx.AllBlock_item() {
		block.Stmts = append(block.Stmts, builder.Visit(item).(Stmt))
	}
	return block
}

// Visit a parse tree produced by SysYFParser#block_item.
func (builder *Builder) VisitBlock_item(ctx *parser.Block_itemContext) interface{} {
	return builder.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

// Visit a parse tree produced by SysYFParser#stmt.
func (builder *Builder) VisitStmt(ctx *parser.StmtContext) interface{} {
	return builder.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

// Visit a parse tree produced by SysYFParser#assign_stmt.
func (builder *Builder) VisitAssign_stmt(ctx *parser.Assign_stmtContext) interface{} {
	stmt := &AssignStmt{}
	stmt.Lhs = builder.Visit(ctx.Lval()).(*LVal)
	stmt.Rhs = builder.Visit(ctx.Exp()).(Expr)
	return stmt
}

// Visit a parse tree produced by SysYFParser#empty_stmt.
func (builder *Builder) VisitEmpty_stmt(ctx *parser.Empty_stmtContext) interface{} {
	return &EmptyStmt{}
}

// Visit a parse tree produced by SysYFParser#exp_stmt.
func (builder *Builder) VisitExp_stmt(ctx *parser.Exp_stmtContext) interface{} {
	return &ExprStmt{Exp: builder.Visit(ctx.Exp()).(Expr)}
}

// Visit a parse tree produced by SysYFParser#if_stmt.
func (builder *Builder) VisitIf_stmt(ctx *parser.If_stmtContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#while_stmt.
func (builder *Builder) VisitWhile_stmt(ctx *parser.While_stmtContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#break_stmt.
func (builder *Builder) VisitBreak_stmt(ctx *parser.Break_stmtContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#continue_stmt.
func (builder *Builder) VisitContinue_stmt(ctx *parser.Continue_stmtContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#return_stmt.
func (builder *Builder) VisitReturn_stmt(ctx *parser.Return_stmtContext) interface{} {
	if ctx.Exp() != nil {
		return &ReturnStmt{Exp: builder.Visit(ctx.Exp()).(Expr)}
	}
	return &ReturnStmt{Exp: nil}
}

// Visit a parse tree produced by SysYFParser#cond_exp.
func (builder *Builder) VisitCond_exp(ctx *parser.Cond_expContext) interface{} {
	if ctx.Exp() != nil {
		return builder.Visit(ctx.Exp())
	}
	op := GetBinaryOp(ctx.GetChild(1).(antlr.TerminalNode).GetText())
	return &BinaryExpr{
		Op:  op,
		Lhs: builder.Visit(ctx.Cond_exp(0)).(Expr),
		Rhs: builder.Visit(ctx.Cond_exp(1)).(Expr),
	}
}

// Visit a parse tree produced by SysYFParser#binaryHigh.
func (builder *Builder) VisitBinaryHigh(ctx *parser.BinaryHighContext) interface{} {
	op := GetBinaryOp(ctx.GetChild(1).(antlr.TerminalNode).GetText())
	return &BinaryExpr{
		Op:  op,
		Lhs: builder.Visit(ctx.Exp(0)).(Expr),
		Rhs: builder.Visit(ctx.Exp(1)).(Expr),
	}
}

// Visit a parse tree produced by SysYFParser#binaryLow.
func (builder *Builder) VisitBinaryLow(ctx *parser.BinaryLowContext) interface{} {
	op := GetBinaryOp(ctx.GetChild(1).(antlr.TerminalNode).GetText())
	return &BinaryExpr{
		Op:  op,
		Lhs: builder.Visit(ctx.Exp(0)).(Expr),
		Rhs: builder.Visit(ctx.Exp(1)).(Expr),
	}
}

// Visit a parse tree produced by SysYFParser#unary.
func (builder *Builder) VisitUnary(ctx *parser.UnaryContext) interface{} {
	exp := builder.Visit(ctx.Exp()).(Expr)
	if ctx.PLUS() != nil {
		return &UnaryExpr{Op: POS, Exp: exp}
	} else if ctx.MINUS() != nil {
		return &UnaryExpr{Op: NEG, Exp: exp}
	}
	return &UnaryExpr{Op: NOT, Exp: exp}
}

// Visit a parse tree produced by SysYFParser#otherexpAlt.
func (builder *Builder) VisitOtherexpAlt(ctx *parser.OtherexpAltContext) interface{} {
	return builder.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

// Visit a parse tree produced by SysYFParser#otherexp.
func (builder *Builder) VisitOtherexp(ctx *parser.OtherexpContext) interface{} {
	if ctx.Exp() != nil {
		exp := builder.Visit(ctx.Exp()).(Expr)
		return &BracketExpr{
			Exp: exp,
		}
	}
	return builder.VisitOneNonTerminalChild(&ctx.BaseParserRuleContext)
}

// Visit a parse tree produced by SysYFParser#func_call.
func (builder *Builder) VisitFunc_call(ctx *parser.Func_callContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#lval.
func (builder *Builder) VisitLval(ctx *parser.LvalContext) interface{} {
	lval := &LVal{}
	lval.Name = ctx.Ident().GetText()
	lval.ArrayIdx = nil
	if ctx.Exp() != nil {
		lval.ArrayIdx = builder.Visit(ctx.Exp()).(Expr)
	}
	return lval
}

// Visit a parse tree produced by SysYFParser#ident.
func (builder *Builder) VisitIdent(ctx *parser.IdentContext) interface{} {
	panic("unreachable")
}

// Visit a parse tree produced by SysYFParser#number.
func (builder *Builder) VisitNumber(ctx *parser.NumberContext) interface{} {
	if ctx.INT() != nil {
		val, err := strconv.Atoi(ctx.INT().GetText())
		if err != nil {
			panic(err)
		}
		return IntLit(val)
	}
	floatStr := ctx.FLOAT().GetText()
	val, err := strconv.ParseFloat(floatStr, 32)
	if err != nil {
		panic(err)
	}
	return FloatLit{
		Val: float32(val),
		Len: len(floatStr),
	}
}

func (builder *Builder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(builder)
}

func (v *Builder) VisitChildren(node antlr.RuleNode) interface{} {
	for _, tree := range node.GetChildren() {
		if tree, ok := (tree).(antlr.RuleNode); ok {
			tree.Accept(v)
		}
	}
	return nil
}
func (v *Builder) VisitTerminal(_ antlr.TerminalNode) interface{} { return nil }
func (v *Builder) VisitErrorNode(_ antlr.ErrorNode) interface{}   { return nil }

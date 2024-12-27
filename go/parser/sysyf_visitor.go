// Code generated from SysYF.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SysYF
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SysYFParser.
type SysYFVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SysYFParser#comp_unit.
	VisitComp_unit(ctx *Comp_unitContext) interface{}

	// Visit a parse tree produced by SysYFParser#global_def.
	VisitGlobal_def(ctx *Global_defContext) interface{}

	// Visit a parse tree produced by SysYFParser#const_decl.
	VisitConst_decl(ctx *Const_declContext) interface{}

	// Visit a parse tree produced by SysYFParser#btype.
	VisitBtype(ctx *BtypeContext) interface{}

	// Visit a parse tree produced by SysYFParser#const_def.
	VisitConst_def(ctx *Const_defContext) interface{}

	// Visit a parse tree produced by SysYFParser#single.
	VisitSingle(ctx *SingleContext) interface{}

	// Visit a parse tree produced by SysYFParser#arrayEmpty.
	VisitArrayEmpty(ctx *ArrayEmptyContext) interface{}

	// Visit a parse tree produced by SysYFParser#arrayMul.
	VisitArrayMul(ctx *ArrayMulContext) interface{}

	// Visit a parse tree produced by SysYFParser#var_decl.
	VisitVar_decl(ctx *Var_declContext) interface{}

	// Visit a parse tree produced by SysYFParser#var_def.
	VisitVar_def(ctx *Var_defContext) interface{}

	// Visit a parse tree produced by SysYFParser#func_def.
	VisitFunc_def(ctx *Func_defContext) interface{}

	// Visit a parse tree produced by SysYFParser#func_fparam.
	VisitFunc_fparam(ctx *Func_fparamContext) interface{}

	// Visit a parse tree produced by SysYFParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SysYFParser#block_item.
	VisitBlock_item(ctx *Block_itemContext) interface{}

	// Visit a parse tree produced by SysYFParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#assign_stmt.
	VisitAssign_stmt(ctx *Assign_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#empty_stmt.
	VisitEmpty_stmt(ctx *Empty_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#exp_stmt.
	VisitExp_stmt(ctx *Exp_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#while_stmt.
	VisitWhile_stmt(ctx *While_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#break_stmt.
	VisitBreak_stmt(ctx *Break_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#continue_stmt.
	VisitContinue_stmt(ctx *Continue_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by SysYFParser#cond_exp.
	VisitCond_exp(ctx *Cond_expContext) interface{}

	// Visit a parse tree produced by SysYFParser#binaryHigh.
	VisitBinaryHigh(ctx *BinaryHighContext) interface{}

	// Visit a parse tree produced by SysYFParser#binaryLow.
	VisitBinaryLow(ctx *BinaryLowContext) interface{}

	// Visit a parse tree produced by SysYFParser#unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by SysYFParser#otherexpAlt.
	VisitOtherexpAlt(ctx *OtherexpAltContext) interface{}

	// Visit a parse tree produced by SysYFParser#otherexp.
	VisitOtherexp(ctx *OtherexpContext) interface{}

	// Visit a parse tree produced by SysYFParser#func_call.
	VisitFunc_call(ctx *Func_callContext) interface{}

	// Visit a parse tree produced by SysYFParser#lval.
	VisitLval(ctx *LvalContext) interface{}

	// Visit a parse tree produced by SysYFParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by SysYFParser#number.
	VisitNumber(ctx *NumberContext) interface{}
}

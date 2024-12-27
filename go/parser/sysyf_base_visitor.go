// Code generated from SysYF.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SysYF
import "github.com/antlr4-go/antlr/v4"

type BaseSysYFVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSysYFVisitor) VisitComp_unit(ctx *Comp_unitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitGlobal_def(ctx *Global_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitConst_decl(ctx *Const_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitBtype(ctx *BtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitConst_def(ctx *Const_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitSingle(ctx *SingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitArrayEmpty(ctx *ArrayEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitArrayMul(ctx *ArrayMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitVar_decl(ctx *Var_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitVar_def(ctx *Var_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitFunc_def(ctx *Func_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitFunc_fparam(ctx *Func_fparamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitBlock_item(ctx *Block_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitAssign_stmt(ctx *Assign_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitEmpty_stmt(ctx *Empty_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitExp_stmt(ctx *Exp_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitWhile_stmt(ctx *While_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitBreak_stmt(ctx *Break_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitContinue_stmt(ctx *Continue_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitCond_exp(ctx *Cond_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitBinaryHigh(ctx *BinaryHighContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitBinaryLow(ctx *BinaryLowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitOtherexpAlt(ctx *OtherexpAltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitOtherexp(ctx *OtherexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitFunc_call(ctx *Func_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitLval(ctx *LvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSysYFVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

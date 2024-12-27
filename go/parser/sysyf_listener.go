// Code generated from SysYF.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SysYF
import "github.com/antlr4-go/antlr/v4"

// SysYFListener is a complete listener for a parse tree produced by SysYFParser.
type SysYFListener interface {
	antlr.ParseTreeListener

	// EnterComp_unit is called when entering the comp_unit production.
	EnterComp_unit(c *Comp_unitContext)

	// EnterGlobal_def is called when entering the global_def production.
	EnterGlobal_def(c *Global_defContext)

	// EnterConst_decl is called when entering the const_decl production.
	EnterConst_decl(c *Const_declContext)

	// EnterBtype is called when entering the btype production.
	EnterBtype(c *BtypeContext)

	// EnterConst_def is called when entering the const_def production.
	EnterConst_def(c *Const_defContext)

	// EnterSingle is called when entering the single production.
	EnterSingle(c *SingleContext)

	// EnterArrayEmpty is called when entering the arrayEmpty production.
	EnterArrayEmpty(c *ArrayEmptyContext)

	// EnterArrayMul is called when entering the arrayMul production.
	EnterArrayMul(c *ArrayMulContext)

	// EnterVar_decl is called when entering the var_decl production.
	EnterVar_decl(c *Var_declContext)

	// EnterVar_def is called when entering the var_def production.
	EnterVar_def(c *Var_defContext)

	// EnterFunc_def is called when entering the func_def production.
	EnterFunc_def(c *Func_defContext)

	// EnterFunc_fparam is called when entering the func_fparam production.
	EnterFunc_fparam(c *Func_fparamContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlock_item is called when entering the block_item production.
	EnterBlock_item(c *Block_itemContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssign_stmt is called when entering the assign_stmt production.
	EnterAssign_stmt(c *Assign_stmtContext)

	// EnterEmpty_stmt is called when entering the empty_stmt production.
	EnterEmpty_stmt(c *Empty_stmtContext)

	// EnterExp_stmt is called when entering the exp_stmt production.
	EnterExp_stmt(c *Exp_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterBreak_stmt is called when entering the break_stmt production.
	EnterBreak_stmt(c *Break_stmtContext)

	// EnterContinue_stmt is called when entering the continue_stmt production.
	EnterContinue_stmt(c *Continue_stmtContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterCond_exp is called when entering the cond_exp production.
	EnterCond_exp(c *Cond_expContext)

	// EnterBinaryHigh is called when entering the binaryHigh production.
	EnterBinaryHigh(c *BinaryHighContext)

	// EnterBinaryLow is called when entering the binaryLow production.
	EnterBinaryLow(c *BinaryLowContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterOtherexpAlt is called when entering the otherexpAlt production.
	EnterOtherexpAlt(c *OtherexpAltContext)

	// EnterOtherexp is called when entering the otherexp production.
	EnterOtherexp(c *OtherexpContext)

	// EnterFunc_call is called when entering the func_call production.
	EnterFunc_call(c *Func_callContext)

	// EnterLval is called when entering the lval production.
	EnterLval(c *LvalContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitComp_unit is called when exiting the comp_unit production.
	ExitComp_unit(c *Comp_unitContext)

	// ExitGlobal_def is called when exiting the global_def production.
	ExitGlobal_def(c *Global_defContext)

	// ExitConst_decl is called when exiting the const_decl production.
	ExitConst_decl(c *Const_declContext)

	// ExitBtype is called when exiting the btype production.
	ExitBtype(c *BtypeContext)

	// ExitConst_def is called when exiting the const_def production.
	ExitConst_def(c *Const_defContext)

	// ExitSingle is called when exiting the single production.
	ExitSingle(c *SingleContext)

	// ExitArrayEmpty is called when exiting the arrayEmpty production.
	ExitArrayEmpty(c *ArrayEmptyContext)

	// ExitArrayMul is called when exiting the arrayMul production.
	ExitArrayMul(c *ArrayMulContext)

	// ExitVar_decl is called when exiting the var_decl production.
	ExitVar_decl(c *Var_declContext)

	// ExitVar_def is called when exiting the var_def production.
	ExitVar_def(c *Var_defContext)

	// ExitFunc_def is called when exiting the func_def production.
	ExitFunc_def(c *Func_defContext)

	// ExitFunc_fparam is called when exiting the func_fparam production.
	ExitFunc_fparam(c *Func_fparamContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlock_item is called when exiting the block_item production.
	ExitBlock_item(c *Block_itemContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssign_stmt is called when exiting the assign_stmt production.
	ExitAssign_stmt(c *Assign_stmtContext)

	// ExitEmpty_stmt is called when exiting the empty_stmt production.
	ExitEmpty_stmt(c *Empty_stmtContext)

	// ExitExp_stmt is called when exiting the exp_stmt production.
	ExitExp_stmt(c *Exp_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitBreak_stmt is called when exiting the break_stmt production.
	ExitBreak_stmt(c *Break_stmtContext)

	// ExitContinue_stmt is called when exiting the continue_stmt production.
	ExitContinue_stmt(c *Continue_stmtContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitCond_exp is called when exiting the cond_exp production.
	ExitCond_exp(c *Cond_expContext)

	// ExitBinaryHigh is called when exiting the binaryHigh production.
	ExitBinaryHigh(c *BinaryHighContext)

	// ExitBinaryLow is called when exiting the binaryLow production.
	ExitBinaryLow(c *BinaryLowContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitOtherexpAlt is called when exiting the otherexpAlt production.
	ExitOtherexpAlt(c *OtherexpAltContext)

	// ExitOtherexp is called when exiting the otherexp production.
	ExitOtherexp(c *OtherexpContext)

	// ExitFunc_call is called when exiting the func_call production.
	ExitFunc_call(c *Func_callContext)

	// ExitLval is called when exiting the lval production.
	ExitLval(c *LvalContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}

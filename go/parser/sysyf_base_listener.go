// Code generated from SysYF.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SysYF
import "github.com/antlr4-go/antlr/v4"

// BaseSysYFListener is a complete listener for a parse tree produced by SysYFParser.
type BaseSysYFListener struct{}

var _ SysYFListener = &BaseSysYFListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSysYFListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSysYFListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSysYFListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSysYFListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterComp_unit is called when production comp_unit is entered.
func (s *BaseSysYFListener) EnterComp_unit(ctx *Comp_unitContext) {}

// ExitComp_unit is called when production comp_unit is exited.
func (s *BaseSysYFListener) ExitComp_unit(ctx *Comp_unitContext) {}

// EnterGlobal_def is called when production global_def is entered.
func (s *BaseSysYFListener) EnterGlobal_def(ctx *Global_defContext) {}

// ExitGlobal_def is called when production global_def is exited.
func (s *BaseSysYFListener) ExitGlobal_def(ctx *Global_defContext) {}

// EnterConst_decl is called when production const_decl is entered.
func (s *BaseSysYFListener) EnterConst_decl(ctx *Const_declContext) {}

// ExitConst_decl is called when production const_decl is exited.
func (s *BaseSysYFListener) ExitConst_decl(ctx *Const_declContext) {}

// EnterBtype is called when production btype is entered.
func (s *BaseSysYFListener) EnterBtype(ctx *BtypeContext) {}

// ExitBtype is called when production btype is exited.
func (s *BaseSysYFListener) ExitBtype(ctx *BtypeContext) {}

// EnterConst_def is called when production const_def is entered.
func (s *BaseSysYFListener) EnterConst_def(ctx *Const_defContext) {}

// ExitConst_def is called when production const_def is exited.
func (s *BaseSysYFListener) ExitConst_def(ctx *Const_defContext) {}

// EnterSingle is called when production single is entered.
func (s *BaseSysYFListener) EnterSingle(ctx *SingleContext) {}

// ExitSingle is called when production single is exited.
func (s *BaseSysYFListener) ExitSingle(ctx *SingleContext) {}

// EnterArrayEmpty is called when production arrayEmpty is entered.
func (s *BaseSysYFListener) EnterArrayEmpty(ctx *ArrayEmptyContext) {}

// ExitArrayEmpty is called when production arrayEmpty is exited.
func (s *BaseSysYFListener) ExitArrayEmpty(ctx *ArrayEmptyContext) {}

// EnterArrayMul is called when production arrayMul is entered.
func (s *BaseSysYFListener) EnterArrayMul(ctx *ArrayMulContext) {}

// ExitArrayMul is called when production arrayMul is exited.
func (s *BaseSysYFListener) ExitArrayMul(ctx *ArrayMulContext) {}

// EnterVar_decl is called when production var_decl is entered.
func (s *BaseSysYFListener) EnterVar_decl(ctx *Var_declContext) {}

// ExitVar_decl is called when production var_decl is exited.
func (s *BaseSysYFListener) ExitVar_decl(ctx *Var_declContext) {}

// EnterVar_def is called when production var_def is entered.
func (s *BaseSysYFListener) EnterVar_def(ctx *Var_defContext) {}

// ExitVar_def is called when production var_def is exited.
func (s *BaseSysYFListener) ExitVar_def(ctx *Var_defContext) {}

// EnterFunc_def is called when production func_def is entered.
func (s *BaseSysYFListener) EnterFunc_def(ctx *Func_defContext) {}

// ExitFunc_def is called when production func_def is exited.
func (s *BaseSysYFListener) ExitFunc_def(ctx *Func_defContext) {}

// EnterFunc_fparam is called when production func_fparam is entered.
func (s *BaseSysYFListener) EnterFunc_fparam(ctx *Func_fparamContext) {}

// ExitFunc_fparam is called when production func_fparam is exited.
func (s *BaseSysYFListener) ExitFunc_fparam(ctx *Func_fparamContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSysYFListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSysYFListener) ExitBlock(ctx *BlockContext) {}

// EnterBlock_item is called when production block_item is entered.
func (s *BaseSysYFListener) EnterBlock_item(ctx *Block_itemContext) {}

// ExitBlock_item is called when production block_item is exited.
func (s *BaseSysYFListener) ExitBlock_item(ctx *Block_itemContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseSysYFListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseSysYFListener) ExitStmt(ctx *StmtContext) {}

// EnterAssign_stmt is called when production assign_stmt is entered.
func (s *BaseSysYFListener) EnterAssign_stmt(ctx *Assign_stmtContext) {}

// ExitAssign_stmt is called when production assign_stmt is exited.
func (s *BaseSysYFListener) ExitAssign_stmt(ctx *Assign_stmtContext) {}

// EnterEmpty_stmt is called when production empty_stmt is entered.
func (s *BaseSysYFListener) EnterEmpty_stmt(ctx *Empty_stmtContext) {}

// ExitEmpty_stmt is called when production empty_stmt is exited.
func (s *BaseSysYFListener) ExitEmpty_stmt(ctx *Empty_stmtContext) {}

// EnterExp_stmt is called when production exp_stmt is entered.
func (s *BaseSysYFListener) EnterExp_stmt(ctx *Exp_stmtContext) {}

// ExitExp_stmt is called when production exp_stmt is exited.
func (s *BaseSysYFListener) ExitExp_stmt(ctx *Exp_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseSysYFListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseSysYFListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BaseSysYFListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BaseSysYFListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BaseSysYFListener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BaseSysYFListener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BaseSysYFListener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BaseSysYFListener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseSysYFListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseSysYFListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterCond_exp is called when production cond_exp is entered.
func (s *BaseSysYFListener) EnterCond_exp(ctx *Cond_expContext) {}

// ExitCond_exp is called when production cond_exp is exited.
func (s *BaseSysYFListener) ExitCond_exp(ctx *Cond_expContext) {}

// EnterBinaryHigh is called when production binaryHigh is entered.
func (s *BaseSysYFListener) EnterBinaryHigh(ctx *BinaryHighContext) {}

// ExitBinaryHigh is called when production binaryHigh is exited.
func (s *BaseSysYFListener) ExitBinaryHigh(ctx *BinaryHighContext) {}

// EnterBinaryLow is called when production binaryLow is entered.
func (s *BaseSysYFListener) EnterBinaryLow(ctx *BinaryLowContext) {}

// ExitBinaryLow is called when production binaryLow is exited.
func (s *BaseSysYFListener) ExitBinaryLow(ctx *BinaryLowContext) {}

// EnterUnary is called when production unary is entered.
func (s *BaseSysYFListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BaseSysYFListener) ExitUnary(ctx *UnaryContext) {}

// EnterOtherexpAlt is called when production otherexpAlt is entered.
func (s *BaseSysYFListener) EnterOtherexpAlt(ctx *OtherexpAltContext) {}

// ExitOtherexpAlt is called when production otherexpAlt is exited.
func (s *BaseSysYFListener) ExitOtherexpAlt(ctx *OtherexpAltContext) {}

// EnterOtherexp is called when production otherexp is entered.
func (s *BaseSysYFListener) EnterOtherexp(ctx *OtherexpContext) {}

// ExitOtherexp is called when production otherexp is exited.
func (s *BaseSysYFListener) ExitOtherexp(ctx *OtherexpContext) {}

// EnterFunc_call is called when production func_call is entered.
func (s *BaseSysYFListener) EnterFunc_call(ctx *Func_callContext) {}

// ExitFunc_call is called when production func_call is exited.
func (s *BaseSysYFListener) ExitFunc_call(ctx *Func_callContext) {}

// EnterLval is called when production lval is entered.
func (s *BaseSysYFListener) EnterLval(ctx *LvalContext) {}

// ExitLval is called when production lval is exited.
func (s *BaseSysYFListener) ExitLval(ctx *LvalContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseSysYFListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseSysYFListener) ExitIdent(ctx *IdentContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSysYFListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSysYFListener) ExitNumber(ctx *NumberContext) {}

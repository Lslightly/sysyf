package ast

import "fmt"

// BinaryOp
type BinaryOp int

// UnaryOp
type UnaryOp int

// Type
type Type int

// BinaryOp
const (
	OR BinaryOp = iota
	AND
	EQ
	NEQ
	LT
	LTE
	GT
	GTE
	ADD
	SUB
	MUL
	DIV
	MOD
	ALLBinaryOp
)

var BinaryOpStr = map[BinaryOp]string{
	OR:  "||",
	AND: "&&",
	EQ:  "==",
	NEQ: "!=",
	LT:  "<",
	LTE: "<=",
	GT:  ">",
	GTE: ">=",
	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	MOD: "%",
}

func GetBinaryOp(op string) BinaryOp {
	for i := range ALLBinaryOp {
		if BinaryOpStr[i] == op {
			return i
		}
	}
	panic(fmt.Sprintf("invalid binary op: %s", op))
}

// UnaryOp
const (
	_ UnaryOp = iota
	POS
	NEG
	NOT
)

// Type
const (
	_ Type = iota
	INT
	FLOAT
	VOID
)

// Expr

type IntLit int
type FloatLit struct {
	Val float32
	Len int
}

type BinaryExpr struct {
	Op       BinaryOp
	Lhs, Rhs Expr
}

type UnaryExpr struct {
	Op  UnaryOp
	Exp Expr
}

type EmptyExpr struct{}

type LVal struct {
	Name string
	/*
		ArrayIdx:
			nil means variable acc
	*/
	ArrayIdx Expr
}

type FuncCallExpr struct {
	FuncName string
	Args     []Expr
}

type BracketExpr struct {
	Exp Expr
}

// Stmt

type AssignStmt struct {
	Lhs *LVal
	Rhs Expr
}

type ExprStmt struct {
	Exp Expr
}

type EmptyStmt struct{}

type VarDeclStmt struct {
	IsConst bool
	Typ     Type
	Defs    []*VarDef
}
type VarDef struct {
	Parent *VarDeclStmt
	Name   string
	/*
		ArrayLen == nil means no ArrayLen
		*ArrayLen == nil means no len definition
	*/
	ArrayLen *Expr
	/*
		InitVals == nil means no init vals expr in AST
		*InitVals is the init vals
	*/
	InitVals *[]Expr
}

type IfStmt struct {
	Cond   Expr
	IfBody Stmt
	/*
		nil means no else body
	*/
	ElseBody Stmt
}

type WhileStmt struct {
	Cond Expr
	Body Stmt
}

type BreakStmt struct{}
type ContinueStmt struct{}

type ReturnStmt struct {
	/*
		nil means no return expr
	*/
	Exp Expr
}

type BlockStmt struct {
	Stmts []Stmt
}

type FuncParam struct {
	Typ     Type
	Name    string
	IsArray bool
}

type FuncDef struct {
	Typ    Type
	Name   string
	Params []*FuncParam
	Body   *BlockStmt
}

type Assembly struct {
	Defs []GlobalDef
}

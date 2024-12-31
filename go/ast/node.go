package ast

// BinaryOp
type BinaryOp int

// UnaryOp
type UnaryOp int

// Type
type Type int

// BinaryOp
const (
	_ BinaryOp = iota
	OR
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
)

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
type FloatLit float32

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

// Stmt

type AssignStmt struct {
	Lhs *LVal
	Rhs Expr
}

type ExprStmt struct {
	Exp Expr
}

type EmptyStmt struct{}

type VarDefStmt struct {
	IsConst bool
	Typ     Type
	Name    string
	/*
		ArrayLen == nil means no ArrayLen
	*/
	ArrayLen *Expr
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

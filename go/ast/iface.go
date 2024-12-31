package ast

type Printer interface {
	Print(ctx *Context) string
}

type GlobalDef interface {
	_gdef
	Printer
}

type Expr interface {
	_expr
	Printer
}

type Stmt interface {
	_stmt
	Printer
}

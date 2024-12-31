module SysYFCompiler

go 1.22.0

require github.com/antlr4-go/antlr/v4 v4.13.1

require (
	github.com/sergi/go-diff v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect
)

replace (
	github.com/antlr/antlr4 => github.com/antlr4-go/antlr/v4 v4.13.1
)
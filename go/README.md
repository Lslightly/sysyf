# SysYF Compiler (Go Version)

## Generating Parser

Execute following instructions first to generate parser. The guide is from [ANTLR go target](https://github.com/antlr/antlr4/blob/master/doc/go-target.md)

```bash
go generate ./...
```

## Test

[./test/testdata/cases/](./test/testdata/cases/) stores the original cases
[./test/testdata/parser](./test/testdata/parser/) stores correct answers of parse tree print.

use `go test ./...` to test.

use `go test -run TestParser` under `test` dir to test parser.

package test

import (
	"SysYFCompiler/ast"
	"SysYFCompiler/parser"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestCreateFile(t *testing.T) {
	__dir__ := filepath.Dir(currentFile())
	f, err := os.Create(filepath.Join(__dir__, "debug_out/Easy/main.sy"))
	if err != nil {
		t.Error(err)
	}
	fmt.Fprintln(f, "1")
}

func dumpDebugOut(content string, phase string, file string, t *testing.T) {
	__dir__ := filepath.Dir(currentFile())
	resFileDir := filepath.Join(__dir__, "debug_out", phase)
	resFilePath := filepath.Join(resFileDir, file)
	err := os.MkdirAll(filepath.Dir(resFilePath), os.ModePerm)
	if err != nil {
		t.Errorf("create dir failed: %s", err.Error())
	}
	resFile, err := os.Create(resFilePath)
	if err != nil {
		t.Errorf("create file failed: %s", err.Error())
	}
	defer resFile.Close()
	fmt.Fprintln(resFile, content)
}

func cmpResultAndAnswer(content, answer string, t *testing.T) {
	answer = strings.TrimSpace(answer)
	content = strings.TrimSpace(content)
	if answer != content {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(answer, content, true)
		if len(diffs) != 0 {
			t.Errorf("%s", dmp.DiffPrettyText(diffs))
		}
	}
}

func getAnswerSYContent(testcase TestCase, root string) string {
	ansCase := testcase
	ansCase.Root = root
	return ansCase.SYContent()
}

func TestParser(t *testing.T) {
	__dir__ := filepath.Dir(currentFile())
	parserTestDataRoot := filepath.Join(__dir__, "testdata/parser")
	for _, testcase := range TestCases {
		t.Run(testcase.Id, func(t *testing.T) {
			content := parser.PrintParseTree(testcase.SYPath(), "Comp_unit")
			dumpDebugOut(content, "parser", testcase.Id+".sy", t)
			cmpResultAndAnswer(content, getAnswerSYContent(testcase, parserTestDataRoot), t)
		})
	}
}

func TestASTBuilder(t *testing.T) {
	__dir__ := filepath.Dir(currentFile())
	parserTestDataRoot := filepath.Join(__dir__, "testdata/parser")
	for _, testcase := range TestCases {
		t.Run(testcase.Id, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("%s failed: %v", testcase.Id, err)
					t.Errorf("%s", debug.Stack())
				}
			}()
			content := ast.PrintAST(testcase.SYPath(), "Comp_unit")
			dumpDebugOut(content, "ast", testcase.Id+".sy", t)
			cmpResultAndAnswer(content, getAnswerSYContent(testcase, parserTestDataRoot), t)
		})
	}
}

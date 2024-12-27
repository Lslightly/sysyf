package test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/*
TestCase contains information of a test case.
Name: mode/Name
Root: the root path of the test case
Root/Name.sy is a testcase
*/
type TestCase struct {
	Id        string // mode/Name
	NeedInput bool
	Root      string // the root path of the test case
}

// InputPath returns the path of the .in file
func (t TestCase) InputPath() string {
	return filepath.Join(t.Root, t.Id+".in")
}

// OutputPath returns the path of the .out file
func (t TestCase) OutputPath() string {
	return filepath.Join(t.Root, t.Id+".out")
}

// SYPath returns the path of the .sy file
func (t TestCase) SYPath() string {
	return filepath.Join(t.Root, t.Id+".sy")
}

// SYContent returns the content of the .sy file
func (t TestCase) SYContent() string {
	content, _ := os.ReadFile(t.SYPath())
	return string(content)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}

// GenTestCases generates test cases from a root directory
func GenTestCases(root string) []TestCase {
	caseSet := make(map[string]TestCase) // dir/name -> TestCase
	result := make([]TestCase, 0)
	fs.WalkDir(os.DirFS(root), ".", func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() || p == "." {
			return nil
		}
		ext := filepath.Ext(p)
		caseid := strings.TrimSuffix(p, ext)
		caseid = strings.TrimPrefix(caseid, root)
		if _, ok := caseSet[caseid]; !ok {
			caseSet[caseid] = TestCase{
				Id:        caseid,
				NeedInput: fileExists(filepath.Join(root, caseid+".in")),
				Root:      root,
			}
		}
		return nil
	})
	for _, v := range caseSet {
		result = append(result, v)
	}
	return result
}

func currentFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// TestCases under cases dir
var TestCases []TestCase = GenTestCases(filepath.Join(filepath.Dir(currentFile()), "testdata/cases"))

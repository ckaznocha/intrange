package intrange_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/ckaznocha/intrange"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.RunWithSuggestedFixes(t, testdata, intrange.Analyzer, "main")
}

func FuzzAnalyzer(f *testing.F) {
	f.Add(`package p

func f() {
    // Basic for loop
    for i := 0; i < 10; i++ {
        println(i)
    }
}`)

	f.Add(`package p

func f() {
    // Loop with len
    arr := []int{1, 2, 3}
    for i := 0; i < len(arr); i++ {
        println(arr[i])
    }
}`)

	f.Add(`package p

func f() {
    // Already using range
    for i := range 10 {
        println(i)
    }
}`)

	f.Add(`package p

func f() {
    // Loop with different increment
    for i := 0; i < 10; i += 1 {
        println(i)
    }
}`)

	f.Add(`package p

func f() {
    // Loop with assignment increment
    for i := 0; i < 10; i = i + 1 {
        println(i)
    }
}`)

	f.Fuzz(func(t *testing.T, code string) {
		fSet := token.NewFileSet()

		f, err := parser.ParseFile(fSet, "test.go", code, parser.ParseComments)
		if err != nil {
			return
		}

		files := []*ast.File{f}
		typesInfo := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
			Uses:  make(map[*ast.Ident]types.Object),
		}

		pkg, err := (&types.Config{}).Check("p", fSet, files, typesInfo)
		if err != nil {
			return
		}

		if _, err := intrange.Analyzer.Run(&analysis.Pass{
			Fset:      fSet,
			Files:     files,
			Pkg:       pkg,
			TypesInfo: typesInfo,
			ResultOf: map[*analysis.Analyzer]any{
				inspect.Analyzer: inspector.New(files),
			},
			Report: func(analysis.Diagnostic) {},
		}); err != nil {
			t.Errorf("Analyzer failed on code:\n%s\nError: %v", code, err)
		}
	})
}

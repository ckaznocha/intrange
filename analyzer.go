package intrange

import (
	"flag"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

//nolint:gochecknoglobals // This is a global variable for the analyzer.
var Analyzer = &analysis.Analyzer{
	Name: "intrange",
	Doc: "intrange is a linter to find places where for loops could" +
		" make use of an integer range.",
	URL:      "https://github.com/ckaznocha/intrange",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Flags: flag.FlagSet{
		Usage: nil,
	},
	RunDespiteErrors: false,
	ResultType:       nil,
	FactTypes:        nil,
}

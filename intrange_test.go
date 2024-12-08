package intrange_test

import (
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"

	"github.com/ckaznocha/intrange"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.RunWithSuggestedFixes(t, testdata, intrange.Analyzer)
}

package varinit_test

import (
	"testing"

	"github.com/akif999/varinit"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, varinit.Analyzer, "a")
}

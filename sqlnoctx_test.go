package sqlnoctx_test

import (
	"testing"

	"github.com/barash-asenov/sqlnoctx"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, sqlnoctx.Analyzer, "a")
}

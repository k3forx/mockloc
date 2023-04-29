package mockloc_test

import (
	"fmt"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/k3forx/mockloc"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzerRun(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	res := analysistest.Run(t, testdata, mockloc.Analyzer, "a")
	fmt.Printf("res: %+v\n", res)
	// cases := map[string]struct{

	// }{}

	// for name, c := range cases {

	// }
}

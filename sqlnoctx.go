package sqlnoctx

import (
	"fmt"

	"github.com/barash-asenov/sqlnoctx/dbfunc"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var Analyzer = &analysis.Analyzer{
	Name:             "sqlnoctx",
	Doc:              "sqlnoctx finds executed db functions without context.Context",
	Run:              run,
	RunDespiteErrors: false,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
	ResultType: nil,
	FactTypes:  nil,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if _, err := dbfunc.Run(pass); err != nil {
		return nil, fmt.Errorf("run: %w", err)
	}

	return nil, nil
}

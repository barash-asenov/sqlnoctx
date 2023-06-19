package dbfunc

import (
	"fmt"
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

// only types with Context alternatives
var nonPreferredFuncNames = []string{
	"(*database/sql.DB).Exec",
	"(*database/sql.DB).Ping",
	"(*database/sql.DB).Prepare",
	"(*database/sql.DB).Query",
	"(*database/sql.DB).QueryRow",
}

func Run(pass *analysis.Pass) (interface{}, error) {
	nonPreferredFuncs := typeFuncs(pass, nonPreferredFuncNames)
	if len(nonPreferredFuncs) == 0 {
		return nil, nil
	}

	reportFuncs := dbCalledFuncs(pass, nonPreferredFuncs)
	report(pass, reportFuncs)

	return nil, nil
}

func dbCalledFuncs(pass *analysis.Pass, dbFuncs []*types.Func) []*Report {
	var reports []*Report

	ssa, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	if !ok {
		panic(fmt.Sprintf("%T is not *buildssa.SSA", pass.ResultOf[buildssa.Analyzer]))
	}

	for _, sf := range ssa.SrcFuncs {
		for _, b := range sf.Blocks {
			for _, instr := range b.Instrs {
				for _, dbFunc := range dbFuncs {
					if analysisutil.Called(instr, nil, dbFunc) {
						ngCalledFunc := &Report{
							Instruction: instr,
							function:    dbFunc,
						}
						reports = append(reports, ngCalledFunc)

						break
					}
				}
			}
		}
	}

	return reports
}

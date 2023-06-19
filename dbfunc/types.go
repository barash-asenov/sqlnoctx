package dbfunc

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

var errNotFound = fmt.Errorf("function not found")

func typeFuncs(pass *analysis.Pass, funcs []string) []*types.Func {
	fs := make([]*types.Func, 0, len(funcs))

	for _, fn := range funcs {
		f, err := typeFunc(pass, fn)
		if err != nil {
			continue
		}

		fs = append(fs, f)
	}

	return fs
}

// current types
// "(*database/sql.DB).Exec",
// "(*database/sql.DB).Ping",
// "(*database/sql.DB).Prepare",
// "(*database/sql.DB).Query",
// "(*database/sql.DB).QueryRow",
func typeFunc(pass *analysis.Pass, funcName string) (*types.Func, error) {
	ss := strings.Split(strings.TrimSpace(funcName), ".")

	if len(ss) != 3 {
		return nil, errNotFound
	}

	pkgname := strings.TrimLeft(ss[0], "(")
	typename := strings.TrimRight(ss[1], ")")

	if pkgname != "" && pkgname[0] == '*' {
		pkgname = pkgname[1:]
		typename = "*" + typename
	}

	typ := analysisutil.TypeOf(pass, pkgname, typename)
	if typ == nil {
		return nil, errNotFound
	}

	m := analysisutil.MethodOf(typ, ss[2])
	if m == nil {
		return nil, errNotFound
	}

	return m, nil
}

package main

import (
	"github.com/barash-asenov/sqlnoctx"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(sqlnoctx.Analyzer) }

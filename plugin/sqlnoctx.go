package main

import (
	"github.com/barash-asenov/sqlnoctx"
	"golang.org/x/tools/go/analysis"
)

// AnalyzerPlugin will be deprecated in the new version. New function shall be defined
// https://github.com/golangci/golangci-lint/pull/3887/fileshttps://github.com/golangci/golangci-lint/pull/3887/files
func New(_ any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{sqlnoctx.Analyzer}, nil
}

var AnalyzerPlugin analyzerPlugin

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{sqlnoctx.Analyzer}
}

package plugin

import (
	"github.com/OWEEN3/loglint/pkg/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type linter struct{}

func (l *linter) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (l *linter) GetLoadMode() string {
	return "types"
}

func New(conf any) (register.LinterPlugin, error) {
	return &linter{}, nil
}

func init() {
	register.Plugin("loglint", New)
}
package main

import (
	"github.com/torafugu2929/go-expr-finder/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.ContextComparisonAnalyzer)
}

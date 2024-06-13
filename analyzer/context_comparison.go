package analyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var ContextComparisonAnalyzer = &analysis.Analyzer{
	Name: "context_comparison",
	Doc:  "reports comparison of context.Context values",
	Run:  runContextComparisonAnalyzer,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func runContextComparisonAnalyzer(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.BinaryExpr:
			XType := pass.TypesInfo.TypeOf(n.X)
			YType := pass.TypesInfo.TypeOf(n.Y)

			if XType.String() == "context.Context" && YType.String() == "context.Context" {
				pass.Reportf(n.Pos(), "context comparison: %s", fmt.Sprint(n.X, n.Op, n.Y))
			}
		}
	})

	return nil, nil
}

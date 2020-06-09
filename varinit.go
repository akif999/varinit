package varinit

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "varinit is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "varinit",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if isHoge(n.Name) || isFuga(n.Name) {
				pass.Reportf(n.Pos(), "identifier is meaningless")
			}
		}
	})

	return nil, nil
}

func isHoge(s string) bool {
	return s == "hoge"
}

func isFuga(s string) bool {
	return s == "fuga"
}

package eval

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

// evalAst returns the value of ast expr for data
func evalAst(expr ast.Expr, data string) any {
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		x := evalAst(expr.X, data)
		y := evalAst(expr.Y, data)
		return Compute(expr.Op, x, y)
	case *ast.ParenExpr:
		return evalAst(expr.X, data)
	case *ast.BasicLit:
		return expr.Value
	}

	return fmt.Errorf("unsuported expr type %s", expr)
}

// parseFunc replaces function call to result
func parseFunc(expr ast.Expr, data string) error {
	astutil.Apply(expr, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.CallExpr:
			id, ok := x.Fun.(*ast.Ident)
			if ok {
				if result, err := Call(id.Name, x.Args, data); err == nil {
					c.Replace(&ast.BasicLit{
						Value: result,
					})
				} else {
					// TODO handle err
					return false
				}
			}
		}
		return true
	})
	return nil
}

// Eval returns the value of expr for data
func Eval(expr string, data string) (any, error) {

	astExpr, err := parser.ParseExpr(expr)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	ast.Print(fset, astExpr)

	if err := parseFunc(astExpr, data); err != nil {
		return nil, err
	}
	ast.Print(fset, astExpr)

	return evalAst(astExpr, data), nil
}

package easeeval

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/samanhappy/easeeval/compute"
	"github.com/samanhappy/easeeval/function"
	log "github.com/sirupsen/logrus"
	"golang.org/x/tools/go/ast/astutil"
)

// evalAst returns the value of ast expr for data
func evalAst(expr ast.Expr, data string) any {
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		x := evalAst(expr.X, data)
		y := evalAst(expr.Y, data)
		return compute.Compute(expr.Op, x, y)
	case *ast.ParenExpr:
		return evalAst(expr.X, data)
	case *ast.BasicLit:
		return expr.Value
	}

	return fmt.Errorf("unsuported expr type %T", expr)
}

// parseFunc replaces function call to result
func parseFunc(expr ast.Expr, data string) (e error) {
	astutil.Apply(expr, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.CallExpr:
			id, ok := x.Fun.(*ast.Ident)
			if ok {
				if result, err := function.Call(id.Name, x.Args, data); err == nil {
					c.Replace(&ast.BasicLit{
						Value: fmt.Sprintf("%v", result),
					})
				} else {
					e = err
				}
			}
		}
		return true
	})
	return
}

// Eval returns the value of expression for data
func Eval(expr string, data string) (any, error) {
	astExpr, err := parser.ParseExpr(expr)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	if log.GetLevel() == log.DebugLevel {
		log.Debugln("ast after expression parse is:")
		ast.Print(fset, astExpr)
	}

	if err := parseFunc(astExpr, data); err != nil {
		return nil, err
	}
	if log.GetLevel() == log.DebugLevel {
		log.Debugln("ast after function parse is:")
		ast.Print(fset, astExpr)
	}

	return evalAst(astExpr, data), nil
}

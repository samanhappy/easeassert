package eval

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

// toInt64 returns toInt64 value for any
func toInt64(x any) int64 {
	switch x := x.(type) {
	case int64:
		return x
	case string:
		i, _ := strconv.ParseInt(x, 10, 0)
		return i
	}
	return 0
}

// evalAst returns the value of ast expr for data
func evalAst(expr ast.Expr, data string) any {
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		x := evalAst(expr.X, data)
		y := evalAst(expr.Y, data)
		switch expr.Op {
		case token.LAND:
			return x.(bool) && y.(bool)
		case token.LOR:
			return x.(bool) || y.(bool)
		case token.GTR:
			return toInt64(x) > toInt64(y)
		case token.LSS:
			return toInt64(x) < toInt64(y)
		case token.ADD:
			return toInt64(x) + toInt64(y)
		case token.SUB:
			return toInt64(x) - toInt64(y)
		case token.MUL:
			return toInt64(x) * toInt64(y)
		case token.QUO:
			return toInt64(x) / toInt64(y)
		}
	case *ast.ParenExpr:
		return evalAst(expr.X, data)
	case *ast.BasicLit:
		return expr.Value
	}

	return fmt.Errorf("unsuported expr type %s", expr)
}

func parseFunc(expr ast.Expr, data string) (err error) {
	astutil.Apply(expr, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.CallExpr:
			id, ok := x.Fun.(*ast.Ident)
			if ok {
				v := ""
				switch id.Name {
				case "jq":
					v, err = ExtractJson(data, x.Args[0].(*ast.BasicLit).Value)
				case "regex":
					v, err = ExtractRegex(data, x.Args[0].(*ast.BasicLit).Value)
				}
				c.Replace(&ast.BasicLit{
					Value: v,
				})
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

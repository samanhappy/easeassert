package eval

import (
	"go/ast"
	"strings"

	jqparser "github.com/savaki/jq"
)

// Call calls function for name and args, then return function result
func Call(name string, args []ast.Expr, data string) (v string, err error) {
	switch name {
	case "jq":
		v, err = jq(data, args[0].(*ast.BasicLit).Value)
	case "regex":
		v, err = regex(data, args[0].(*ast.BasicLit).Value)
	}
	// TODO support other functions

	return v, err
}

// Jq querys value from json
func jq(data string, expr string) (string, error) {

	x := strings.Replace(expr, "\"", "", -1)

	op, err := jqparser.Parse(x)
	if err != nil {
		return "", err
	}

	value, err := op.Apply([]byte(data))
	if err != nil {
		return "", err
	}

	return strings.Replace(string(value), "\"", "", -1), nil
}

// regex gets result for text regex
func regex(data string, expr string) (string, error) {
	// TODO
	return "", nil
}

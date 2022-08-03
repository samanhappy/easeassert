package function

import (
	"fmt"
	"go/ast"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	jqparser "github.com/savaki/jq"
)

// Call calls function for name and args, then return function result
func Call(name string, exprs []ast.Expr, data string) (v any, err error) {
	args, err := parseArgs(exprs)
	if err != nil {
		return "", err
	}

	switch name {
	case "jq":
		v, err = jq(args, data)
	case "regex":
		v, err = regex(args, data)
	case "unixTime":
		v, err = unixTime(args, data)
	case "now":
		v, err = now(args, data)
	}
	// TODO support other functions

	return v, err
}

// Jq querys value from json
func jq(args []string, data string) (string, error) {
	if len(args) == 1 {
		op, err := jqparser.Parse(args[0])
		if err != nil {
			return "", err
		}

		value, err := op.Apply([]byte(data))
		if err != nil {
			return "", err
		}

		return strings.Replace(string(value), "\"", "", -1), nil
	}
	return "", fmt.Errorf("expected args size 1 but got %v", len(args))
}

// regex gets result for text regex
func regex(args []string, data string) (string, error) {
	// TODO
	return "", nil
}

// unixTime return Unix time seconds for time string
func unixTime(args []string, data string) (int64, error) {
	if len(args) == 1 {
		if t, err := dateparse.ParseAny(args[0]); err == nil {
			return t.Unix(), nil
		} else {
			return 0, err
		}
	}
	return 0, fmt.Errorf("expected args size 2 but got %v", len(args))
}

// now return Unix time seconds for now
func now(args []string, data string) (int64, error) {
	return time.Now().Unix(), nil
}

// parseArg parse arg to string value
func parseArgs(exprs []ast.Expr) ([]string, error) {
	args := make([]string, len(exprs))
	for i, expr := range exprs {
		if lit, ok := expr.(*ast.BasicLit); ok {
			args[i] = strings.Replace(lit.Value, "\"", "", -1)
		} else {
			return nil, fmt.Errorf("wrong expr type:%s", expr)
		}
	}
	return args, nil
}

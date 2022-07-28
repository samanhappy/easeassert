package eval

import (
	"strings"

	"github.com/savaki/jq"
)

func ExtractJson(data string, expr string) (string, error) {

	x := strings.Replace(expr, "\"", "", -1)

	op, err := jq.Parse(x)
	if err != nil {
		return "", err
	}

	value, err := op.Apply([]byte(data))
	if err != nil {
		return "", err
	}

	return strings.Replace(string(value), "\"", "", -1), nil
}

func ExtractRegex(data string, expr string) (string, error) {
	// TODO
	return "", nil
}

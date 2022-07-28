package eval

import (
	"strings"

	"github.com/savaki/jq"
)

func ExtractorJson(data string, expr string) string {
	x := strings.Replace(expr, "\"", "", -1)
	op, _ := jq.Parse(x)
	value, _ := op.Apply([]byte(data))
	return strings.Replace(string(value), "\"", "", -1)
}

func ExtractorRegex(data string, expr string) string {
	// TODO
	return ""
}

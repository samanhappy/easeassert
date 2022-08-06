package typecast

import (
	"go/token"
	"strconv"
)

// Cast converts types for different operators
func Cast(op token.Token, x any, y any) (any, any) {
	xc, yc, xok, yok := x, y, false, false
	switch op {
	case token.LAND, token.LOR, token.NOT:
		xc, xok = x.(bool)
		yc, yok = y.(bool)
	case token.ADD, token.SUB, token.MUL, token.QUO, token.GTR, token.LSS, token.EQL, token.NEQ, token.LEQ, token.GEQ:
		// TODO only support convert to int64 for arithmetic and relational operators
		xc, xok = toInt64(x)
		yc, yok = toInt64(y)
	}
	if xok && yok {
		return xc, yc
	}
	return x, y
}

// toInt64 returns toInt64 value for any
func toInt64(x any) (int64, bool) {
	switch x := x.(type) {
	case int:
		return int64(x), true
	case int8:
		return int64(x), true
	case int16:
		return int64(x), true
	case int32:
		return int64(x), true
	case int64:
		return x, true
	case string:
		i, _ := strconv.ParseInt(x, 10, 0)
		return i, true
	}
	return 0, false
}

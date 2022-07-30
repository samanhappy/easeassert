package eval

import (
	"go/token"
)

// Calculate get result for different operators and two params
func Calculate(op token.Token, x any, y any) any {
	x, y = Convert(op, x, y)
	switch op {
	case token.LAND:
		return x.(bool) && y.(bool)
	case token.LOR:
		return x.(bool) || y.(bool)
	case token.GTR:
		return x.(int64) > y.(int64)
	case token.LSS:
		return x.(int64) < y.(int64)
	case token.ADD:
		return x.(int64) + y.(int64)
	case token.SUB:
		return x.(int64) - y.(int64)
	case token.MUL:
		return x.(int64) * y.(int64)
	case token.QUO:
		return x.(int64) / y.(int64)
	}
	// TODO support other operators
	return nil
}

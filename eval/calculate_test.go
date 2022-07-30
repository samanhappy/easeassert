package eval

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	r := Calculate(token.ADD, 1, 2)
	assert.Equal(t, int64(3), r)
}

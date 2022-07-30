package eval

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	r := Compute(token.ADD, 1, 2)
	assert.Equal(t, int64(3), r)
}

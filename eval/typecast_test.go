package eval

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCast(t *testing.T) {
	x, y := Cast(token.LAND, true, false)
	assert.Equal(t, true, x)
	assert.Equal(t, false, y)

	x, y = Cast(token.ADD, "1", "2")
	assert.Equal(t, int64(1), x)
	assert.Equal(t, int64(2), y)
}

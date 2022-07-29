package eval

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	x, y := Convert(token.LAND, true, false)
	assert.Equal(t, true, x)
	assert.Equal(t, false, y)

	x, y = Convert(token.ADD, "1", "2")
	assert.Equal(t, int64(1), x)
	assert.Equal(t, int64(2), y)
}

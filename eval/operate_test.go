package eval

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperate(t *testing.T) {
	r := Operate(token.ADD, 1, 2)
	assert.Equal(t, int64(3), r)
}

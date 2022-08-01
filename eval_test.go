package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	r, err := Eval(`1 + jq(".k")`, `{"k":"2"}`)
	assert.Nil(t, err)
	assert.Equal(t, int64(3), r)
}

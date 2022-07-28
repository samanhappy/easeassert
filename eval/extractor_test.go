package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractorJson(t *testing.T) {
	v := ExtractorJson(`{"k":2}`, ".k")
	assert.Equal(t, "2", v)
}

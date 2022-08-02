package easeeval

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	r, err := Eval(`1 + jq(".k")`, `{"k":"2"}`)
	assert.Nil(t, err)
	assert.Equal(t, int64(3), r)
}

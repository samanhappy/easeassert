package easeeval

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	r, err := Eval(`1 + jq(".k")`, `{"k":"2"}`)
	assert.Nil(t, err)
	assert.Equal(t, int64(3), r)
}

func TestEvalTime(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	r, err := Eval(`unixTime(jq(".time")) < now()`, `{"time":"2022-08-01T12:00:00"}`)
	fmt.Println(r)
	assert.Nil(t, err)
	assert.Equal(t, true, r)
}

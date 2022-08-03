package function

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJq(t *testing.T) {
	v, e := jq([]string{".k"}, `{"k":2}`)
	assert.Nil(t, e)
	assert.Equal(t, "2", v)

	v, e = jq([]string{".data.[0].orderPayAt"}, `{
		"code": 0,
		"msg": "success",
		"data": [
			{
				"orderPayAt": "2022-07-28T20:55:06"
			}
		]
	}`)
	assert.Nil(t, e)
	assert.Equal(t, "2022-07-28T20:55:06", v)
}

func TestTime(t *testing.T) {
	v, err := unixTime([]string{"2022-08-01T12:00:00"}, "")
	assert.Nil(t, err)
	fmt.Println(v)
	assert.NotEqual(t, "0", v)
}

func TestNow(t *testing.T) {
	v, err := now([]string{}, "")
	assert.Nil(t, err)
	fmt.Println(v)
	assert.NotEqual(t, "0", v)
}

package dingjson

import (
	"github.com/dingdinglz/dingtools/dinglog"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	Id   int    `json:"id"`
	What string `json:"what"`
}

func TestJson(t *testing.T) {
	json, err := NewFromStruct(&Test{
		Id:   1,
		What: "yyyy",
	})
	assert.NoError(t, err)
	logger := dinglog.NewLogger()
	logger.Info(json.ConventToStr())
	assert.Equal(t, json.ConventToStr(), "{\"id\":1,\"what\":\"yyyy\"}")
	var jsonstruct Test
	err = json.ConventToStruct(&jsonstruct)
	assert.NoError(t, err)
	logger.Info(jsonstruct)
	assert.Equal(t, jsonstruct, Test{
		Id:   1,
		What: "yyyy",
	})
	what, err := json.GetString("what")
	assert.NoError(t, err)
	logger.Info(what)
	assert.Equal(t, what, "yyyy")
}

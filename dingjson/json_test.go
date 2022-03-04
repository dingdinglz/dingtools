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

type Test2 struct {
	Child Test   `json:"child"`
	Msg   string `json:"msg"`
}

type Test3 struct {
	Arr []string `json:"arr"`
}

type Test4 struct {
	What []Test `json:"what"`
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
	i, err := json.GetInt("id")
	assert.NoError(t, err)
	logger.Info(i)
	assert.Equal(t, i, int64(1))
	json2, err := NewFromStruct(&Test2{
		Test{
			Id:   1,
			What: "222",
		},
		"d",
	})
	assert.NoError(t, err)
	logger.Info(json2.ConventToStr())
	json3, err := json2.Get("child")
	assert.NoError(t, err)
	logger.Info(json3.ConventToStr())
	json4, err := NewFromStruct(&Test3{Arr: []string{"111", "222", "333"}})
	assert.NoError(t, err)
	json4.ArrayEach(func(jsonin *DingJson) {
		logger.Info(jsonin.ConventToStr())
	}, "arr")
	json5, err := NewFromStruct(&Test4{What: []Test{
		{1, "2222"},
		{2, "3333"},
	}})
	assert.NoError(t, err)
	json5.ArrayEach(func(jsonin2 *DingJson) {
		logger.Info(jsonin2.ConventToStr())
		logger.Info(jsonin2.GetString("what"))
	}, "what")
}

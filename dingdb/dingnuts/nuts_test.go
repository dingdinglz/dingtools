package dingnuts

import (
	"github.com/dingdinglz/dingtools/dinglog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDingNuts(t *testing.T) {
	logger := dinglog.NewLogger()
	dingNuts, err := NewDingNuts("./testdb")
	assert.NoError(t, err)
	assert.NotEqual(t, dingNuts, nil)

	err = dingNuts.SetValue([]byte("testname"), []byte("testvalue"))
	assert.NoError(t, err)
	value, err := dingNuts.GetValue([]byte("testname"))
	assert.NoError(t, err)
	assert.Equal(t, value, []byte("testvalue"))
	assert.Equal(t, string(value), "testvalue")
	logger.Info(string(value))

	testbucket := dingNuts.GetBucket("testbucket")
	err = testbucket.SetValue([]byte("test"), []byte("我是测试数据"))
	assert.NoError(t, err)
	value, err = testbucket.GetValue([]byte("test"))
	assert.NoError(t, err)
	assert.Equal(t, value, []byte("我是测试数据"))
	assert.Equal(t, string(value), "我是测试数据")
	logger.Info(string(value))

	err = dingNuts.BackUp("./backup")
	assert.NoError(t, err)

	err = testbucket.BackUp("./backup")
	assert.Error(t, err)

	err = dingNuts.Close()
	assert.NoError(t, err)
}

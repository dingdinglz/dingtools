package dinglog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	assert.Equal(t, version, Version())
	logger := NewLogger()
	logger.SetLevel(Level_Debug)
	logger.Info("I am Logger", "第二段", 5, "第三段")
}

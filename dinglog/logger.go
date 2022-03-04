package dinglog

import (
	"fmt"
	"github.com/gookit/color"
	"time"
)

const (
	version     string = "v0.1"
	Level_Debug int    = 1
	Level_Info  int    = 2
	Level_Warn  int    = 3
	Level_Error int    = 4
)

func init() {
	logger := NewLogger()
	logger.Info("dinglog", "version:", version)
}

// NewLogger 新建一个DingLogger组件
func NewLogger() *DingLogger {
	return &DingLogger{
		TimeFormat: "2006-01-02 15:04:05",
		Level:      Level_Info,
	}
}

// SetLevel 设置日志等级
//
// Level_Info info等级（推荐）
//
// Level_Debug debug等级（开发测试用）
//
// Level_Warn warn等级
//
// Level_Error error等级
func (l *DingLogger) SetLevel(level int) {
	l.Level = level
}

// Info 日志-info
func (l *DingLogger) Info(Objects ...interface{}) {
	if l.Level > Level_Info {
		return
	}
	timeStr := time.Now().Format(l.TimeFormat)
	show := "[" + color.Green.Text("info") + "]"
	var outs []interface{}
	outs = append(outs, timeStr)
	outs = append(outs, show)
	for _, i := range Objects {
		outs = append(outs, i)
	}
	fmt.Println(outs...)
}

type DingLogger struct {
	TimeFormat string
	Level      int
}

// Version 返回版本号
func Version() string {
	return version
}

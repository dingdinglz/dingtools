package dingruntime

import (
	"github.com/dingdinglz/dingtools/dinglog"
	"path"
	"runtime"
)

const (
	version string = "v0.1"
)

func init() {
	logger := dinglog.NewLogger()
	logger.Info("dingruntime", "version:", version)
}

// GetRunFuncName 获取当前的函数名称
func GetRunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// GetLocation 获取当前的文件名和行号
func GetLocation() (fileName string, line int) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fileName = path.Base(file)
	return
}

package dingjson

import "github.com/dingdinglz/dingtools/dinglog"

const (
	version string = "v0.1"
)

func init() {
	logger := dinglog.NewLogger()
	logger.Info("dingjson", "version:", version)
}

// Version 返回dingjson的版本号
func Version() string {
	return version
}

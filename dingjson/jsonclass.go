package dingjson

import (
	"encoding/json"
	"github.com/dingdinglz/dingtools/dinglog"
)

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

// NewFromStruct 通过struct生成dingjson类型（推荐）
func NewFromStruct(str interface{}) (*DingJson, error) {
	marshal, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}
	return &DingJson{src: marshal}, nil
}

// NewFromString 通过string生成dingjson类型
func NewFromString(str string) *DingJson {
	return &DingJson{src: []byte(str)}
}

// NewFromBytes 通过[]byte生成dingjson类型
func NewFromBytes(src []byte) *DingJson {
	return &DingJson{src: src}
}

type DingJson struct {
	src []byte
}

// ConventToStr 返回string的json
func (d *DingJson) ConventToStr() string {
	return string(d.src)
}

// GetSrc 返回[]byte类型的json
func (d *DingJson) GetSrc() []byte {
	return d.src
}

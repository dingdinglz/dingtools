package dingjson

import (
	"encoding/json"
	"github.com/buger/jsonparser"
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

// ConventToStruct 将json对象转换为struct类型
func (d *DingJson) ConventToStruct(v interface{}) error {
	return json.Unmarshal(d.src, v)
}

// GetString 取-string
func (d *DingJson) GetString(keys ...string) (string, error) {
	str, err := jsonparser.GetString(d.src, keys...)
	return str, err
}

// GetInt 取-int
func (d *DingJson) GetInt(keys ...string) (int64, error) {
	i, err := jsonparser.GetInt(d.src, keys...)
	return i, err
}

// Get 取子json
func (d *DingJson) Get(keys ...string) (*DingJson, error) {
	get, _, _, err := jsonparser.Get(d.src, keys...)
	return NewFromBytes(get), err
}

func (d *DingJson) ArrayEach(eachFunc func(json *DingJson), keys ...string) {
	jsonparser.ArrayEach(d.src, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		eachFunc(NewFromBytes(value))
	}, keys...)
}

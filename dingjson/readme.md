# dingjson

```sh
go get -u github.com/dingdinglz/dingtools/dingjson
```

## 生成一个json

```go
type Test struct {
	Id   int    `json:"id"`
	What string `json:"what"`
}

json, err := NewFromStruct(&Test{
		Id:   1,
		What: "yyyy",
	})
	
json.ConventToStr()
```

## 解析一个json

```go
json.GetInt()
json.GetString()
json.Get()
json.ArrayEach()
```

### ArrayEach示例

```go
type Test struct {
	Id   int    `json:"id"`
	What string `json:"what"`
}

type Test4 struct {
	What []Test `json:"what"`
}

json5, err := NewFromStruct(&Test4{What: []Test{
		{1, "2222"},
		{2, "3333"},
	}})
json5.ArrayEach(func(jsonin *DingJson) {
		logger.Info(jsonin.ConventToStr())
		logger.Info(jsonin.GetString("what"))
	}, "what")
```

## 完整示例请查看json_test.go文件
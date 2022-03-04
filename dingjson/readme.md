# dingjson

```sh
go get -u github.com/dingdinglz/dingtools/dingjson
```

## 注意：无论是生成还是解析，你都需要通过NewFrom***来生成一个dingjson型的对象。通过操作该对象，可以完成生成，解析，转换等操作

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

## 调用一个api，解析json，下载图片的例子

[随机p站图片下载](./_example)

## 完整示例请查看json_test.go文件
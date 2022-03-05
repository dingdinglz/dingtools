# dingnuts

封装自数据库[nutsdb](https://github.com/xujiajun/nutsdb)

```sh
go get -u github.com/dingdinglz/dingtools/dingdb/dingnuts
```

## 示例代码

[代码地址](./_example)

```go
package main

import (
	"github.com/dingdinglz/dingtools/dingdb/dingnuts"
	"github.com/dingdinglz/dingtools/dinglog"
	"time"
)

func main() {
	current := time.Now()
	logger := dinglog.NewLogger()
	logger.Info("启动数据库...")
	dingNuts, err := dingnuts.NewDingNuts("./testdb")
	if err != nil {
		logger.Error("数据库启动失败：", err.Error())
		return
	}
	err = dingNuts.SetValue([]byte("name"), []byte("value"))
	if err != nil {
		logger.Error("设置数据失败：", err.Error())
		return
	}
	value, err := dingNuts.GetValue([]byte("name"))
	if err != nil {
		logger.Error("查询数据失败：", err.Error())
		return
	}
	logger.Info("查询数据：", string(value))
	logger.Info("开新bucket存数据")
	testbucket := dingNuts.GetBucket("testbucket")
	err = testbucket.SetValue([]byte("name"), []byte("我在bucket里"))
	if err != nil {
		logger.Error("设置数据失败：", err.Error())
		return
	}
	value, err = testbucket.GetValue([]byte("name"))
	if err != nil {
		logger.Error("查询数据失败：", err.Error())
		return
	}
	logger.Info("查询数据：", string(value))
	logger.Info("正在备份数据库")
	err = dingNuts.BackUp("./backup")
	if err != nil {
		logger.Error("备份失败！", err.Error())
		return
	}
	sincetime := time.Since(current)
	logger.Info("共耗时：", sincetime.String())
}

```

## [完整文档](https://pkg.go.dev/github.com/dingdinglz/dingtools/dingdb/dingnuts)
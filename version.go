package dingtools

import "fmt"

var version string = "v0.1"

// Version 获取当前版本号
// 例如：var a = Version()
// fmt.Println(a)
func Version() string {
	return version
}

// OutVersion 输出版本号
func OutVersion()  {
	fmt.Println("dingtools",version)
}
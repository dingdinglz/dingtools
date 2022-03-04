// Package main
// @Description: dingjson的example，下面的程序会调用一个p站随机图片的api，解析该json并下载图片
package main

import (
	"github.com/dingdinglz/dingtools/dingjson"
	"github.com/dingdinglz/dingtools/dinglog"
	"io/ioutil"
	"net/http"
	"time"
)

var logger *dinglog.DingLogger

// GetImageUrl
// 通过解析json，返回url
func GetImageUrl() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.lolicon.app/setu/v2?proxy=o.i.edcms.pw", nil)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	logger.Info("json数据获取成功！")
	nowtime := time.Now()
	json := dingjson.NewFromBytes(bodyText)
	var res string = ""
	json.ArrayEach(func(json *dingjson.DingJson) {
		res, _ = json.GetString("urls", "original")
	}, "data")
	sincetime := time.Since(nowtime)
	logger.Info("耗时：", sincetime.String())
	logger.Info("取出图片url：", res)
	return res
}

// DownLoad 下载获取到的图片
func DownLoad(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	req.Header.Set("Referer", "https://pixivic.com/")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	ioutil.WriteFile("test.png", bodyText, 0666)
	logger.Info("下载成功！")
	return
}

func main() {
	logger = dinglog.NewLogger()
	url := GetImageUrl()
	DownLoad(url)
}

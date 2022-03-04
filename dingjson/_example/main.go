package main

import (
	"github.com/dingdinglz/dingtools/dingjson"
	"github.com/dingdinglz/dingtools/dinglog"
)

func main() {
	json, _ := dingjson.NewFromStruct(struct {
		Test string `json:"test"`
	}{"111"})
	logger := dinglog.NewLogger()
	logger.Info(json.ConventToStr())
}

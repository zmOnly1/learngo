package main

import (
	"learngo2/crawal/engine"
	"learngo2/crawal/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/",
		ParseFunc: parser.PrintCityList,
	})
}

package main

import (
	"learngo2/crawal/engine"
	"learngo2/crawal/scheduler"
	"learngo2/crawal/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/",
		ParseFunc: parser.ParseCityList,
	})
}

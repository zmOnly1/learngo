package main

import (
	"learngo2/crawal/engine"
	"learngo2/crawal/persist"
	"learngo2/crawal/scheduler"
	"learngo2/crawal/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine4{
		Scheduler:   &scheduler.QueueScheduler3{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/",
		ParseFunc: parser.ParseCityList,
	})
}

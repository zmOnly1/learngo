package main

import (
	"learngo2/crawal/engine"
	"learngo2/crawal/persist"
	"learngo2/crawal/scheduler"
	"learngo2/crawal/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine4{
		Scheduler:   &scheduler.QueueScheduler3{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/",
		ParseFunc: parser.ParseCityList,
	})
}

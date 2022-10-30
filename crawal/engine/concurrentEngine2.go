package engine

import (
	"log"
)

type ConcurrentEngine2 struct {
	Scheduler   Scheduler2
	WorkerCount int
}

type Scheduler2 interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine2) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker2(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	count := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", count, item)
			count++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker2(out chan ParseResult, s Scheduler2) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			parseResult, err := Worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}

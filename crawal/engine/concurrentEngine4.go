package engine

type ConcurrentEngine4 struct {
	Scheduler   Scheduler4
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler4 interface {
	ReadyNotifie4
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifie4 interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine4) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker4(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(item Item) {
				e.ItemChan <- item
			}(item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicate request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker4(in chan Request, out chan ParseResult, ready ReadyNotifie4) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			parseResult, err := Worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}

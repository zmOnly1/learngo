package scheduler

import (
	"learngo2/crawal/engine"
)

type SimpleScheduler3 struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler3) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler3) WorkerReady(requests chan engine.Request) {
}

func (s *SimpleScheduler3) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler3) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

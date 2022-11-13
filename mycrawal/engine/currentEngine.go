package engine

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

type CurrentEngine struct {
}

type Scheduler struct {
	requestChan chan int
	workerChan  chan chan int
}

func (s *Scheduler) summit(r int) {
	s.requestChan <- r
}

func (s *Scheduler) Ready(in chan int) {
	s.workerChan <- in
}

func (s *Scheduler) NewWorkChan() chan int {
	return make(chan int)
}

func (s *Scheduler) Run() chan bool {
	s.requestChan = make(chan int)
	s.workerChan = make(chan chan int)

	timeOut := make(chan bool)
	go func() {
		var requests []int
		var workers []chan int
		for {
			var activeWorker chan int
			var activeRequest int
			if len(requests) > 0 && len(workers) > 0 {
				activeRequest = requests[0]
				activeWorker = workers[0]
			}
			select {
			case r := <-s.requestChan:
				requests = append(requests, r)
			case w := <-s.workerChan:
				workers = append(workers, w)
			case activeWorker <- activeRequest:
				requests = requests[1:]
				workers = workers[1:]
			case <-time.After(1000 * time.Millisecond):
				log.Println("timeout: no task left")
				timeOut <- true
			}
		}
	}()
	return timeOut
}

func (e CurrentEngine) Run(seeds ...int) {
	out := make(chan string)
	scheduler := Scheduler{}
	timeout := scheduler.Run()
	go func() {
		<-timeout
		close(out)
	}()

	workerCnt := 5
	for i := 0; i < workerCnt; i++ {
		createWorker(scheduler, scheduler.NewWorkChan(), out)
	}
	for _, r := range seeds {
		log.Printf("3:%d\n", r)
		if isDuplicate(r) {
			continue
		}
		scheduler.summit(r)
	}
	cnt := 0
	for {
		parseResult, ok := <-out
		if !ok {
			log.Println("engine exit.")
			break
		}
		for i, r := range parseResult {
			log.Printf("got result:%d, %s\n", i, string(r))
			rr, _ := strconv.Atoi(string(r))

			if isDuplicate(rr) {
				//log.Printf("duplicate number:%d", rr)
				continue
			}
			scheduler.summit(rr)
			cnt++
		}
		if cnt > 10000 {
			break
		}
	}
}

var existedNum = make(map[int]bool)

func isDuplicate(r int) bool {
	if existedNum[r] {
		return true
	}
	existedNum[r] = true
	return false
}

func createWorker(scheduler Scheduler, in chan int, out chan string) {
	go func() {
		for {
			scheduler.Ready(in)
			r, ok := <-in
			if !ok {
				log.Println("worker exit.")
				break
			}
			result, err := myWorker(r)
			if err != nil {
				panic(err)
			}
			out <- result
		}
	}()
}

func myWorker(r int) (string, error) {
	return strconv.Itoa(r * rand.Intn(10)), nil
}

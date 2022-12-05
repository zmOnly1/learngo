package main

import (
	"context"
	"fmt"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dataChan := make(chan string)
	stopChan := make(chan struct{})
	tracker := &Tracker{
		ch:   dataChan,
		stop: stopChan,
	}
	go tracker.Run()
	go tracker.Event(ctx, "this event")
	go tracker.Event(ctx, "this event")
	go tracker.Event(ctx, "this event")
	time.Sleep(5 * time.Second)
	tracker.Shutdown(ctx)
}

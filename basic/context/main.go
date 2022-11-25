package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func main() {
	//contextTimeout()
	contextWithValue()
}

func contextWithValue() {
	ctx := context.WithValue(context.Background(), "trace_id", 12345)
	ret, ok := ctx.Value("trace_id").(int)
	fmt.Printf("ret:%d, ok: %v\n", ret, ok)
}

func contextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	go func() {
		resp, err2 := client.Do(req)
		pack := Result{r: resp, err: err2}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("timeout", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := io.ReadAll(res.r.Body)
		fmt.Println("server response: %s", out)
	}
}

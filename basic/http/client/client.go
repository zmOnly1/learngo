package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	//requestWithClientAndHeader()
	//requestWithDefaultClientAndHeader()

	//commonRequest()
	test()
}

func requestWithClientAndHeader() {
	request, err := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	request.Header.Add("User-Agent", "test-Agent")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(request)
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

func requestWithDefaultClientAndHeader() {
	request, err := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	request.Header.Add("User-Agent", "test-Agent")

	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

func commonRequest() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

func test() {
	var url = []string{
		"http://www.baidu.com",
		"http://google,com",
		"http://taobao.com",
	}

	for _, v := range url {
		c := http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, time.Second*2)
				},
			},
		}
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err: %v\n", v, err)
			continue
		}
		fmt.Printf("head succ, status: %v\n", resp.Status)
	}

}

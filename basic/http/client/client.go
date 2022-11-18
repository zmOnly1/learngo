package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	requestWithClientAndHeader()
	//requestWithDefaultClientAndHeader()

	//commonRequest()
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

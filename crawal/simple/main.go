package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	//e := determineEncoding(resp.Body)
	//newReader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	printCityList(all)

}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[\da-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, name, certain := charset.DetermineEncoding(bytes, "")
	fmt.Printf("name: %s, certain: %v\n", name, certain)
	return e
}

package main

import (
	"fmt"
	"learngo2/basic/interface/infra"
)

type Queue []interface{}

func main() {
	retriever := getRetriever()

	i := retriever.(infra.Retriever)
	fmt.Println(i.Get("https://www.baidu.com"))

	switch retriever.(type) {
	case infra.Retriever:
		fmt.Println("type is infra.Retriever")
	}
	fmt.Println(retriever.Get("https://www.baidu.com"))

}

type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(url string, form map[string]string)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) {
	s.Get("")
	s.Post("", map[string]string{})
}

func getRetriever() Retriever {
	return infra.Retriever{}
}

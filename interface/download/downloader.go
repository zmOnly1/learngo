package main

import (
	"fmt"
	"learngo2/interface/infra"
)

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

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return infra.Retriever{}
}

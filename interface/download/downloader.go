package main

import (
	"fmt"
	"learngo2/interface/infra"
)

func main() {
	retriever := getRetriever()
	fmt.Println(retriever.Get("https://www.baidu.com"))

}

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return infra.Retriever{}
}

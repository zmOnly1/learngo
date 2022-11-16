package main

import (
	"bufio"
	"fmt"
	"learngo2/sort/node"
	"os"
)

func main() {
	//mergeFile()
	mergeDemo()
}

func mergeFile() {
	const filename = "small.in"
	const count = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := node.RandomSource(count)
	writer := bufio.NewWriter(file)
	node.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = node.ReaderSource(bufio.NewReader(file), -1)
	maxPrintCnt := 100
	pn := 0
	for v := range p {
		fmt.Println(v)
		pn++
		if pn == maxPrintCnt {
			break
		}
	}
}

func mergeDemo() {
	p := node.Merge(
		node.InMemSort(
			node.ArraysSource(3, 2, 6, 7, 4)),
		node.InMemSort(
			node.ArraysSource(7, 4, 0, 3, 2, 13, 8)),
	)
	for val := range p {
		fmt.Println(val)
	}
}

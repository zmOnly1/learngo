package main

import (
	"bufio"
	"fmt"
	node2 "learngo2/basic/sort/node"
	"os"
	"strconv"
)

func main() {
	basename := "small"
	filenameIn := basename + ".in"
	filenameOut := basename + ".out"
	fileSize := 512
	chunkCount := 4
	//commonWork(filenameIn, fileSize, chunkCount, filenameOut)
	networkWork(filenameIn, fileSize, chunkCount, filenameOut)
}

func commonWork(filenameIn string, fileSize int, chunkCount int, filenameOut string) {
	p := createPipeline(filenameIn, fileSize, chunkCount)
	writeToFile(p, filenameOut)
	printFile(filenameOut)
}
func networkWork(filenameIn string, fileSize int, chunkCount int, filenameOut string) {
	p := createNetworkPipeline(filenameIn, fileSize, chunkCount)
	writeToFile(p, filenameOut)
	printFile(filenameOut)
	//time.Sleep(time.Hour)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := node2.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	node2.WriterSink(writer, p)
}

func createPipeline(filename string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	node2.Init()

	var sortResults []<-chan int
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := node2.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, node2.InMemSort(source))
	}
	return node2.MergeN(sortResults...)

}

func createNetworkPipeline(filename string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	node2.Init()

	var sortAddr []string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := node2.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		node2.NetworkSink(addr, node2.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	var sortResults []<-chan int
	for _, addr := range sortAddr {
		sortResults = append(sortResults, node2.NetworkSource(addr))
	}
	return node2.MergeN(sortResults...)

}

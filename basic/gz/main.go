package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func init() {
	fmt.Println("init..l.")
}

func main() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		panic(err)
	}
	r = bufio.NewReader(fz)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("gz:", line)
	}
}

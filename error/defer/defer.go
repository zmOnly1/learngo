package main

import (
	"bufio"
	"fmt"
	"learngo2/functional/fib"
	"os"
)

func main() {
	writeFile("fib.txt")
}

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

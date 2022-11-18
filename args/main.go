package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input conf path")
	flag.IntVar(&logLevel, "d", 0, "please input log level")

	flag.Parse()

	fmt.Println(confPath, logLevel)
	args := os.Args
	fmt.Println(len(args))
	fmt.Println(args) //main.exe -c c://abc -d 11]

}

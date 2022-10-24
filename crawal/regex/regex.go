package main

import (
	"fmt"
	"regexp"
)

const text = `My email is testabc@doc.com
 aaa123@doc.com.cn
 bbb123@doc.com
`

func main() {
	//commonReg()
	findAllMatch()
}

func commonReg() {
	re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	match := re.FindString(text)
	fmt.Println(match)
}
func findAllMatch() {
	re := regexp.MustCompile(`(\w+)@(\w+)(\.[\w.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}

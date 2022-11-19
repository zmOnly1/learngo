package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string `json:"username"`
	Age  int
}

func main() {
	s := &student{
		Name: "xiaoming",
		Age:  12,
	}
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	a := 10
	d, e := json.Marshal(a)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(d))

}

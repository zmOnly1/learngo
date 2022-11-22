package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

type Class struct {
	ClzName string
	P       Person
}

func main() {
	//simpleTemplate()
	//templateWith()
	foreach()
}

func foreach() {
	t, err := template.ParseFiles("basic/template/foreach.html")
	if err != nil {
		panic(err)
	}
	var list []Person
	p1 := Person{
		Name: "Mary1",
	}
	p2 := Person{
		Name: "Mary2",
	}
	p3 := Person{
		Name: "Mary3",
	}
	list = append(list, p1)
	list = append(list, p2)
	list = append(list, p3)
	if err = t.Execute(os.Stdout, list); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
func templateWith() {
	t, err := template.ParseFiles("basic/template/template_with.html")
	if err != nil {
		panic(err)
	}
	p := Person{
		Name: "Mary",
		Age:  31,
	}
	c := Class{
		ClzName: "class one",
		P:       p,
	}
	if err = t.Execute(os.Stdout, c); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

func simpleTemplate() {
	t, err := template.ParseFiles("basic/template/index.html")
	if err != nil {
		panic(err)
	}
	p := Person{
		Name: "Mary",
		Age:  31,
	}
	if err = t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//string()
	//mset()
	//hset()
	//expiry()
	list()
}

func list() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	_, err = c.Do("lpush", "book_list", "abc", "efg", 300)
	if err != nil {
		panic(err)
	}
	r, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func expiry() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		panic(err)
	}
}
func hset() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	_, err = c.Do("HSet", "books", "abc", 100)
	if err != nil {
		panic(err)
	}
	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func mset() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	_, err = c.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		panic(err)
	}
	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		panic(err)
	}
	for _, v := range r {
		fmt.Println(v)
	}
}

func string() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		panic(err)
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

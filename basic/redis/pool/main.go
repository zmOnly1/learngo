package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {

	pool := &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	conn := pool.Get()
	defer conn.Close()
	string(conn)
}

func string(c redis.Conn) {
	_, err := c.Do("Set", "abc", 100)
	if err != nil {
		panic(err)
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

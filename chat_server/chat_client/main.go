package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

func login(conn net.Conn) error {
	return nil
}

func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	err = login(conn)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			fmt.Println("exit, bye")
			return
		}
		_, err := conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Println("write data error")
			return
		}

	}

}

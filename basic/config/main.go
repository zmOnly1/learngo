package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	iniConf, err := config.NewConfig("ini", "basic/config/ini.conf")
	if err != nil {
		panic(err)
	}
	ip := iniConf.String("server::listen_ip")
	port, _ := iniConf.Int("server::listen_port")

	fmt.Printf("ip: %s, port: %d\n", ip, port)

}

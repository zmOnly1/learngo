package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./basic/logs/logcollect.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(configStr))
	logs.SetLogger("file", string(configStr))

	logs.Debug("user name: %s", "liming")
	logs.Trace("user name: %s", "xiaoming")
	logs.Warn("user name: %s", "xili")

}

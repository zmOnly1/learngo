package main

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
}
func (c *MainController) Index() {
	c.Ctx.WriteString("hello Index")
}

func (c *MainController) Html() {
	c.TplName = "index.html"
}
func (c *MainController) Json() {
	m := make(map[string]interface{})
	m["code"] = 200
	m["message"] = "success"
	c.Data["json"] = m
	c.ServeJSON(true)
}

func main() {
	web.Router("/", &MainController{})
	web.Router("/index", &MainController{}, "*:Index")
	web.Router("/html", &MainController{}, "*:Html")
	web.Router("/json", &MainController{}, "*:Json")
	web.Run()
}

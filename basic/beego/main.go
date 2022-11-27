package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}
func (this *MainController) Index() {
	this.Ctx.WriteString("hello Index")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Router("/index", &MainController{}, "*:Index")
	beego.Run()
}

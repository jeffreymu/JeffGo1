package main

import (
	"github.com/astaxie/beego"
)


type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Hello world, It's my web app fo Golang!")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}

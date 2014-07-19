package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego"
	this.Data["Email"] = "astaxie"
	this.TplNames = "index.tpl"
}

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	this.Data["json"] = `{h:1,b:2}`
	this.ServeJson()
}

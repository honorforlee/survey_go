package main

import (
	"github.com/astaxie/beego"
	_ "github.com/sydbfhwj/survey_go/routers"
)

func main() {
	beego.EnableAdmin = true
	beego.Run()
}

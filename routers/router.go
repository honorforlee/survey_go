package routers

import (
	"github.com/astaxie/beego"
	"github.com/sydbfhwj/survey_go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/launch", &controllers.RoomController{}, "get:IsServiceAlive;post:Launch")
	// beego.RESTRouter("/hello", &controllers.HelloController{})
}

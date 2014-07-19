package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sydbfhwj/survey_go/models"
)

type RoomController struct {
	beego.Controller
}

/**
	判断服务是否可用的接口
**/
func (this *RoomController) IsServiceAlive() {
	this.Ctx.Output.SetStatus(200)
	this.Ctx.Output.Body([]byte("service is alive"))
	return
}

func (this *RoomController) Launch() {
	launchParam := struct {
		IsPublic        bool
		MaxPlayerCount  int32
		ProtocolVersion string
		AppVersion      string
	}{}

	datas := make([]byte, 1024)
	dataLen, _ := this.Ctx.Request.Body.Read(datas)
	realData := datas[:dataLen]
	beego.Info(string(realData))

	if err := json.Unmarshal(realData, &launchParam); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	r := models.NewRoom(launchParam.IsPublic, launchParam.MaxPlayerCount)
	res := struct{ AccessCode string }{r.AccessCode}
	this.Data["json"] = res
	this.ServeJson()
}

func isValidSppVersion(protocolVersion string, appVersion string) bool {
	return true
}

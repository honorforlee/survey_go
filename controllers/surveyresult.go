package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	. "github.com/sydbfhwj/survey_go/models"
)

type CommitSurveyContoller struct {
	beego.Controller
}

func (this *CommitSurveyContoller) Post() {
	beego.Info("CommitSurveyContoller Post")

	surveyRes := &SurveyResult{}
	json.Unmarshal(this.Ctx.Input.RequestBody, surveyRes)

	this.Ctx.WriteString("Handle Json " + surveyRes.DeviceInfo.DeviceName)
}

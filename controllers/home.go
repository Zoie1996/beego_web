package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.Data["IsHome"] = true
	// 检查是否登录成功
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	var err error
	self.Data["Topics"], err = models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	self.TplName = "home.html"
}

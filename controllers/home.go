package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.Data["IsHome"] = true
	// 检查是否登录成功
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	self.TplName = "home.html"
}
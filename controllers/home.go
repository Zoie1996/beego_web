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
	cid := self.Input().Get("cid")
	if len(cid) == 0 {
		cid = "0"
	}

	self.Data["Topics"], err = models.GetAllTopics(cid, true)
	if err != nil {
		beego.Error(err)
	}
	self.TplName = "home.html"
	self.Data["categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

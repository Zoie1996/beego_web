package controllers

import (
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (self *TopicController) Get() {
	self.Data["IsTopic"] = true
	// 检查是否登录成功
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	self.TplName = "topic.html"
}

func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
}

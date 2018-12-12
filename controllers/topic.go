package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (self *TopicController) Get() {
	self.TplName = "topic.html"
	self.Data["IsTopic"] = true
	// 检查是否登录成功
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	var err error
	self.Data["Topics"], err = models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
}
func (self *TopicController) Post() {
	if !checkAccount(self.Ctx) {
		self.Redirect("/login", 302)
	}
	title := self.Input().Get("title")
	content := self.Input().Get("content")
	if len(title) == 0 || len(content) == 0 {
		self.Redirect("/topic/add", 302)
		return
	}
	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic", 302)
}

func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
}

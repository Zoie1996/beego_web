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
	// 检查是否登录
	// if !checkAccount(self.Ctx) {
	// 	self.Redirect("/login", 302)
	// }
	id := self.Input().Get("id")
	title := self.Input().Get("title")
	content := self.Input().Get("content")
	category := self.Input().Get("category")
	if len(title) == 0 || len(content) == 0 {
		self.Redirect("/topic/add", 302)
		return
	}
	err := models.AddTopic(id, title, content, category)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic", 302)
}

func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
	var err error
	self.Data["categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

func (self *TopicController) Del() {
	err := models.DelTopic(self.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic", 302)
	return
}

func (self *TopicController) View() {
	self.TplName = "topic_view.html"
	var err error
	self.Data["Topic"], err = models.GetTopic(self.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["TID"] = self.Ctx.Input.Param("0")
}
func (self *TopicController) Modify() {
	self.TplName = "topic_modify.html"
	Topic, err := models.GetTopicModify(self.Ctx.Input.Param("0"))

	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = Topic

	self.Data["TID"] = self.Ctx.Input.Param("0")
	categories, err1 := models.GetAllCategories()
	if err1 != nil {
		beego.Error(err1)
	}
	for _, category := range categories {
		if category.ID == Topic.Category.ID {
			self.Data["IsCategory"] = true
		} else {
			self.Data["IsCategory"] = false
		}
	}
	self.Data["categories"] = categories
}

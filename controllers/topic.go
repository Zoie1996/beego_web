package controllers

import (
	"myblog/models"
	"path"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

// Get 获取文章
func (self *TopicController) Get() {
	self.TplName = "topic.html"
	self.Data["IsTopic"] = true
	// 检查是否登录成功
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	var err error
	self.Data["Topics"], err = models.GetAllTopics("0", false)
	if err != nil {
		beego.Error(err)
	}
}

// Post 添加文章
func (self *TopicController) Post() {
	// 检查是否登录
	if !checkAccount(self.Ctx) {
		self.Redirect("/login", 302)
	}
	id := self.Input().Get("id")
	title := self.Input().Get("title")
	content := self.Input().Get("content")
	category := self.Input().Get("category")
	// 获取附件
	f, fh, err := self.GetFile("uploadname")
	if err != nil {
		beego.Error(err)
	}
	// 关闭文件

	var attachment string
	if fh != nil {
		attachment = fh.Filename
		beego.Info(attachment)
		err = self.SaveToFile("uploadname", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
		defer f.Close()
	}
	if len(title) == 0 || len(content) == 0 {
		self.Redirect("/topic/add", 302)
		return
	}
	err = models.AddTopic(id, title, content, category, attachment)
	if err != nil {
		beego.Error(err)
	}
	if len(id) == 0 {
		self.Redirect("/topic", 302)
	} else {
		self.Redirect("/topic/view/"+id, 302)

	}
}

// Add 添加文章
func (self *TopicController) Add() {
	self.TplName = "topic_add.html"
	var err error
	self.Data["categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

// Del 删除文章
func (self *TopicController) Del() {
	id := self.Input().Get("id")
	cid := self.Input().Get("cid")
	err := models.DelTopic(id, cid)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic", 302)
	return
}

// View 浏览文章
func (self *TopicController) View() {
	self.TplName = "topic_view.html"
	tid := self.Ctx.Input.Param("0")
	var err error
	self.Data["Topic"], err = models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["replies"], err = models.GetReplies(tid)
	self.Data["IsLogin"] = checkAccount(self.Ctx)

}

// Modify 文章详情
func (self *TopicController) Modify() {
	self.TplName = "topic_modify.html"
	Topic, err := models.GetTopicModify(self.Ctx.Input.Param("0"))

	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = Topic
	// self.Data["TID"] = self.Ctx.Input.Param("0")
	categories, err1 := models.GetAllCategories()
	if err1 != nil {
		beego.Error(err1)
	}
	self.Data["categories"] = categories
}

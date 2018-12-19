package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (self *ReplyController) Add() {
	tid := self.Input().Get("tid")
	name := self.Input().Get("nikename")
	content := self.Input().Get("content")
	err := models.AddReply(tid, name, content)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic/view/"+tid, 302)
}

func (self *ReplyController) Delete() {
	if !checkAccount(self.Ctx) {
		return
	}
	id := self.Input().Get("rid")
	tid := self.Input().Get("tid")
	err := models.DeleteReply(id,tid)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic/view/"+tid, 302)

}

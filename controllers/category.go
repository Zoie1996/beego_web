package controllers

import (
	"myblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (self *CategoryController) Get() {
	op := self.Input().Get("op")
	switch op {
	case "add":
		name := self.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err == nil {
			// 如果添加成功，重定向到首页
			self.Redirect("/category", 301)
		}
	case "del":
		id := self.Input().Get("id")
		if len(id) == 0 {
			break
		}

	}
	self.Data["IsTopic"] = true
	// 检查是否登录成功
	self.Data["IsLogin"] = checkAccount(self.Ctx)
	self.TplName = "category.html"
	var err error
	self.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

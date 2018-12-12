package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	// 退出时删除cookie，跳转到首页
	IsExit := self.Input().Get("exit") == "true"
	if IsExit {
		self.Ctx.SetCookie("username", "", -1, "/")
		self.Ctx.SetCookie("password", "", -1, "/")
		self.Redirect("/", 301)
		return
	}
	self.TplName = "login.html"
}
func (self *LoginController) Post() {
	// 用户登录并保存cookie
	username := self.Input().Get("username")
	password := self.Input().Get("password")
	remember := self.Input().Get("remember") == "1"
	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0
		if remember {
			maxAge = 1<<10 - 1
		}
		self.Ctx.SetCookie("username", username, maxAge, "/")
		self.Ctx.SetCookie("password", password, maxAge, "/")
		self.Redirect("/", 301)
		return
	}
	self.Data["Success"] = false
	self.Data["errMsg"] = "用户名或密码错误"

}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value
	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value
	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}

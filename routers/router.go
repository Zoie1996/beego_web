package routers

import (
	"myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	// 自动路由
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/category", &controllers.CategoryController{})
}

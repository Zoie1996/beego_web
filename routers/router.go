package routers

import (
	"myblog/controllers"
	"os"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	// 自动路由
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	// 创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	// 作为静态文件
	beego.SetStaticPath("/attachment", "attachment")

}

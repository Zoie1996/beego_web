package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.TplName = "home.html"
}
// func (c *MainController) Get() {
// 	c.TplName = "home.html"
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.Data["Hello"] = "Hello guys!"
// 	c.Data["true"] = true
// 	c.Data["false"] = false

// 	type u struct {
// 		Name string
// 		Age  int
// 		Sex  string
// 	}
// 	user := &u{
// 		Name: "Joe",
// 		Age:  13,
// 		Sex:  "男"}
// 	c.Data["User"] = user

// 	nums := []int{1,2,3,4,5,6,7,8,9,0}
// 	c.Data["nums"] = nums

// }

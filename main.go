package main

import (
	"myblog/models"
	_ "myblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 初始化数据库
	models.RegisterDB()
}
func main() {
	// 开启Debug
	orm.Debug = true
	// 自动建表 参数二：false不重建表 参数三：打印相关信息
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITES_DRIVER = "sqlite3"
)

type Category struct {
	ID            int64     //分类ID
	Title         string    // 分类标题
	Created       time.Time `orm:"index;auto_now_add;type(datetime)"` // 创建时间
	Views         int64     `orm:"index"`
	TopicTime     time.Time `orm:"index;auto_now_add;type(datetime)"` // 文章时间
	TopicCount    int64     // 文章统计
	TopicasUserID int64     //用户ID
}

type Topic struct {
	ID             int64
	UID            int64     // 用户ID
	Title          string    // 文章标题
	Content        string    `orm:"size(5000)"`                        //文章内容
	Attachment     string    `orm:"null"`                              // 附件
	Created        time.Time `orm:"index;auto_now_add;type(datetime)"` //创建时间
	Updated        time.Time `orm:"index;auto_now;type(datetime)"`     //更新时间
	Views          int64
	Author         string
	ReplyTime      time.Time `orm:"index;null"` // 最后回复时间
	Replycount     int64     `orm:"default(0)"` //回复统计
	ReplylastUsrID int64     `orm:"null"`       // 恢复用户ID
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {

		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITES_DRIVER, orm.DRSqlite)
	// 默认数据库名称"default" 驱动名称 数据库名称 最大连接数
	orm.RegisterDataBase("default", _SQLITES_DRIVER, _DB_NAME, 10)
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	Categories := make([]*Category, 0)
	qs := o.QueryTable("Category")
	_, err := qs.All(&Categories)
	return Categories, err
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	category := &Category{Title: name}
	qs := o.QueryTable("Category")
	err := qs.Filter("title", name).One(category)
	if err == nil {
		// 已经找到同名分类名
		return err
	}
	// 添加分类
	_, err = o.Insert(category)
	if err != nil {
		// 插入失败
		return err
	}
	// 没有发生错误，返回nil
	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	category := &Category{ID: cid}
	// 删除该分类
	_, err = o.Delete(category)
	return err

}

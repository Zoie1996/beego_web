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
	Views         int64     `orm:"index"`                             //六看
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
	Views          int64     // 浏览
	Author         string    // 作者
	ReplyTime      time.Time `orm:"auto_now;type(datetime)"` // 最后回复时间
	Replycount     int64     `orm:"default(0)"`              // 回复统计
	ReplylastUsrID int64     `orm:"null"`                    // 恢复用户ID
}

// RegisterDB 注册数据库
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

// AddTopic 添加文章
func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{Title: title, Content: content}
	// 添加文章
	_, err := o.Insert(topic)
	return err

}

// GetAllTopics 获取所有文章列表
func GetAllTopics(idDecs bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("Topic")
	var err error
	if idDecs {
		// 根据时间倒序排序
		_, err = qs.OrderBy("-created").All(&topics)

	} else {
		_, err = qs.All(&topics)

	}
	return topics, err
}

// GetAllCategories 获取所有分类列表
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	qs := o.QueryTable("Category")
	_, err := qs.All(&categories)
	return categories, err
}

// AddCategory 添加分类
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
	// 返回err 或 nil
	return err
}

// DelCategory 删除分类
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

package models

import (
	// "os"
	// "path"
	"strconv"
	"time"

	// "github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// _DB_NAME        = "data/beeblog.db"
	_SQLITES_DRIVER = "mysql"
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
	Content        string    `orm:"size(5000)"`                        // 文章内容
	Attachment     string    `orm:"null"`                              // 附件
	Created        time.Time `orm:"index;auto_now_add;type(datetime)"` // 创建时间
	Updated        time.Time `orm:"index;auto_now;type(datetime)"`     // 更新时间
	Views          int64     // 浏览
	Author         string    // 作者
	ReplyTime      time.Time `orm:"auto_now;type(datetime)"` // 最后回复时间
	Replycount     int64     `orm:"default(0)"`              // 回复统计
	ReplylastUsrID int64     `orm:"null"`                    // 恢复用户ID
	Category       *Category `orm:"rel(one);default(1)"`     // 文章分类
}

type Comment struct {
	ID      int64     // 评论ID
	Topic   *Topic    `orm:"rel(one)"` // 文章ID
	Name    string    // 评论昵称
	Content string    `orm:"size(1000)"`                        // 评论内容
	Created time.Time `orm:"index;auto_now_add;type(datetime)"` // 创建时间

}

// RegisterDB 注册数据库
func RegisterDB() {
	// if !com.IsExist(_DB_NAME) {
	// 	os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
	// 	os.Create(_DB_NAME)
	// }
	orm.RegisterDriver(_SQLITES_DRIVER, orm.DRMySQL)
	// 默认数据库名称"default" 驱动名称 数据库名称 最大连接数
	orm.RegisterDataBase("default", _SQLITES_DRIVER, "root:root@/myblog?charset=utf8")
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
}

// AddReply 添加评论

func AddReply(tid, name, content string) error {
	topicid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{ID: topicid}
	comment := &Comment{Topic: topic, Name: name, Content: content}
	_, err = o.Insert(comment)
	return err
}

// GetReplies 获取文章评论
func GetReplies(id string) (replies []*Comment, err error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	topic := &Topic{ID: tid}
	o := orm.NewOrm()
	replies = make([]*Comment, 0)
	qs := o.QueryTable("Comment")
	_, err = qs.Filter("Topic", topic).All(&replies)
	return replies, err

}

// AddTopic 添加文章
func AddTopic(id, title, content, c string) error {
	cid, err := strconv.ParseInt(c, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	category := &Category{ID: cid}

	if len(id) == 0 {
		// 添加文章
		topic := &Topic{Title: title, Content: content, Category: category}

		_, err = o.Insert(topic)
	} else {
		// 修改文章
		err = ModifyTopic(id, title, content, category)
	}
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

// GetTopic 获取文章详情
func GetTopic(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

// GetTopicModify 获取文章详情
func GetTopicModify(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(topic)
	if err != nil {
		return nil, err
	}
	return topic, err
}

func ModifyTopic(id, title, content string, category *Category) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{ID: tid}
	err = o.Read(topic)
	if err == nil {
		topic.Title = title
		topic.Content = content
		topic.Category = category
		o.Update(topic)
	}
	return err
}

// DelTopic 删除文章
func DelTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{ID: tid}
	_, err = o.Delete(topic)
	return err

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

package models

import (
	// "os"
	// "path"
	"os"
	"path"
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
	ID            int64     // 分类ID
	Title         string    // 分类标题
	Created       time.Time `orm:"index;auto_now_add;type(datetime)"` // 创建时间
	Views         int64     `orm:"index"`                             // 浏览量
	TopicTime     time.Time `orm:"index;auto_now_add;type(datetime)"` // 文章时间
	TopicCount    int64     // 文章统计
	TopicasUserID int64     // 用户ID
	Topic         []*Topic  `orm:"reverse(many)"` // 文章分类反向关系
}

type Topic struct {
	ID             int64
	UID            int64      // 用户ID
	Title          string     // 文章标题
	Content        string     `orm:"size(5000)"`                        // 文章内容
	Attachment     string     `orm:"null"`                              // 附件
	Created        time.Time  `orm:"index;auto_now_add;type(datetime)"` // 创建时间
	Updated        time.Time  `orm:"index;auto_now;type(datetime)"`     // 更新时间
	Views          int64      // 浏览
	Author         string     // 作者
	ReplyTime      time.Time  `orm:"auto_now_add;type(datetime)"` // 最后回复时间
	Replycount     int64      `orm:"default(0)"`                  // 回复统计
	ReplylastUsrID int64      `orm:"null"`                        // 回复用户ID
	Category       *Category  `orm:"rel(fk)"`                     // 文章分类
	Comment        []*Comment `orm:"reverse(many)"`               // 文章评论反向关系
}

type Comment struct {
	ID      int64     // 评论ID
	Topic   *Topic    `orm:"rel(fk)"` // 文章ID
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

	// dbType := beego.AppConfig.String("db_type")
	// //连接名称
	// dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	// //数据库名称
	// dbName := beego.AppConfig.String(dbType + "::db_name")
	// //数据库连接用户名
	// dbUser := beego.AppConfig.String(dbType + "::db_user")
	// //数据库连接用户名
	// dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	// //数据库IP（域名）
	// dbHost := beego.AppConfig.String(dbType + "::db_host")
	// //数据库端口
	// dbPort := beego.AppConfig.String(dbType + "::db_port")
	// switch dbType {
	// case "sqlite3":
	// 	orm.RegisterDataBase(dbAlias, dbType, dbName)
	// case "mysql":
	// 	dbCharset := beego.AppConfig.String(dbType + "::db_charset")
	// 	orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
	// 		dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	// }
	// 默认数据库名称"default" 驱动名称 数据库名称 最大连接数
	orm.RegisterDataBase("default", _SQLITES_DRIVER, "root:root@/myblog?charset=utf8&loc=Asia%2FShanghai")
	// orm.RegisterDataBase("default", _SQLITES_DRIVER, "root:root@tcp(127.0.0.1:3306)/myblog?charset=utf8", 30)
	// orm.Debug = true
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
}

func DeleteReply(id, tid string) error {
	rid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	topicid, err1 := strconv.ParseInt(tid, 10, 64)
	if err1 != nil {
		return err1
	}
	o := orm.NewOrm()
	reply := &Comment{ID: rid}
	if o.Read(reply) == nil {
		_, err = o.Delete(reply)
		if err != nil {
			return nil
		}
	}

	// topic := new(Topic)
	// qs := o.QueryTable("topic")
	// err = qs.Filter("id", topicid).One(topic)
	// if err == nil {
	// 	topic.Replycount--
	// 	o.Update(topic)
	// }
	replies := make([]*Comment, 0)
	_, err = o.QueryTable("comment").Filter("topic_id", tid).OrderBy("-created").All(&replies)
	topic := &Topic{ID: topicid}
	if o.Read(topic) == nil {
		topic.Replycount = int64(len(replies))
		topic.ReplyTime = replies[0].Created
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}
	return err
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
	topic = new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", topicid).One(topic)
	if err == nil {
		topic.Replycount++
		topic.ReplyTime = time.Now()
		o.Update(topic)
	}
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
func AddTopic(id, title, content, categoryID, attachment string) error {
	cid, err := strconv.ParseInt(categoryID, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	category := &Category{ID: cid}

	if len(id) == 0 {
		// 添加文章
		topic := &Topic{Title: title, Content: content, Category: category, Attachment: attachment}
		_, err = o.Insert(topic)
		category = new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("id", cid).One(category)
		if err != nil {
			return err
		}
		category.TopicCount++
		_, err = o.Update(category)
	} else {
		// 修改文章
		err = ModifyTopic(id, title, content, attachment, category)
	}
	return err

}

// GetAllTopics 获取所有文章列表
func GetAllTopics(id string, idDecs bool) ([]*Topic, error) {

	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("Topic")
	if cid > 0 {
		// 分类
		qs = qs.Filter("category_id", cid)
	}
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

// ModifyTopic 修改文章
func ModifyTopic(id, title, content, attachment string, category *Category) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{ID: tid}
	err = o.Read(topic)
	// 修改文章分类判断
	if topic.Category != category {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("id", topic.Category.ID).One(cate)
		if err != nil {
			return err
		}
		cate.TopicCount--
		_, err = o.Update(cate)
		if err != nil {
			return err
		}
		err = qs.Filter("id", category.ID).One(cate)
		if err != nil {
			return err
		}
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	var OldAttach string
	if err == nil {
		OldAttach = topic.Attachment
		topic.Title = title
		topic.Content = content
		topic.Category = category
		topic.Attachment = attachment
		o.Update(topic)
	}
	// 更新后删除旧文件
	if len(OldAttach) > 0 {
		os.Remove(path.Join("attachment", OldAttach))
	}
	return err
}

// DelTopic 删除文章
func DelTopic(id, cid string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	categoryid, err1 := strconv.ParseInt(cid, 10, 64)
	if err1 != nil {
		return err1
	}
	o := orm.NewOrm()
	topic := &Topic{ID: tid}
	category := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", categoryid).One(category)
	if err != nil {
		return err
	}
	category.TopicCount--
	_, err = o.Update(category)
	if err != nil {
		return err
	}
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

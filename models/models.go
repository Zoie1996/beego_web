package models

import (
	"os"
	"path"
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
	ID            int64
	Title         string
	Created       time.Time `orm:"index"`
	Views         int64     `orm:"index"`
	TopicTime     time.Time `orm:"index"`
	TopicCount    int64
	TopicasUserID int64
}

type Topic struct {
	ID             int64
	UID            int64
	Title          string
	Content        string `orm:"size(5000)"`
	Attaclment     string
	Created        time.Time `orm:"index"`
	Updated        time.Time `orm:"index"`
	Views          int64
	Author         string
	ReplyTime      time.Time `orm:"index"`
	Replycount     int64
	ReplylastUsrID int64
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

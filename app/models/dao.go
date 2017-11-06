package models

import (
	"gopkg.in/mgo.v2"
)

const (
	DbName            = "goblog"
	BlogCollection    = "blogs"
	CommentCollection = "gb_comments"
	MessageCollection = "gb_messages"
	HistoryCollection = "gb_histories"
	EmailCollection   = "gb_emails"
	BaseYear          = 2017
)

type Dao struct {
	session *mgo.Session
}

func NewDao() (*Dao, error) {
	//mongodb数据库连接
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	return &Dao{session}, nil
}

func (d *Dao) Close() {
	d.session.Close()
}

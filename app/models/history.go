package models

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type History struct {
	Year  int
	Blogs []Blog
}

//插入归档
func (dao *Dao) InsertHistory(history *History) error {
	historyCollection := dao.session.DB(DbName).C(HistoryCollection)
	err := historyCollection.Insert(history)
	if err != nil {
		revel.WARN.Printf("Unable to save history: %v error %v", history, err)
	}
	return err
}

//查找归档
func (dao *Dao) FindHistory() []History {
	historyCollection := dao.session.DB(DbName).C(HistoryCollection)
	his := []History{}
	query := historyCollection.Find(bson.M{}).Sort("-year")
	query.All(&his)
	return his
}

//删除归档
func (dao *Dao) RemoveAll() error {
	historyCollection := dao.session.DB(DbName).C(HistoryCollection)
	_, err := historyCollection.RemoveAll(bson.M{})
	if err != nil {
		revel.WARN.Printf("Unable to remove all: error %v", err)
	}
	return err
}

//创建所有归档
func (dao *Dao) CreateAllHistory() {
	dao.RemoveAll()
	var end int = time.Now().Year()
	for i := BaseYear; i <= end; i++ {
		history := new(History)
		history.Year = i
		dao.InsertHistory(history)
	}
}

package model

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var DBOrmInstance *xorm.Engine

func CreateDBInstance() {
	var err error
	DBOrmInstance, err = xorm.NewEngine("sqlite3", "./data/stardict.db")
	if err != nil {
		logrus.Error(err)
		return
	}
	DBOrmInstance.ShowSQL(true)

	// err = DBOrmInstance.CreateTables(&WordModel{})
	// if err != nil {
	// 	logrus.Error(err)
	// 	return
	// }
}

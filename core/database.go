package core

import (
	"dance/conf"
	"io/ioutil"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	mysql *gorm.DB
)

func GetDB() *gorm.DB {
	return mysql
}

func InitDB() {
	logger := log.New(ioutil.Discard, "", log.LstdFlags)

	confs := conf.Config.Database
	dbs, err := gorm.Open("mysql", confs)
	if err != nil {
		conf.MainLog.Errorf("连接数据库异常：%v", err.Error())
	}
	dbs.SetLogger(logger)
	dbs.DB().SetMaxOpenConns(100)
	dbs.DB().SetConnMaxIdleTime(time.Hour)
	dbs.DB().SetConnMaxLifetime(time.Hour)

	mysql = dbs
}

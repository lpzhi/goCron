package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/jinzhu/gorm"
	"goCron/prepare/dottask/pkg/setting"
)

var (
	db, db_role *gorm.DB
)

type dbConfig struct {
	dbType      string
	dbName      string
	user        string
	password    string
	host        string
	tablePrefix string
}

type BaseMode struct {
}

/*
连接默认数据库 default_database role_total  ..._central库
*/
func init() {
	db = ConectDb("default_database")
	db_role = ConectDb("role_total")
}

/**
连接数据库
*/
func ConectDb(dbType string) *gorm.DB {

	config := getDbConfig(dbType)
	dbConnect, err := gorm.Open(config.dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.user,
		config.password,
		config.host,
		config.dbName))

	if err != nil {
		log.Fatal("conect db err:%v", err)
	}

	dbConnect.SingularTable(true)
	dbConnect.LogMode(true)
	dbConnect.DB().SetMaxIdleConns(10)
	dbConnect.DB().SetMaxOpenConns(100)

	return dbConnect
}

/*
根据配置获取数据库实例
*/
func (baaseMode *BaseMode) getDb(config interface{}) {
}

/*
根据传入的数据库类型 role_total default_database
*/
func getDbConfig(dbype string) dbConfig {
	sec, err := setting.Cfg.GetSection(dbype)
	config := dbConfig{}

	if err != nil {
		log.Fatal("fail to get section: %v  err: %v", dbype, err)
	}

	config.dbName = sec.Key("NAME").String()
	config.dbType = sec.Key("TYPE").String()
	config.host = sec.Key("HOST").String()
	config.password = sec.Key("PASSWORD").String()
	config.user = sec.Key("USER").String()
	config.tablePrefix = sec.Key("TABLE_PREFIX").String()
	return config
}

func CloseDB() {
	defer db.Close()
}

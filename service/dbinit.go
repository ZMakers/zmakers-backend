package service

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


type database struct {
	MaxConn int
	MaxOpen int
}
var databaseConfig = new(database) //设置全局的引用型指针变量


var Db *gorm.DB

func DbInit() *gorm.DB {
	username := "root"  // account
	password := "Zjj19911031" // password
	host := "127.0.0.1" // db address
	port := 3306 // db port
	Dbname := "zmakers" // db name
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//mysqlDb, err := Db.DB()
	//if err != nil {
	//	panic(fmt.Sprintf("get mysql db error %s", err))
	//}
	//database := getConfig()
	//mysqlDb.SetMaxIdleConns(database.MaxConn)
	//mysqlDb.SetMaxOpenConns(database.MaxOpen)
	return Db
}

func DbClose() {
	DB, err := Db.DB()
	if err != nil {
		panic("get DB error!")
	}
	err = DB.Close()
	if err != nil {
		panic("close DB error!")
	}
}


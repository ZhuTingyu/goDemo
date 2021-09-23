package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//config2 "github.com/xiaoqinghong/hi_robot/server/config"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var thisDB *gorm.DB

func Init() {
	var connection = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "yupaoapp", "123.57.208.24", 3306, "hi_robot_test")
	var err error
	fmt.Println(connection)
	thisDB, err = gorm.Open("mysql", connection)
	//thisDB, err = gorm.Open("sqlite3", config.Config().Sqlite3.Path)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库链接成功")
}

func BeginTx() (tx *gorm.DB) {
	tx = thisDB.Begin()
	return
}

type Dao struct {
	db *gorm.DB
}

func (d *Dao) SetTx(tx *gorm.DB) {
	if d != nil {
		d.db = tx
	}
}

func (d *Dao) GetDB() (db *gorm.DB) {
	if d != nil && d.db != nil {
		db = d.db
	} else {
		db = thisDB
	}
	return
}

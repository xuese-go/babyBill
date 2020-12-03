package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xuese-go/babyBill/structs"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/babyBill?charset=utf8&parseTime=true&loc=Local"
	if Db, err = gorm.Open("mysql", dsn); err != nil {
		log.Println("数据库连接失败")
		log.Println(err.Error())
	} else {
		log.Println("数据库连接成功")
		//打印sql
		Db.LogMode(true)
		//创建表
		tables := make([]interface{}, 0)
		tables = append(tables, &structs.Record{})

		for k := range tables {
			Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(tables[k])
		}
	}
}

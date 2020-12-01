package service

import (
	"github.com/boltdb/bolt"
	uuid "github.com/satori/go.uuid"
	"github.com/xuese-go/babyBill/db"
	"log"
)

type Record struct {
	Uuid   string
	Dates  string
	Money  string
	Matter string
}

func Save(record Record) error {
	//插入数据
	err := db.Db.Update(func(tx *bolt.Tx) error {

		//取出叫"record_table"的表
		b := tx.Bucket([]byte("record_table"))

		//往表里面存储数据
		if b != nil {
			//插入的键值对数据类型必须是字节数组
			err := b.Put([]byte(uuid.NewV4().String()), []byte(record.Dates+","+record.Money+","+record.Matter))
			if err != nil {
				log.Println(err.Error())
				return err
			}
		}

		//一定要返回nil
		return nil
	})

	//更新数据库失败
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

package db

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

var Db *bolt.DB

func init() {
	//打开我的数据库当前目录中的数据文件。
	//如果它不存在，它将被创建。
	//同时只能打开一次，所以防止无限等待则设置超时时间
	Db, err := bolt.Open("babyBill.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	//	创建表
	err = Db.Update(func(tx *bolt.Tx) error {

		//判断要创建的表是否存在
		b := tx.Bucket([]byte("record_table"))
		if b == nil {

			//创建叫"record_table"的表
			_, err := tx.CreateBucket([]byte("record_table"))
			if err != nil {
				log.Println("表record_table创建失败", err.Error())
			}
		}

		//一定要返回nil
		return nil
	})

	//更新数据库失败
	if err != nil {
		log.Println(err)
	}

	//	读写事务
	//	err := db.Update(func(tx *bolt.Tx) error {
	//		...
	//		return nil
	//	})
	//只读事务
	//	err := db.View(func(tx *bolt.Tx) error {
	//		...
	//		return nil
	//	})

	//	手动事务
	// Start a writable transaction.
	//tx, err := db.Begin(true)
	//if err != nil {
	//	return err
	//}
	//defer tx.Rollback()
	//
	//// Use the transaction...
	//_, err := tx.CreateBucket([]byte("MyBucket"))
	//if err != nil {
	//	return err
	//}
	//
	//// Commit the transaction and check for error.
	//if err := tx.Commit(); err != nil {
	//	return err
	//}
}

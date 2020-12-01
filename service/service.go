package service

import (
	"github.com/boltdb/bolt"
	uuid "github.com/satori/go.uuid"
	"log"
)

func Save(dates string, money string, matter string) error {
	db, err := bolt.Open("babyBill.db", 0600, &bolt.Options{Timeout: 1 * 5000})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//插入数据
	err = db.Update(func(tx *bolt.Tx) error {

		//取出叫"record_table"的表
		b := tx.Bucket([]byte("record_table"))

		//往表里面存储数据
		if b != nil {
			//插入的键值对数据类型必须是字节数组
			err := b.Put([]byte(uuid.NewV4().String()), []byte(dates+","+money+","+matter))
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

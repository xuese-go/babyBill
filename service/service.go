package service

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type Record struct {
	Dates  string `json:"dates"`
	Money  string `json:"money"`
	Matter string `json:"matter"`
}

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
			model := new(Record)
			model.Dates = dates
			model.Money = money
			model.Matter = matter
			s, _ := json.Marshal(model)
			var t int64 = time.Now().Unix()
			var dateStr string = time.Unix(t, 0).Format("2006-01-02 15:04:05")
			err := b.Put([]byte(dateStr), s)
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

func Find(dates string) ([]*Record, error) {
	db, err := bolt.Open("babyBill.db", 0600, &bolt.Options{Timeout: 1 * 5000})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data := make([]*Record, 0)
	//查询数据
	err = db.View(func(tx *bolt.Tx) error {

		//取出叫"record_table"的表
		b := tx.Bucket([]byte("record_table"))

		if b != nil {
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				log.Printf("key=%s, value=%s\n", k, v)
				m := new(Record)
				s := fmt.Sprintf("%s", v)
				b := []byte(s)
				if err := json.Unmarshal(b, &m); err != nil {
					log.Println(err)
				} else {
					data = append(data, m)
				}
			}
		}

		//一定要返回nil
		return nil
	})

	//更新数据库失败
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return data, nil
}

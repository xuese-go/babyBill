package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xuese-go/babyBill/db"
	"github.com/xuese-go/babyBill/structs"
	"github.com/xuese-go/babyBill/util"
	"log"
)

/**
新增用户
*/
func Save(record structs.Record) structs.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()

	//填充其它数据
	uid := uuid.NewV4().String()
	record.Uuid = uid

	//新增数据
	t := tx.Create(record)
	if t.Error != nil {
		t.Rollback()
		return structs.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	t.Commit()
	return structs.ResponeStruct{Success: true, Msg: "操作成功"}
}

/**
删除
*/
func DeleteById(uuid string) structs.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	var u structs.Record
	if err := tx.First(&u, "uuid = ?", uuid).Delete(&u).Error; err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return structs.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	tx.Commit()
	return structs.ResponeStruct{Success: true, Msg: "操作成功"}
}

/**
修改
*/
func Update(record structs.Record) structs.ResponeStruct {
	dba := db.Db
	tx := dba.Begin()
	var u structs.Record
	if err := dba.First(&u, "uuid = ?", record.Uuid).Error; err != nil {
		log.Println(err.Error())
		return structs.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	if record.Dates != "" {
		u.Dates = record.Dates
	}
	if record.Money != 0 {
		u.Money = record.Money
	}
	if record.Remarks != "" {
		u.Remarks = record.Remarks
	}
	t := tx.Save(&u)
	if t.Error != nil {
		t.Rollback()
		log.Println(t.Error.Error())
		return structs.ResponeStruct{Success: false, Msg: "失败"}
	}
	t.Commit()
	return structs.ResponeStruct{Success: true, Msg: "成功"}
}

/**
根据id查询
*/
func One(uuid string) structs.ResponeStruct {
	dba := db.Db
	var u structs.Record
	if err := dba.First(&u, "uuid = ?", uuid).Error; err != nil {
		log.Println(err.Error())
		return structs.ResponeStruct{Success: false, Msg: "查询错误"}
	}
	return structs.ResponeStruct{Success: true, Msg: "操作成功", Data: u}
}

/**
分页查询
*/
func Page(pageNum int, pageSize int, record structs.Record) structs.ResponeStruct {
	//为了不影响后边的操作  所以需要使用新的变量
	dba := db.Db
	us := make([]structs.Record, 0)

	//查询条件
	if record.Dates != "" {
		dba = dba.Where("dates like ?", "%"+record.Dates+"%")
	}
	//总记录数
	var count int
	dba = dba.Find(&us).Count(&count)
	if dba.Error != nil {
		log.Println(dba.Error.Error())
		return structs.ResponeStruct{Success: false, Msg: "操作失败"}
	}

	//分页信息
	if pageNum > 0 && pageSize > 0 {
		dba = dba.Order("dates")
		dba = dba.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}
	//查询
	if err := dba.Table("record_table").Select([]string{"uuid", "dates", "money", "remarks"}).Scan(&us).Error; err != nil {
		log.Println(err.Error())
		return structs.ResponeStruct{Success: false, Msg: "操作失败"}
	}
	return structs.ResponeStruct{Success: true, Data: us, Page: util.PageUtil(count, pageSize, pageNum)}
}

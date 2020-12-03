package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/babyBill/service"
	"github.com/xuese-go/babyBill/structs"
	"log"
	"strconv"
)

/*
新增用户
*/
func Save(ctx *gin.Context) {
	var recor structs.Record
	if err := ctx.ShouldBind(&recor); err != nil {
		structs.Respone(ctx, structs.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.Panic(err.Error())
		return
	}
	respond := service.Save(recor)
	structs.Respone(ctx, respond)
}

/*
根据主键删除
*/
func Delete(ctx *gin.Context) {
	uuid := ctx.Param("deleteId")
	respond := service.DeleteById(uuid)
	structs.Respone(ctx, respond)
}

/**
根据id修改
*/
func Update(ctx *gin.Context) {
	uuid := ctx.Param("putId")
	var recor structs.Record
	if err := ctx.ShouldBind(&recor); err != nil {
		structs.Respone(ctx, structs.ResponeStruct{Success: false, Msg: "参数绑定错误"})
		log.Panic(err.Error())
		return
	}

	recor.Uuid = uuid
	respond := service.Update(recor)
	structs.Respone(ctx, respond)
}

/**
根据主键查询
*/
func One(ctx *gin.Context) {
	uuid := ctx.Param("getId")
	respond := service.One(uuid)
	structs.Respone(ctx, respond)
}

/**
分页
*/
func Page(ctx *gin.Context) {
	pageNum := ctx.Query("pageNum")
	pageSize := ctx.Query("pageSize")
	dates := ctx.Query("dates")

	var recor structs.Record
	recor.Dates = dates
	n, _ := strconv.Atoi(pageNum)
	s, _ := strconv.Atoi(pageSize)
	res := service.Page(n, s, recor)

	structs.Respone(ctx, res)
}

/**
统计 根据传的时间参数 sql灵活统计  date >= d and date <= d
*/
func Statistics(ctx *gin.Context) {
	dates := ctx.Query("dates")
	if dates != "" {
		res := service.Statistics(dates)
		structs.Respone(ctx, res)
	} else {
		structs.Respone(ctx, structs.ResponeStruct{Success: false, Msg: "需要输入时间"})
	}
}

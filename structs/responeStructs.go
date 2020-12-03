package structs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
统一返回
*/
type ResponeStruct struct {
	//业务是否处理成功(查询类是否有返回数据)
	Success bool `json:"success"`
	//消息
	Msg string `json:"msg"`
	//数据
	Data interface{} `json:"data"`
	//	分页信息
	Page interface{} `json:"page"`
}

func Respone(context *gin.Context, resp ResponeStruct) {
	context.JSON(http.StatusOK, resp)
}

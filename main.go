package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuese-go/babyBill/controller"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	winExecError = map[uint32]string{
		0:  "The system is out of memory or resources.",
		2:  "The .exe file is invalid.",
		3:  "The specified file was not found.",
		11: "The specified path was not found.",
	}
)

func init() {
	// 加载默认配置
	r := gin.Default()

	//日志
	file, _ := os.Create("sys.log")
	log.SetOutput(file) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// 路由
	routers(r)

	// 启动并监听默认8080端口
	_ = r.Run()

}

func main() {
}

// 路由绑定路径集合
func routers(r *gin.Engine) {

	//模板路径-html文件地址
	r.LoadHTMLGlob("views/**/*")
	//静态文件路径
	r.Static("/static", "static")

	//页面路由
	index := r.Group("/")
	{
		//页面处理
		ind := index.Group("/")
		//主页面-登录页面
		ind.GET("/", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index/index.html", nil)
		})
	}

	//api路由
	apis := r.Group("/api")
	{
		record := apis.Group("/record")
		record.POST("/record", controller.Save)
		record.DELETE("/record/:deleteId", controller.Delete)
		record.PUT("/record/:putId", controller.Update)
		record.GET("/record/:getId", controller.One)
		record.GET("/records", controller.Page)
		//统计
		record.GET("/statistics", controller.Statistics)

	}

}

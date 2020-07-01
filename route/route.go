package route

import (
	"github.com/gin-gonic/gin"
	"machine-geek.cn/recruit-server/common"
	"machine-geek.cn/recruit-server/controller"
)

func Route(engine *gin.Engine) {
	// 静态文件服务
	engine.Static("/static", "./static")
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 上传文件
	engine.POST("/upload/file", controller.UploadFile)
	engine.POST("/upload/picture", controller.UploadPicture)
	// 路由
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, common.Ok())
	})
}

package route

import (
	"github.com/gin-gonic/gin"
	"machine-geek.cn/recruit-server/common"
	"machine-geek.cn/recruit-server/controller"
	"net/http"
)

func Route(engine *gin.Engine) {
	// 跨域请求
	engine.Use(cors())
	// 静态文件服务
	engine.Static("/static", "./static")
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 上传文件
	engine.POST("/upload/file", controller.UploadFile)
	// 上传图片
	engine.POST("/upload/picture", controller.UploadPicture)
	// 测试路由
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, common.Ok())
	})
	// 登陆Admin
	engine.POST("/admin/login", controller.LoginAdmin)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

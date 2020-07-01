package route

import "github.com/gin-gonic/gin"

func Route(engine *gin.Engine) {
	// 静态文件服务
	engine.Static("/static", "./static")
	// 路由
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"ok":  true,
			"msg": "success",
		})
	})
}

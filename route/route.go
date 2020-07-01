package route

import "github.com/gin-gonic/gin"

func Route(engine *gin.Engine) {
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"ok":  true,
			"msg": "success",
		})
	})
}

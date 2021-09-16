package router

import "github.com/gin-gonic/gin"

func SetupRouter(request *gin.Engine) *gin.Engine {
	request.GET("/wechat/ready", func(ctx *gin.Context) {
		ctx.String(200, "success")
	})

	return request
}

package router

import (
	"github.com/gin-gonic/gin"
	"wechat-go/controller/user"
)

func SetupRouter(request *gin.Engine) *gin.Engine {
	request.GET("/wechat/ready", func(ctx *gin.Context) {
		ctx.String(200, "success")
	})
	request.POST("/wechat/login", user.Login)
	return request
}

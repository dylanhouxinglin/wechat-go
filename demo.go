package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 小写读取不成功
type loginParams struct {
	Nickname 	string 		`json:"nickName"  binding:"required"`
	LoginTime	string   	`json:"loginTime" binding:"required"`
}


func setupRouter()  {
	router := gin.Default()
	router.GET("/wechat/ready", func(ctx *gin.Context) {
		ctx.String(200, "success")
	})
	router.POST("/wechat/login", func(ctx *gin.Context) {
		var in loginParams
		if err := ctx.ShouldBindJSON(&in); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{ "status": "200", "nickname":in.Nickname})
	})
	_ = http.ListenAndServe(":8989", router)
}

func main()  {
	setupRouter()
}
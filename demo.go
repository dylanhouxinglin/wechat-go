package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter()  {
	router := gin.Default()
	router.GET("/api/test", func(ctx *gin.Context) {
		ctx.String(200, "success")
	})
	_ = http.ListenAndServe(":8989", router)
}

func main()  {
	setupRouter()
}
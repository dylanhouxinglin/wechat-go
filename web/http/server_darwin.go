package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"wechat-go/router"
)


func Run() error {
	appEngine := gin.Default()
	appServer := router.SetupRouter(appEngine)
	server := &http.Server{
		Addr: 			viper.GetString(`server.address`),
		ReadTimeout: 	viper.GetDuration(`server.readTimeout`)*time.Second,
		WriteTimeout:   viper.GetDuration(`server.writeTimeout`)*time.Second,
		Handler: 		appServer,
	}
	//_ = http.ListenAndServe(server.Addr, appServer)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
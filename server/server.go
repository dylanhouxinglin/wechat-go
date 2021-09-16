package server

import (
	"fmt"
	"wechat-go/web/http"
)

func Serve()  {
	err := http.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

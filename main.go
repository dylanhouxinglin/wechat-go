package main

import (
	"wechat-go/preload"
	"wechat-go/server"
)

func main()  {
	preload.Init()
	server.Serve()
}